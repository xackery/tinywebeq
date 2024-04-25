package cache

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/model"
	"github.com/xackery/tinywebeq/tlog"
	"gopkg.in/yaml.v3"
)

var (
	mu            sync.RWMutex
	isInitialized bool
	fileCache     map[string]int = make(map[string]int)
)

type cacheSource int

const (
	SourceCacheNone cacheSource = iota
	SourceCacheMemory
	SourceCacheFile
	SourceCacheSqlite
)

type cacheEntry struct {
	expiration int64
	data       model.CacheIdentifier
}

func Init(ctx context.Context, isCacheFlush bool) error {
	if isInitialized {
		return nil
	}
	isInitialized = true
	if isCacheFlush {
		tlog.Debugf("Flushing cache...")
		err := os.RemoveAll("cache")
		if err != nil {
			return fmt.Errorf("remove cache: %w", err)
		}
		go maintain()
		return nil
	}

	err := dbliteInit(ctx)
	if err != nil {
		return fmt.Errorf("dblite init: %w", err)
	}

	readFileCacheIndex()
	go maintain()
	return nil
}

// Write writes data to cache
func Write(ctx context.Context, path string, data model.CacheIdentifier) error {
	mu.Lock()
	defer mu.Unlock()

	var err error
	if config.Get().MemCache.IsEnabled {
		err = WriteMemoryCache(ctx, path, data)
		if err != nil {
			return fmt.Errorf("write memory cache: %w", err)
		}
	}

	if config.Get().SqliteCache.IsEnabled {
		err = writeSqliteCache(ctx, path, data)
		if err != nil {
			return fmt.Errorf("write sqlite cache: %w", err)
		}
	}

	if config.Get().FileCache.IsEnabled {
		err = writeFileCache(ctx, path, data)
		if err != nil {
			return fmt.Errorf("write file cache: %w", err)
		}
	}

	return nil
}

// Read reads data from cache
func Read(ctx context.Context, path string) (model.CacheIdentifier, cacheSource, bool) {
	mu.Lock()
	defer mu.Unlock()
	size := 4000
	if config.Get().MemCache.IsEnabled {
		entry, ok := memCache[path]
		if ok {
			if entry.expiration > time.Now().Unix() {
				//tlog.Debugf("Memcache read: %s, expiration: %d", path, entry.expiration)
				return entry.data, SourceCacheMemory, true
			}
			tlog.Debugf("Memcache expired: %s", path)
			delete(memCache, path)
			memCacheSize -= size
		}
	}

	if config.Get().SqliteCache.IsEnabled {
		data, ok := readSqliteCache(path)
		if ok {
			//tlog.Debugf("SqliteCache read: %s, expiration: %d", path, data.Expiration())
			return data, SourceCacheSqlite, true
		}
	}

	if config.Get().FileCache.IsEnabled {
		expiration, ok := fileCache[path]
		if ok {
			if expiration < int(time.Now().Unix()) {
				tlog.Debugf("Filecache expired: %s", path)
				delete(fileCache, path)
				err := os.Remove("cache/" + path)
				if err != nil {
					tlog.Warnf("Remove file (skipping cache): %v", err)
				}
				err = writeFileCacheIndex()
				if err != nil {
					tlog.Warnf("Write file cache index (skipping cache): %v", err)
				}
				return nil, SourceCacheFile, false
			}

			//tlog.Debugf("Filecache read: %s, expiration: %d", path, expiration)
			r, err := os.Open("cache/" + path)
			if err != nil {
				tlog.Warnf("Open file (skipping cache): %v", err)
				return nil, SourceCacheFile, false
			}
			defer r.Close()

			var cacheData model.CacheIdentifier
			err = nil
			if strings.HasPrefix(path, "item/") {
				cacheData = &model.Item{}
				err = yaml.NewDecoder(r).Decode(cacheData)
			}
			if err != nil {
				tlog.Warnf("Decode (skipping cache): %v", err)
				return nil, SourceCacheFile, false
			}
			return cacheData, SourceCacheFile, true
		}
	}

	tlog.Debugf("Cache miss: %s", path)
	return nil, SourceCacheNone, false
}

func maintain() {

	time.Sleep(5 * time.Second)

	tickerMemCache := time.NewTicker(time.Duration(config.Get().MemCache.TruncateSchedule) * time.Second)
	defer tickerMemCache.Stop()
	tickerFileCache := time.NewTicker(time.Duration(config.Get().FileCache.TruncateSchedule) * time.Second)
	defer tickerFileCache.Stop()
	tickerSqliteCache := time.NewTicker(time.Duration(config.Get().SqliteCache.TruncateSchedule) * time.Second)
	defer tickerSqliteCache.Stop()

	for {
		select {
		case <-tickerMemCache.C:
			truncateMemCache()
		case <-tickerFileCache.C:
			truncateFileCache()
		case <-tickerSqliteCache.C:
			truncateSqliteCache()
		}
	}
}

func Close() error {
	return sqliteClose()
}
