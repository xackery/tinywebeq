package main

import (
	"context"
	"crypto"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"time"

	"github.com/go-acme/lego/challenge/tlsalpn01"
	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/challenge/http01"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
	"github.com/rs/zerolog"
	"go.uber.org/zap"
	"golang.org/x/crypto/acme/autocert"

	"github.com/xackery/tinywebeq/questparse"
	"github.com/xackery/tinywebeq/repo"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/handlers"
	"github.com/xackery/tinywebeq/image"
	"github.com/xackery/tinywebeq/recipe"
	"github.com/xackery/tinywebeq/store"
	"github.com/xackery/tinywebeq/template"
	"github.com/xackery/tinywebeq/tlog"
)

// Version is the build version
var Version string

type application struct {
	templates fs.FS
	logger    *zap.SugaredLogger
	handlers  *handlers.Handlers
	db        *repo.Repo
}

func main() {
	err := run()
	if err != nil {
		tlog.Errorf("Critical fail: %s", err)
		os.Exit(1)
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		log.Println(http.ListenAndServe("localhost:8082", nil))
	}()
	var err error

	_, err = config.NewConfig(ctx)
	if err != nil {
		return fmt.Errorf("config.NewConfig: %w", err)
	}
	tlog.SetLevel(zerolog.InfoLevel)

	if config.Get().IsDebugEnabled {
		tlog.SetLevel(zerolog.DebugLevel)
	}

	tlog.Init(nil, os.Stdout)

	dbNew, err := repo.New(tlog.Sugar, "peqdb:peqdb@tcp(192.168.2.10:3306)/peq?parseTime=true&columnsWithAlias=true")
	if err != nil {
		return fmt.Errorf("repo.New: %w", err)
	}

	app := &application{
		templates: template.FS,
		logger:    tlog.Sugar,
		handlers:  handlers.New(tlog.Sugar, template.FS),
		db:        dbNew,
	}

	args := os.Args
	if len(args) < 2 {
		fmt.Println("no arguments provided")
		usage()
	}
	isCacheFlush := false
	switch args[1] {
	case "version":
		fmt.Println(Version)
		return nil
	case "quest":
		return questParse(ctx)
	case "recipe":
		return recipeParse(ctx)
	case "help":
		usage()
	case "flush":
		isCacheFlush = true
	case "letsencrypt":
		err = letsencrypt()
		if err != nil {
			return fmt.Errorf("letsencrypt: %w", err)
		}
	case "server":
	default:
		fmt.Println("unknown command:", args[1])
		usage()
	}
	if Version == "" {
		Version = "1.x.x EXPERIMENTAL"
	}
	tlog.Infof("Starting tinywebeq %s", Version)

	if isCacheFlush {
		err = os.RemoveAll("cache")
		if err != nil {
			return fmt.Errorf("remove cache: %w", err)
		}
		tlog.Infof("Cache flushed")
		return nil
	}

	err = db.Init(ctx)
	if err != nil {
		return fmt.Errorf("db.Init: %w", err)
	}

	err = store.Init(ctx)
	if err != nil {
		return fmt.Errorf("store.Init: %w", err)
	}

	err = os.MkdirAll("cache", 0755)
	if err != nil {
		return fmt.Errorf("make cache: %w", err)
	}

	err = image.Init(ctx)
	if err != nil {
		return fmt.Errorf("image.Init: %w", err)
	}

	certPath := config.Get().Site.LetsEncrypt.CertPath
	server := &http.Server{
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	if config.Get().Site.LetsEncrypt.IsEnabled {
		tlog.Debugf("Letsencrypt enabled")
		server.Addr = fmt.Sprintf(":%d", 443)
		m := &autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(config.Get().Site.LetsEncrypt.Domains...),
			Cache:      autocert.DirCache(certPath),
		}
		server.TLSConfig = m.TLSConfig()
	} else {
		server.Addr = fmt.Sprintf("%s:%d", config.Get().Site.Host, config.Get().Site.Port)
	}

	server.Handler = app.routes()

	tlog.Infof("Listening on %s", server.Addr)
	if config.Get().Site.LetsEncrypt.IsEnabled {
		return server.ListenAndServeTLS("", "")
	}
	return server.ListenAndServe()
}

func usage() {
	fmt.Println("Usage: tinywebeq [server|quest|letsencrypt|version|flush|help]")
	os.Exit(1)
}

// LetsEncryptUser is a user for letsencrypt
type LetsEncryptUser struct {
	Email        string
	Registration *registration.Resource
	key          crypto.PrivateKey
}

// GetEmail returns the email
func (u *LetsEncryptUser) GetEmail() string {
	return u.Email
}

// GetRegistration returns the registration
func (u *LetsEncryptUser) GetRegistration() *registration.Resource {
	return u.Registration
}

