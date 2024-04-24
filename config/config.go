package config

import (
	"context"
	"fmt"
	"os"
	"runtime"

	"github.com/jbsmith7741/toml"
	"github.com/rs/zerolog"
	"github.com/xackery/tinywebeq/tlog"
)

var (
	cfg Config
)

// Config represents a configuration parse
type Config struct {
	IsDebugEnabled   bool        `toml:"is_debug_enabled" desc:"Default true, enables verbose logging"`
	MaxLevel         int         `toml:"max_level" desc:"Default 65, Maximum level a character can obtain. Used by spell.playeronlyspells, spell.info, item.is_search_only_player_obtainable"`
	CurrentExpansion int         `toml:"current_expansion" desc:"Default 0 (Classic), Current expansion for the server. Filters out spells that are not available to players."`
	Server           Server      `toml:"server" desc:"Web Server configuration"`
	Database         Database    `toml:"database" desc:"Database configuration"`
	Item             Item        `toml:"item" desc:"Item configuration"`
	Spell            Spell       `toml:"spell" desc:"Spell configuration"`
	Npc              Npc         `toml:"npc" desc:"NPC configuration"`
	MemCache         MemCache    `toml:"mem_cache" desc:"Memory cache configuration"`
	SqliteCache      SqliteCache `toml:"sqlite_cache" desc:"Sqlite cache configuration"`
	FileCache        FileCache   `toml:"file_cache" desc:"File cache configuration"`
}

type Server struct {
	Host        string      `toml:"host" desc:"Default localhost, what IP to bind to, can also have it empty to bind all"`
	Port        int         `toml:"port" desc:"Default 8080, what port to bind to, if you're not using a proxy and use TLS set to 443"`
	LetsEncrypt LetsEncrypt `toml:"lets_encrypt" desc:"LetsEncrypt configuration"`
}

type LetsEncrypt struct {
	IsEnabled bool     `toml:"is_enabled" desc:"Default false, WARNING: To use LetsEncrypt, you need to expose ports 5001 and 5002 for challenge requests. You also need to not use wildcard domains"`
	Email     string   `toml:"email" desc:"Default example@email.com, what email to register for lets encrypt"`
	CertPath  string   `toml:"cert" desc:"Default certs/, where to store letsencrypt certs"`
	Domains   []string `toml:"domains" desc:"Default example.com,list of domains to register for letsencrypt, do not use wildcard domains"`
	IsProd    bool     `toml:"is_prod" desc:"Default false, uses staging to ensure letsencrypt settings are good. Once you succeed, flip to true to use production certs"`
}

type Database struct {
	Host     string `toml:"host" desc:"Default localhost, where to connect for database"`
	Port     int    `toml:"port" desc:"Default 3306, what port to connect to for database"`
	Username string `toml:"username" desc:"Default peq, Database username"`
	Name     string `toml:"name" desc:"Default peq, Database name"`
	Password string `toml:"password" desc:"Default peqpass, Database password"`
}

type Item struct {
	IsEnabled           bool        `toml:"is_enabled" desc:"Default true, enables item endpoints"`
	IsCacheEnabled      bool        `toml:"is_cache_enabled" desc:"Default true, enables item caching"`
	IsDiscoveredOnly    bool        `toml:"is_discovered_only" desc:"Default false, only allow discovered items to be viewed"`
	DiscoverCacheReload int         `toml:"discover_cache_reload" desc:"Default 640, in seconds how long before trying to refresh a cached non-discovered item as discovered, lowering this value can cause a bigger tax on sql"`
	Search              ItemSearch  `toml:"search" desc:"Item search configuration"`
	Preview             ItemPreview `toml:"preview" desc:"Item preview configuration"`
}

type ItemSearch struct {
	IsEnabled                bool `toml:"is_enabled" desc:"Default false, makes a search box appear in the item view for finding other items"`
	IsBleveEnabled           bool `toml:"is_bleve_enabled" desc:"Default false, costs ~200MB of memory to do names searches in memory with bleve indexing"`
	IsOnlyPlayableObtainable bool `toml:"is_only_player_obtainable" desc:"Default false, only items determined player obtainable come up in search"`
}

type ItemPreview struct {
	BGColor    string `toml:"bg_color" desc:"Default #313338, background color for item preview"`
	FGColor    string `toml:"fg_color" desc:"Default #DBDEE1, foreground text color for item preview"`
	FontNormal string `toml:"font" desc:"Default goregular.ttf, if changed place a .ttf file same path as binary"`
	FontBold   string `toml:"font_bold" desc:"Default gobold.ttf, if changed place a .ttf file same path as binary"`
}

