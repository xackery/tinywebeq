package main

import (
	"bytes"
	"context"
	"crypto"
	"fmt"
	"io"
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
	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/image"
	"github.com/xackery/tinywebeq/item"
	"github.com/xackery/tinywebeq/npc"
	"github.com/xackery/tinywebeq/player"
	"github.com/xackery/tinywebeq/quest"
	"github.com/xackery/tinywebeq/quest/parse"
	"github.com/xackery/tinywebeq/recipe"
	"github.com/xackery/tinywebeq/site"
	"github.com/xackery/tinywebeq/spell"
	"github.com/xackery/tinywebeq/store"
	"github.com/xackery/tinywebeq/tlog"
	"github.com/xackery/tinywebeq/zone"
	"golang.org/x/crypto/acme/autocert"
)

// Version is the build version
var Version string

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

	err = site.Init()
	if err != nil {
		return fmt.Errorf("site.Init: %w", err)
	}

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

	err = item.Init()
	if err != nil {
		return fmt.Errorf("item.Init: %w", err)
	}
	err = player.Init()
	if err != nil {
		return fmt.Errorf("player.Init: %w", err)
	}
	err = spell.Init()
	if err != nil {
		return fmt.Errorf("spell.Init: %w", err)
	}
	err = npc.Init()
	if err != nil {
		return fmt.Errorf("npc.Init: %w", err)
	}

	err = quest.Init()
	if err != nil {
		return fmt.Errorf("quest.Init: %w", err)
	}

	err = zone.Init()
	if err != nil {
		return fmt.Errorf("zone.Init: %w", err)
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

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	mux.HandleFunc("/item/view/", item.View)
	mux.HandleFunc("/item/search", item.Search)
	mux.HandleFunc("/item/preview.png", item.PreviewImage)
	mux.HandleFunc("/player/view/", player.View)
	mux.HandleFunc("/spell/view", spell.View)
	mux.HandleFunc("/spell/search", spell.Search)
	mux.HandleFunc("/spell/preview.png", spell.PreviewImage)
	mux.HandleFunc("/npc/view/", npc.View)
	mux.HandleFunc("/npc/peek/", npc.Peek)
	mux.HandleFunc("/npc/search", npc.Search)
	mux.HandleFunc("/npc/preview.png", npc.PreviewImage)
	mux.HandleFunc("/quest/view/", quest.View)
	mux.HandleFunc("/quest/search", quest.Search)
	mux.HandleFunc("/quest/preview.png", quest.PreviewImage)
	// mux.HandleFunc("/recipe/view/", recipe.View)
	// mux.HandleFunc("/recipe/search", recipe.Search)
	// mux.HandleFunc("/recipe/preview.png", recipe.PreviewImage)
	mux.HandleFunc("/zone/view/", zone.View)
	mux.HandleFunc("/zone/search", zone.Search)
	mux.HandleFunc("/zone/preview.png", zone.PreviewImage)
	mux.HandleFunc("/css/style.css", func(w http.ResponseWriter, r *http.Request) {
		fi, err := site.TemplateFS().Open("style.css")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer fi.Close()
		data, err := io.ReadAll(fi)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		http.ServeContent(w, r, "style.css", time.Now(), bytes.NewReader(data))
	})

	mux.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))).ServeHTTP(w, r)
	})

	server.Handler = mux

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
func (u LetsEncryptUser) GetRegistration() *registration.Resource {
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
		return letsencryptRenew()
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

	err = parse.Parse(ctx, config.Get().Quest.ActiveConcurrency)
	if err != nil {
		return fmt.Errorf("questParse: %w", err)
	}
	return nil
}

func letsencryptRenew() error {
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
