package config

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"sync"

	"github.com/jbsmith7741/toml"
	"github.com/rs/zerolog"
	"github.com/xackery/tinywebeq/tlog"
)

var (
	cfg Config
	mu  sync.RWMutex
)

// Config represents a configuration parse
type Config struct {
	IsDebugEnabled   bool     `toml:"is_debug_enabled" desc:"Default true, enables verbose logging"`
	MaxLevel         int32    `toml:"max_level" desc:"Default 65, Maximum level a character can obtain. Used by spell.playeronlyspells, spell.info, item.is_search_only_player_obtainable"`
	CurrentExpansion int      `toml:"current_expansion" desc:"Default 0 (Classic), Current expansion for the server. Filters out spells that are not available to players."`
	Site             Site     `toml:"server" desc:"Web Site configuration"`
	Database         Database `toml:"database" desc:"Database configuration"`
	Item             Item     `toml:"item" desc:"Item configuration"`
	Spell            Spell    `toml:"spell" desc:"Spell configuration"`
	Npc              Npc      `toml:"npc" desc:"NPC configuration"`
	Quest            Quest    `toml:"quest" desc:"Quest parser configuration"`
	Recipe           Recipe   `toml:"recipe" desc:"Recipe parser configuration"`
	Player           Player   `toml:"player" desc:"Player parser configuration"`
	Zone             Zone     `toml:"zone" desc:"Zone parser configuration"`
	MemCache         MemCache `toml:"mem_cache" desc:"Memory cache configuration"`
}

type Site struct {
	Host        string      `toml:"host" desc:"Default localhost, what IP to bind to, can also have it empty to bind all"`
	Port        int         `toml:"port" desc:"Default 8080, what port to bind to, if you're not using a proxy and use TLS set to 443"`
	LetsEncrypt LetsEncrypt `toml:"lets_encrypt" desc:"LetsEncrypt configuration"`
	BaseURL     string      `toml:"base_url" desc:"Default http://localhost:8080, base url for the site, e.g. https://foo.com"`
	GoogleTag   string      `toml:"google_tag" desc:"Default empty, google tag manager id"`
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
	IsMemorySearchEnabled    bool `toml:"is_memory_search_enabled" desc:"Default false, enables memory search for items"`
	IsOnlyPlayableObtainable bool `toml:"is_only_player_obtainable" desc:"Default false, only items determined player obtainable come up in search"`
}

type ItemPreview struct {
	IsEnabled  bool   `toml:"is_enabled" desc:"Default true, enables item preview"`
	BGColor    string `toml:"bg_color" desc:"Default #313338, background color for item preview"`
	FGColor    string `toml:"fg_color" desc:"Default #DBDEE1, foreground text color for item preview"`
	FontNormal string `toml:"font" desc:"Default goregular.ttf, if changed place a .ttf file same path as binary"`
	FontBold   string `toml:"font_bold" desc:"Default gobold.ttf, if changed place a .ttf file same path as binary"`
}

type Player struct {
	IsEnabled      bool          `toml:"is_enabled" desc:"Default true, enables player endpoints"`
	IsCacheEnabled bool          `toml:"is_cache_enabled" desc:"Default true, enables player caching"`
	Search         PlayerSearch  `toml:"search" desc:"Player search configuration"`
	Preview        PlayerPreview `toml:"preview" desc:"Player preview configuration"`
}

type PlayerSearch struct {
	IsEnabled             bool `toml:"is_enabled" desc:"Default false, makes a search box appear in the player view for finding other players"`
	IsMemorySearchEnabled bool `toml:"is_memory_search_enabled" desc:"Default false, enables memory search for players"`
}

type PlayerPreview struct {
	IsEnabled  bool   `toml:"is_enabled" desc:"Default true, enables player preview"`
	BGColor    string `toml:"bg_color" desc:"Default #313338, background color for player preview"`
	FGColor    string `toml:"fg_color" desc:"Default #DBDEE1, foreground text color for player preview"`
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
	IsEnabled             bool `toml:"is_enabled" desc:"Default false, makes a search box appear in the spell view for finding other spells"`
	IsMemorySearchEnabled bool `toml:"is_memory_search_enabled" desc:"Default false, enables memory search for spells"`
	IsOnlyPlayerSpells    bool `toml:"is_only_player_spells" desc:"Default false, only spells determined player obtainable come up in search"`
}

type SpellPreview struct {
	IsEnabled  bool   `toml:"is_enabled" desc:"Default true, enables spell preview"`
	BGColor    string `toml:"bg_color" desc:"Default #313338, background color for spell preview"`
	FGColor    string `toml:"fg_color" desc:"Default #DBDEE1, foreground text color for spell preview"`
	FontNormal string `toml:"font" desc:"Default goregular.ttf, if changed place a .ttf file same path as binary"`
	FontBold   string `toml:"font_bold" desc:"Default gobold.ttf, if changed place a .ttf file same path as binary"`
}