// GetPrivateKey returns the private key
func (u *LetsEncryptUser) GetPrivateKey() crypto.PrivateKey {
	return u.key
}

func letsencrypt() error {
	if !config.Get().Site.LetsEncrypt.IsEnabled {
		return fmt.Errorf("letsencrypt is disabled in config, enable it first")
	}
	if len(config.Get().Site.LetsEncrypt.Email) == 0 {
		return fmt.Errorf("letsencrypt email is not set in config")
	}
	if len(config.Get().Site.LetsEncrypt.Domains) < 1 {
		return fmt.Errorf("letsencrypt domains is not set in config")
	}

	certPath := config.Get().Site.LetsEncrypt.CertPath
	// remove trailing slash
	certPath = filepath.Clean(certPath)
	certPath = filepath.ToSlash(certPath)

	fi, err := os.Stat(certPath)
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("os.Stat: %w", err)
		}
		tlog.Infof("%s does not exist, creating", certPath)
		err = os.MkdirAll(certPath, 0755)
		if err != nil {
			return fmt.Errorf("os.MkdirAll: %w", err)
		}
		fi, err = os.Stat(certPath)
		if err != nil {
			return fmt.Errorf("os.Stat: %w", err)
		}
	}
	if !fi.IsDir() {
		return fmt.Errorf("letsencrypt cert path is not a directory")
	}

	var privateKey crypto.PrivateKey
	// check if private key exists
	keyPath := config.Get().Site.LetsEncrypt.CertPath + "/key.pem"
	if _, err := os.Stat(keyPath); err == nil {
		tlog.Infof("%s exists, doing renew instead", keyPath)
		return letsEncryptRenew()
	}
	tlog.Infof("%s does not exist, generating", keyPath)

	privateKey, err = certcrypto.GeneratePrivateKey(certcrypto.RSA2048)
	if err != nil {
		return fmt.Errorf("certcrypto.GeneratePrivateKey: %w", err)
	}

	myUser := LetsEncryptUser{
		Email: config.Get().Site.LetsEncrypt.Email,
		key:   privateKey,
	}

	leConfig := lego.NewConfig(&myUser)

	leConfig.CADirURL = lego.LEDirectoryStaging
	if config.Get().Site.LetsEncrypt.IsProd {
		leConfig.CADirURL = lego.LEDirectoryProduction
	}

	leConfig.Certificate.KeyType = certcrypto.RSA2048

	// A client facilitates communication with the CA server.
	client, err := lego.NewClient(leConfig)
	if err != nil {
		return fmt.Errorf("lego.NewClient: %w", err)
	}

	// We specify an HTTP port of 5002 and an TLS port of 5001 on all interfaces
	// because we aren't running as root and can't bind a listener to port 80 and 443
	// (used later when we attempt to pass challenges). Keep in mind that you still
	// need to proxy challenge traffic to port 5002 and 5001.
	err = client.Challenge.SetHTTP01Provider(http01.NewProviderServer("", "5002"))
	if err != nil {
		return fmt.Errorf("client.Challenge.SetHTTP01Provider: %w", err)
	}
	err = client.Challenge.SetTLSALPN01Provider(tlsalpn01.NewProviderServer("", "5001"))
	if err != nil {
		return fmt.Errorf("client.Challenge.SetTLSALPN01Provider: %w", err)
	}

	tlog.Infof("Listening on :5002 for HTTP-01 challenge and :5001 for TLS-ALPN-01 challenge")

	// New users will need to register
	reg, err := client.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true})
	if err != nil {
		return fmt.Errorf("register: %w", err)
	}
	myUser.Registration = reg

	request := certificate.ObtainRequest{
		Domains: config.Get().Site.LetsEncrypt.Domains,
		Bundle:  true,
	}
	certificates, err := client.Certificate.Obtain(request)
	if err != nil {
		return fmt.Errorf("Certificate.Obtain: %w", err)
	}

	// Each certificate comes back with the cert bytes, the bytes of the client's
	// private key, and a certificate URL. SAVE THESE TO DISK.
	fmt.Printf("%#v\n", certificates)

	// Save certs
	err = os.WriteFile(certPath+"/cert.pem", certificates.Certificate, 0644)
	if err != nil {
		return fmt.Errorf("write cert.pem: %w", err)
	}

	err = os.WriteFile(certPath+"/key.pem", certificates.PrivateKey, 0644)
	if err != nil {
		return fmt.Errorf("write key.pem: %w", err)
	}

	err = os.WriteFile(certPath+"/chain.pem", certificates.IssuerCertificate, 0644)
	if err != nil {
		return fmt.Errorf("write chain.pem: %w", err)
	}

	err = os.WriteFile(certPath+"/csr.pem", certificates.CSR, 0644)
	if err != nil {
		return fmt.Errorf("write csr.pem: %w", err)
	}

	tlog.Infof("Letsencrypt done!")
	return nil
}