type Spell struct {
	IsEnabled          bool         `toml:"is_enabled" desc:"Default true, spell endpoints are enabled"`
	IsCacheEnabled     bool         `toml:"is_cache_enabled" desc:"Default true, spell caching is enabled"`
	IsSpellInfoEnabled bool         `toml:"is_spell_info_enabled" desc:"Default true, gives spell effect detailed info (based to mq spell details)"`
	Search             SpellSearch  `toml:"search" desc:"Spell search configuration"`
	Preview            SpellPreview `toml:"preview" desc:"Spell preview configuration"`
}

type SpellSearch struct {
	IsEnabled          bool `toml:"is_enabled" desc:"Default false, makes a search box appear in the spell view for finding other spells"`
	IsBleveEnabled     bool `toml:"is_bleve_enabled" desc:"Default false, when enabled it costs ~200MB of memory to do names searches with no DB hit"`
	IsOnlyPlayerSpells bool `toml:"is_only_player_spells" desc:"Default false, only spells determined player obtainable come up in search"`
}

type SpellPreview struct {
	BGColor    string `toml:"bg_color" desc:"Default #313338, background color for spell preview"`
	FGColor    string `toml:"fg_color" desc:"Default #DBDEE1, foreground text color for spell preview"`
	FontNormal string `toml:"font" desc:"Default goregular.ttf, if changed place a .ttf file same path as binary"`
	FontBold   string `toml:"font_bold" desc:"Default gobold.ttf, if changed place a .ttf file same path as binary"`
}

type Npc struct {
	IsEnabled      bool       `toml:"is_enabled" desc:"Default true, enables npc endpoints"`
	IsCacheEnabled bool       `toml:"is_cache_enabled" desc:"Default true, enables npc caching"`
	Search         NpcSearch  `toml:"search" desc:"Npc search configuration"`
	Preview        NpcPreview `toml:"preview" desc:"Npc preview configuration"`
}

type NpcSearch struct {
	IsEnabled      bool `toml:"is_enabled" desc:"Default false, makes a search box appear in the npc view for finding other npcs"`
	IsBleveEnabled bool `toml:"is_bleve_enabled" desc:"Default false, costs ~200MB of memory to do names searches in memory with bleve indexing"`
}

type NpcPreview struct {
	BGColor    string `toml:"bg_color" desc:"Default #313338, background color for npc preview"`
	FGColor    string `toml:"fg_color" desc:"Default #DBDEE1, foreground text color for npc preview"`
	FontNormal string `toml:"font" desc:"Default goregular.ttf, if changed place a .ttf file same path as binary"`
	FontBold   string `toml:"font_bold" desc:"Default gobold.ttf, if changed place a .ttf file same path as binary"`
}

type FileCache struct {
	IsEnabled        bool `toml:"is_enabled" desc:"Default false, when a page is requested it's data is cached to the cache subfolder"`
	MaxFiles         int  `toml:"max_files" desc:"Default 1000, maximum number of files to keep in cache"`
	TruncateSchedule int  `toml:"truncate_schedule" desc:"Default 25200, (25200 is seconds = 7 hours), when this hits a scheduler fires that reviews file cache and truncates expired items, freeing up memory"`
	Expiration       int  `toml:"expiration" desc:"How long to keep file cache in seconds, default 21600 seconds (6 hours) "`
}

type SqliteCache struct {
	IsEnabled        bool `toml:"is_enabled" desc:"Default true, when a page is requested it's data is cached to the sqlite cache"`
	TruncateSchedule int  `toml:"truncate_schedule" desc:"Default 25200, (25200 is seconds = 7 hours), when this hits a scheduler fires that reviews file cache and truncates expired items, freeing up memory"`
	Expiration       int  `toml:"expiration" desc:"How long to keep file cache in seconds, default 21600 seconds (6 hours) "`
}

type MemCache struct {
	IsEnabled        bool `toml:"is_enabled" desc:"If true, memory cache is enabled, default true"`
	MaxMemory        int  `toml:"max_memory" desc:"Default 150000000, (150000000 bytes = 150 MB) how much maximum memory should be used for caching page data (this does not include bleve search memory)"`
	TruncateSchedule int  `toml:"truncate_schedule" desc:"Default 600, (600s = 10 minutes), when this hits a scheduler fires that reviews memory cache and truncates expired items, freeing up memory"`
	Expiration       int  `toml:"expiration" desc:"How long to keep memory cache in seconds, default 300 seconds (5 minutes)"`
}