type Zone struct {
	IsEnabled bool        `toml:"is_enabled" desc:"Default true, enables zone endpoints"`
	Search    ZoneSearch  `toml:"search" desc:"Zone search configuration"`
	Preview   ZonePreview `toml:"preview" desc:"Zone preview configuration"`
}

type ZoneSearch struct {
	IsEnabled             bool `toml:"is_enabled" desc:"Default true, makes a search box appear in the zone view for finding other zones"`
	IsMemorySearchEnabled bool `toml:"is_memory_search_enabled" desc:"Default false, enables memory search for zones"`
}

type ZonePreview struct {
	IsEnabled  bool   `toml:"is_enabled" desc:"Default true, enables zone preview"`
	BGColor    string `toml:"bg_color" desc:"Default #313338, background color for zone preview"`
	FGColor    string `toml:"fg_color" desc:"Default #DBDEE1, foreground text color for zone preview"`
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
	IsEnabled             bool `toml:"is_enabled" desc:"Default false, makes a search box appear in the npc view for finding other npcs"`
	IsMemorySearchEnabled bool `toml:"is_memory_search_enabled" desc:"Default false, enables memory search for npcs"`
}

type NpcPreview struct {
	IsEnabled  bool   `toml:"is_enabled" desc:"Default true, enables npc preview"`
	BGColor    string `toml:"bg_color" desc:"Default #313338, background color for npc preview"`
	FGColor    string `toml:"fg_color" desc:"Default #DBDEE1, foreground text color for npc preview"`
	FontNormal string `toml:"font" desc:"Default goregular.ttf, if changed place a .ttf file same path as binary"`
	FontBold   string `toml:"font_bold" desc:"Default gobold.ttf, if changed place a .ttf file same path as binary"`
}

type Quest struct {
	IsEnabled                   bool         `toml:"is_enabled" desc:"Default true, enables quest features"`
	Path                        string       `toml:"path" desc:"Default quests/, where to find quest files"`
	ActiveConcurrency           int          `toml:"active_concurrency" desc:"Default 100, how many quests to process at once when the quests dedicated command is ran (this impacts connection count)"`
	IsBackgroundScanningEnabled bool         `toml:"is_background_scanning_enabled" desc:"Default false, when true, a background scanner will scan for new quests and update the cache at ScanSchedule"`
	BackgroundScanConcurrency   int          `toml:"background_scan_concurrency" desc:"Default 10, how many quests to process at once when the background scanner is running"`
	ScanSchedule                int          `toml:"scan_schedule" desc:"Default 25200, (25200 is seconds = 7 hours), when this hits a scheduler fires that reviews quest cache to rebuild it"`
	Preview                     QuestPreview `toml:"preview" desc:"Quest preview configuration"`
	Search                      QuestSearch  `toml:"search" desc:"Quest search configuration"`
}

type QuestPreview struct {
	IsEnabled  bool   `toml:"is_enabled" desc:"Default true, enables quest preview"`
	BGColor    string `toml:"bg_color" desc:"Default #313338, background color for quest preview"`
	FGColor    string `toml:"fg_color" desc:"Default #DBDEE1, foreground text color for quest preview"`
	FontNormal string `toml:"font" desc:"Default goregular.ttf, if changed place a .ttf file same path as binary"`
	FontBold   string `toml:"font_bold" desc:"Default gobold.ttf, if changed place a .ttf file same path as binary"`
}

type QuestSearch struct {
	IsEnabled             bool `toml:"is_enabled" desc:"Default false, makes a search box appear in the quest view for finding other quests"`
	IsMemorySearchEnabled bool `toml:"is_memory_search_enabled" desc:"Default false, enables memory search for quests"`
}

type Recipe struct {
	IsEnabled                   bool `toml:"is_enabled" desc:"Default true, enables recipe features"`
	ActiveConcurrency           int  `toml:"active_concurrency" desc:"Default 100, how many recipes to process at once when the recipes dedicated command is ran (this impacts connection count)"`
	IsBackgroundScanningEnabled bool `toml:"is_background_scanning_enabled" desc:"Default false, when true, a background scanner will scan for new recipes and update the cache at ScanSchedule"`
	BackgroundScanConcurrency   int  `toml:"background_scan_concurrency" desc:"Default 10, how many recipes to process at once when the background scanner is running"`
	ScanSchedule                int  `toml:"scan_schedule" desc:"Default 25200 (25200 is seconds = 7 hours), when this hits a scheduler fires that reviews recipe cache to rebuild it"`
}

type MemCache struct {
	IsEnabled        bool  `toml:"is_enabled" desc:"If true, memory cache is enabled, default true"`
	MaxMemory        int   `toml:"max_memory" desc:"Default 150000000, (150000000 bytes = 150 MB) how much maximum memory should be used for caching page data (this does not include bleve search memory)"`
	TruncateSchedule int   `toml:"truncate_schedule" desc:"Default 600, (600s = 10 minutes), when this hits a scheduler fires that reviews memory cache and truncates expired items, freeing up memory"`
	Duration         int64 `toml:"duration" desc:"How long to keep memory cache in seconds, default 300 seconds (5 minutes)"`
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

	tlog.SetLevel(zerolog.InfoLevel)
	if cfg.IsDebugEnabled {
		tlog.SetLevel(zerolog.DebugLevel)
	}

	err = cfg.Verify()
	if err != nil {
		return nil, fmt.Errorf("verify: %w", err)
	}

	return &cfg, nil
}