func questParse(ctx context.Context) error {
	err := db.Init(ctx)
	if err != nil {
		return fmt.Errorf("db.Init: %w", err)
	}

	err = store.Init(ctx)
	if err != nil {
		return fmt.Errorf("store.Init: %w", err)
	}

	err = questparse.Parse(ctx, config.Get().Quest.ActiveConcurrency)
	if err != nil {
		return fmt.Errorf("questParse: %w", err)
	}
	return nil
}

func letsEncryptRenew() error {
	certPath := config.Get().Site.LetsEncrypt.CertPath
	keyPath := certPath + "/key.pem"
	data, err := os.ReadFile(keyPath)
	if err != nil {
		return fmt.Errorf("os.ReadFile: %w", err)
	}
	privateKey, err := certcrypto.ParsePEMPrivateKey(data)
	if err != nil {
		return fmt.Errorf("certcrypto.ParsePrivateKey: %w", err)
	}

	myUser := LetsEncryptUser{
		Email: config.Get().Site.LetsEncrypt.Email,
		key:   privateKey,
	}

	leConfig := lego.NewConfig(&myUser)

	leConfig.CADirURL = lego.LEDirectoryStaging
	if config.Get().Site.LetsEncrypt.IsProd {
		leConfig.CADirURL = lego.LEDirectoryProduction
	}

	leConfig.Certificate.KeyType = certcrypto.RSA2048

	// A client facilitates communication with the CA server.
	client, err := lego.NewClient(leConfig)
	if err != nil {
		return fmt.Errorf("lego.NewClient: %w", err)
	}

	// We specify an HTTP port of 5002 and an TLS port of 5001 on all interfaces
	// because we aren't running as root and can't bind a listener to port 80 and 443

	err = client.Challenge.SetHTTP01Provider(http01.NewProviderServer("", "5002"))
	if err != nil {
		return fmt.Errorf("client.Challenge.SetHTTP01Provider: %w", err)
	}

	// err = client.Challenge.SetTLSALPN01Provider(tlsalpn01.NewProviderServer("", "5001"))
	// if err != nil {
	// 	return fmt.Errorf("client.Challenge.SetTLSALPN01Provider: %w", err)
	// }

	return nil
}

// func (s *CertificatesStorage) SaveResource(certRes *certificate.Resource) {
// 	domain := certRes.Domain

// 	// We store the certificate, private key and metadata in different files
// 	// as web servers would not be able to work with a combined file.
// 	err := s.WriteFile(domain, certExt, certRes.Certificate)
// 	if err != nil {
// 		log.Fatalf("Unable to save Certificate for domain %s\n\t%v", domain, err)
// 	}

// 	if certRes.IssuerCertificate != nil {
// 		err = s.WriteFile(domain, issuerExt, certRes.IssuerCertificate)
// 		if err != nil {
// 			log.Fatalf("Unable to save IssuerCertificate for domain %s\n\t%v", domain, err)
// 		}
// 	}

// 	// if we were given a CSR, we don't know the private key
// 	if certRes.PrivateKey != nil {
// 		err = s.WriteCertificateFiles(domain, certRes)
// 		if err != nil {
// 			log.Fatalf("Unable to save PrivateKey for domain %s\n\t%v", domain, err)
// 		}
// 	} else if s.pem || s.pfx {
// 		// we don't have the private key; can't write the .pem or .pfx file
// 		log.Fatalf("Unable to save PEM or PFX without private key for domain %s. Are you using a CSR?", domain)
// 	}

// 	jsonBytes, err := json.MarshalIndent(certRes, "", "\t")
// 	if err != nil {
// 		log.Fatalf("Unable to marshal CertResource for domain %s\n\t%v", domain, err)
// 	}

// 	err = s.WriteFile(domain, resourceExt, jsonBytes)
// 	if err != nil {
// 		log.Fatalf("Unable to save CertResource for domain %s\n\t%v", domain, err)
// 	}
// }

// func (s *CertificatesStorage) ReadResource(domain string) certificate.Resource {
// 	raw, err := s.ReadFile(domain, resourceExt)
// 	if err != nil {
// 		log.Fatalf("Error while loading the meta data for domain %s\n\t%v", domain, err)
// 	}

// 	var resource certificate.Resource
// 	if err = json.Unmarshal(raw, &resource); err != nil {
// 		log.Fatalf("Error while marshaling the meta data for domain %s\n\t%v", domain, err)
// 	}

// 	return resource
// }

func recipeParse(ctx context.Context) error {
	err := db.Init(ctx)
	if err != nil {
		return fmt.Errorf("db.Init: %w", err)
	}

	err = store.Init(ctx)
	if err != nil {
		return fmt.Errorf("store.Init: %w", err)
	}

	err = recipe.Parse(ctx)
	if err != nil {
		return fmt.Errorf("item.Parse: %w", err)
	}
	return nil
}