// NewConfig creates a new configuration
func NewConfig(ctx context.Context) (*Config, error) {
	var f *os.File
	cfg = Config{}
	path := "tinywebeq.conf"

	isNewConfig := false
	fi, err := os.Stat(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("config info: %w", err)
		}
		f, err = os.Create(path)
		if err != nil {
			return nil, fmt.Errorf("create tinywebeq.conf: %w", err)
		}
		fi, err = os.Stat(path)
		if err != nil {
			return nil, fmt.Errorf("new config info: %w", err)
		}
		isNewConfig = true
	}
	if !isNewConfig {
		f, err = os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("open config: %w", err)
		}
	}

	defer f.Close()
	if fi.IsDir() {
		return nil, fmt.Errorf("tinywebeq.conf is a directory, should be a file")
	}

	if isNewConfig {
		enc := toml.NewEncoder(f)
		enc.Encode(defaultLabel())

		fmt.Println("a new tinywebeq.conf file was created. Please open this file and configure tinywebeq, then run it again.")
		if runtime.GOOS == "windows" {
			option := ""
			fmt.Println("press a key then enter to exit.")
			fmt.Scan(&option)
		}
		os.Exit(0)
	}

	_, err = toml.DecodeReader(f, &cfg)
	if err != nil {
		return nil, fmt.Errorf("decode tinywebeq.conf: %w", err)
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if cfg.IsDebugEnabled {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	err = cfg.Verify()
	if err != nil {
		return nil, fmt.Errorf("verify: %w", err)
	}

	return &cfg, nil
}

// Verify returns an error if configuration appears off
func (c *Config) Verify() error {
	if c.SqliteCache.Expiration < 1 {
		tlog.Warnf("SqliteCache.Expiration is unset, setting to 21600")
		c.SqliteCache.Expiration = 21600
	}
	if c.SqliteCache.TruncateSchedule < 1 {
		tlog.Warnf("SqliteCache.TruncateSchedule is unset, setting to 25200")
		c.SqliteCache.TruncateSchedule = 25200
	}
	if c.FileCache.Expiration < 1 {
		tlog.Warnf("FileCache.Expiration is unset, setting to 21600")
		c.FileCache.Expiration = 21600
	}
	if c.FileCache.TruncateSchedule < 1 {
		tlog.Warnf("FileCache.TruncateSchedule is unset, setting to 25200")
		c.FileCache.TruncateSchedule = 25200
	}
	if c.MemCache.Expiration < 1 {
		tlog.Warnf("MemCache.Expiration is unset, setting to 300")
		c.MemCache.Expiration = 300
	}
	if c.MemCache.TruncateSchedule < 1 {
		tlog.Warnf("MemCache.TruncateSchedule is unset, setting to 600")
		c.MemCache.TruncateSchedule = 600
	}
	if c.MemCache.MaxMemory < 1 {
		tlog.Warnf("MemCache.MaxMemory is unset, setting to 150000000")
		c.MemCache.MaxMemory = 150000000
	}
	if c.Server.Port < 1 {
		tlog.Warnf("Server.Port is unset, setting to 8080")
		c.Server.Port = 8080
	}
	if c.Server.Host == "" {
		tlog.Warnf("Server.Host is unset, setting to localhost")
		c.Server.Host = "localhost"
	}
	if c.Database.Host == "" {
		tlog.Warnf("Database.Host is unset, setting to localhost")
		c.Database.Host = "localhost"
	}

	return nil
}

func defaultLabel() Config {
	cfg := Config{
		IsDebugEnabled: true,
		MaxLevel:       65,
		Server: Server{
			Host: "localhost",
			Port: 8080,
			LetsEncrypt: LetsEncrypt{
				IsEnabled: false,
				Email:     "example@email.com",
				CertPath:  "./certs",
				Domains:   []string{"example.com", "www.example.com"},
				IsProd:    false,
			},
		},
		MemCache: MemCache{
			IsEnabled:        true,
			MaxMemory:        150000000,
			TruncateSchedule: 600,
			Expiration:       300,
		},
		SqliteCache: SqliteCache{
			IsEnabled:        true,
			TruncateSchedule: 25200,
			Expiration:       21600,
		},
		FileCache: FileCache{
			IsEnabled:        false,
			MaxFiles:         1000,
			TruncateSchedule: 25200,
			Expiration:       21600,
		},
		Item: Item{
			IsEnabled:           true,
			IsCacheEnabled:      true,
			IsDiscoveredOnly:    false,
			DiscoverCacheReload: 640,
			Search: ItemSearch{
				IsEnabled:                false,
				IsBleveEnabled:           false,
				IsOnlyPlayableObtainable: false,
			},
			Preview: ItemPreview{
				BGColor:    "#313338",
				FGColor:    "#DBDEE1",
				FontNormal: "goregular.ttf",
				FontBold:   "gobold.ttf",
			},
		},
		Spell: Spell{
			IsEnabled:          true,
			IsCacheEnabled:     true,
			IsSpellInfoEnabled: true,
			Search: SpellSearch{
				IsEnabled:          false,
				IsBleveEnabled:     false,
				IsOnlyPlayerSpells: false,
			},
			Preview: SpellPreview{
				BGColor:    "#313338",
				FGColor:    "#DBDEE1",
				FontNormal: "goregular.ttf",
				FontBold:   "gobold.ttf",
			},
		},
		Database: Database{
			Host:     "localhost",
			Port:     3306,
			Username: "peq",
			Name:     "peq",
			Password: "peqpass",
		},
	}

	return cfg
}

func Get() *Config {
	return &cfg
}