func NewTestConfig(ctx context.Context) (*Config, error) {
	cfg = defaultLabel()
	err := cfg.Verify()
	if err != nil {
		return nil, fmt.Errorf("verify: %w", err)
	}
	return &cfg, nil
}

// Verify returns an error if configuration appears off
func (c *Config) Verify() error {

	if c.MemCache.Duration < 1 {
		tlog.Warnf("MemCache.Expiration is unset, setting to 300")
		c.MemCache.Duration = 300
	}
	if c.MemCache.TruncateSchedule < 1 {
		tlog.Warnf("MemCache.TruncateSchedule is unset, setting to 600")
		c.MemCache.TruncateSchedule = 600
	}
	if c.MemCache.MaxMemory < 1 {
		tlog.Warnf("MemCache.MaxMemory is unset, setting to 150000000")
		c.MemCache.MaxMemory = 150000000
	}
	if c.Site.Port < 1 {
		tlog.Warnf("Server.Port is unset, setting to 8080")
		c.Site.Port = 8080
	}
	if c.Site.Host == "" {
		tlog.Warnf("Server.Host is unset, setting to localhost")
		c.Site.Host = "localhost"
	}
	if c.Site.BaseURL == "" {
		return fmt.Errorf("Server.BaseURL is unset, please set to the base url of the site")
	}
	if c.Database.Host == "" {
		tlog.Warnf("Database.Host is unset, setting to localhost")
		c.Database.Host = "localhost"
	}

	if c.Quest.ActiveConcurrency < 1 {
		tlog.Warnf("Quest.ActiveConcurrency is unset, setting to 100")
		c.Quest.ActiveConcurrency = 100
	}
	if c.Quest.BackgroundScanConcurrency < 1 {
		tlog.Warnf("Quest.BackgroundScanConcurrency is unset, setting to 10")
		c.Quest.BackgroundScanConcurrency = 10
	}

	if c.Recipe.ActiveConcurrency < 1 {
		tlog.Warnf("Recipe.ActiveConcurrency is unset, setting to 100")
		c.Recipe.ActiveConcurrency = 100
	}

	if c.Recipe.BackgroundScanConcurrency < 1 {
		tlog.Warnf("Recipe.BackgroundScanConcurrency is unset, setting to 10")
		c.Recipe.BackgroundScanConcurrency = 10
	}

	return nil
}

func defaultLabel() Config {
	cfg := Config{
		IsDebugEnabled: true,
		MaxLevel:       65,
		Site: Site{
			Host:    "localhost",
			Port:    8080,
			BaseURL: "http://localhost:8080",
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
			Duration:         300,
		},
		Item: Item{
			IsEnabled:           true,
			IsCacheEnabled:      true,
			IsDiscoveredOnly:    false,
			DiscoverCacheReload: 640,
			Search: ItemSearch{
				IsEnabled:                false,
				IsOnlyPlayableObtainable: false,
			},
			Preview: ItemPreview{
				IsEnabled:  true,
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
				IsEnabled:          true,
				IsOnlyPlayerSpells: false,
			},
			Preview: SpellPreview{
				IsEnabled:  true,
				BGColor:    "#313338",
				FGColor:    "#DBDEE1",
				FontNormal: "goregular.ttf",
				FontBold:   "gobold.ttf",
			},
		},
		Quest: Quest{
			IsEnabled:                   true,
			Path:                        "quests/",
			ScanSchedule:                25200,
			IsBackgroundScanningEnabled: false,
			BackgroundScanConcurrency:   10,
			ActiveConcurrency:           100,
			Preview: QuestPreview{
				IsEnabled:  true,
				BGColor:    "#313338",
				FGColor:    "#DBDEE1",
				FontNormal: "goregular.ttf",
				FontBold:   "gobold.ttf",
			},
			Search: QuestSearch{
				IsEnabled: false,
			},
		},
		Recipe: Recipe{
			IsEnabled:                   true,
			ScanSchedule:                25200,
			IsBackgroundScanningEnabled: false,
			BackgroundScanConcurrency:   10,
			ActiveConcurrency:           100,
		},

		Npc: Npc{
			IsEnabled:      true,
			IsCacheEnabled: true,
			Search: NpcSearch{
				IsEnabled: false,
			},
			Preview: NpcPreview{
				IsEnabled:  true,
				BGColor:    "#313338",
				FGColor:    "#DBDEE1",
				FontNormal: "goregular.ttf",
				FontBold:   "gobold.ttf",
			},
		},
		Zone: Zone{
			IsEnabled: true,
			Search: ZoneSearch{
				IsEnabled: true,
			},
			Preview: ZonePreview{
				IsEnabled:  true,
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
	mu.RLock()
	defer mu.RUnlock()
	return &cfg
}
