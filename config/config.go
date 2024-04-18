package config

import (
	"context"
	"fmt"
	"os"
	"runtime"

	"github.com/jbsmith7741/toml"
	"github.com/rs/zerolog"
)

var (
	cfg Config
)

// Config represents a configuration parse
type Config struct {
	Debug               bool      `toml:"debug" desc:"tinywebeq Configuration\n\n# Debug messages are displayed. This will cause console to be more verbose, but also more informative"`
	IsDiscoveredOnly    bool      `toml:"is_discovered_only" desc:"If true, only discovered items are viewable, default false"`
	DiscoverCacheReload int       `toml:"discover_cache_reload" desc:"How long before trying to refresh a non-discovered item as discovered, lowering this value can cause a bigger tax on sql. default 640 seconds"`
	Database            Database  `toml:"database" desc:"Database configuration"`
	MemCache            MemCache  `toml:"mem_cache" desc:"Memory cache configuration"`
	FileCache           FileCache `toml:"file_cache" desc:"File cache configuration"`
	IsSpellInfoEnabled  bool      `toml:"is_spell_info_enabled" desc:"If true, spell info is enabled (similar to mq spell details), default true"`
}

type Database struct {
	Host     string `toml:"host" desc:"Database host"`
	Port     int    `toml:"port" desc:"Database port"`
	Username string `toml:"username" desc:"Database username"`
	Name     string `toml:"name" desc:"Database name"`
	Password string `toml:"password" desc:"Database password"`
}

type FileCache struct {
	IsEnabled        bool `toml:"is_file_cache_enabled" desc:"If true, file cache is enabled, default true"`
	MaxFiles         int  `toml:"max_files" desc:"Maximum number of files to keep in cache, default 1000"`
	TruncateSchedule int  `toml:"truncate_schedule" desc:"How often to truncate file cache in seconds, default 25200 seconds (7 hours)"`
	Expiration       int  `toml:"expiration" desc:"How long to keep file cache in seconds, default 21600 seconds (6 hours) "`
}

type MemCache struct {
	IsEnabled        bool `toml:"is_mem_cache_enabled" desc:"If true, memory cache is enabled, default true"`
	MaxMemory        int  `toml:"max_memory" desc:"Maximum size of memory cache in bytes, default 150000000 (150 MB)"`
	TruncateSchedule int  `toml:"truncate_schedule" desc:"How often to truncate memory cache in seconds, default 600 seconds (10 minutes)"`
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
	if cfg.Debug {
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

	return nil
}

func defaultLabel() Config {
	cfg := Config{
		Debug:               true,
		IsDiscoveredOnly:    false,
		DiscoverCacheReload: 640,
		MemCache: MemCache{
			IsEnabled:        true,
			MaxMemory:        150000000,
			TruncateSchedule: 600,
			Expiration:       300,
		},
		FileCache: FileCache{
			IsEnabled:        true,
			MaxFiles:         1000,
			TruncateSchedule: 25200,
			Expiration:       21600,
		},
		Database: Database{
			Host:     "localhost",
			Port:     3306,
			Username: "peq",
			Name:     "peq",
			Password: "peqpass",
		},
		IsSpellInfoEnabled: true,
	}

	return cfg
}

func Get() *Config {
	return &cfg
}
