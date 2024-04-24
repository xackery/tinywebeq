package cache

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
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
	memCache      map[string]*cacheEntry = make(map[string]*cacheEntry)
	memCacheSize  int
	fileCache     map[string]int = make(map[string]int)
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
		err = writeMemoryCache(ctx, path, data)
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

func writeMemoryCache(ctx context.Context, path string, data model.CacheIdentifier) error {
	if !config.Get().MemCache.IsEnabled {
		return nil
	}

	size := 4000
	_, ok := memCache[path]
	if ok {
		memCacheSize -= size
		memCache[path] = &cacheEntry{
			expiration: time.Now().Add(time.Minute * time.Duration(config.Get().MemCache.Expiration)).Unix(),
			data:       data,
		}
		memCacheSize += size
		tlog.Debugf("Memcache overwrite: %s, expiration: %d", path, memCache[path].expiration)
		return nil
	}

	if memCacheSize+size > config.Get().MemCache.MaxMemory {
		tlog.Debugf("Memcache full, skipping: %s", path)
		return nil
	}

	memCache[path] = &cacheEntry{
		expiration: time.Now().Add(time.Minute * time.Duration(config.Get().MemCache.Expiration)).Unix(),
		data:       data,
	}
	memCacheSize += size
	tlog.Debugf("Memcache write: %s, expiration: %d (%d total size)", path, memCache[path].expiration, memCacheSize)
	return nil
}

func writeFileCache(ctx context.Context, path string, data model.CacheIdentifier) error {
	if !config.Get().FileCache.IsEnabled {
		return nil
	}

	if len(fileCache) > config.Get().FileCache.MaxFiles {
		tlog.Debugf("Filecache full, skipping: %s", path)
		return nil
	}

	fileCache[path] = int(time.Now().Add(time.Minute * time.Duration(config.Get().FileCache.Expiration)).Unix())

	err := writeFileCacheIndex()
	if err != nil {
		tlog.Errorf("Write file cache index: %v", err)
	}

	basePath := filepath.Dir(path)
	err = os.MkdirAll("cache/"+basePath, 0755)
	if err != nil {
		return fmt.Errorf("make %s: %w", basePath, err)
	}

	cw, err := os.Create("cache/" + path)
	if err != nil {
		return fmt.Errorf("create %s: %w", path, err)
	}
	defer cw.Close()

	err = yaml.NewEncoder(cw).Encode(data)
	if err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	tlog.Debugf("Filecache write: %s", path)
	return nil
}

// Read reads data from cache
func Read(ctx context.Context, path string) (model.CacheIdentifier, bool) {
	mu.Lock()
	defer mu.Unlock()
	size := 4000
	if config.Get().MemCache.IsEnabled {
		entry, ok := memCache[path]
		if ok {
			if entry.expiration > time.Now().Unix() {
				tlog.Debugf("Memcache read: %s, expiration: %d", path, entry.expiration)
				return entry.data, true
			}
			tlog.Debugf("Memcache expired: %s", path)
			delete(memCache, path)
			memCacheSize -= size
		}
	}

	if config.Get().SqliteCache.IsEnabled {
		data, ok := readSqliteCache(path)
		if ok {
			tlog.Debugf("SqliteCache read: %s, expiration: %d", path, data.Expiration())
			return data, true
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
				return nil, false
			}

			tlog.Debugf("Filecache read: %s, expiration: %d", path, expiration)
			r, err := os.Open("cache/" + path)
			if err != nil {
				tlog.Warnf("Open file (skipping cache): %v", err)
				return nil, false
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
				return nil, false
			}

			err = writeMemoryCache(ctx, path, cacheData)
			if err != nil {
				tlog.Warnf("Write memory cache: %v", err)
			}
			return cacheData, true
		}
	}

	tlog.Debugf("Cache miss: %s", path)
	return nil, false
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

func truncateMemCache() {
	if !config.Get().MemCache.IsEnabled {
		return
	}

	size := 4000
	time.Sleep(time.Duration(config.Get().MemCache.TruncateSchedule))
	tlog.Debugf("Memcache truncate schedule running...")
	start := time.Now()
	mu.Lock()
	for path, entry := range memCache {
		if entry.expiration > time.Now().Unix() {
			continue
		}
		tlog.Debugf("Memcache expired: %s", path)
		memCacheSize -= size
		delete(memCache, path)
	}
	mu.Unlock()
	tlog.Debugf("Memcache truncate schedule complete in %s", time.Since(start))
}

func truncateFileCache() {
	if !config.Get().FileCache.IsEnabled {
		return
	}
	time.Sleep(time.Duration(config.Get().FileCache.TruncateSchedule))
	tlog.Debugf("Filecache truncate schedule running...")
	start := time.Now()
	mu.Lock()
	for path, expiration := range fileCache {
		if expiration > int(time.Now().Unix()) {
			continue
		}
		tlog.Debugf("Filecache expired: %s", path)
		delete(fileCache, path)
		err := os.Remove("cache/" + path)
		if err != nil {
			tlog.Errorf("Remove file: %v", err)
		}
	}
	err := writeFileCacheIndex()
	if err != nil {
		tlog.Errorf("Write file cache index: %v", err)
	}

	mu.Unlock()
	tlog.Debugf("Filecache truncate schedule complete in %s", time.Since(start))
}

func writeFileCacheIndex() error {
	err := os.MkdirAll("cache", 0755)
	if err != nil {
		return fmt.Errorf("make cache: %w", err)
	}
	w, err := os.Create("cache/index.yaml")
	if err != nil {
		return fmt.Errorf("write cache index: %w", err)
	}
	defer w.Close()

	err = yaml.NewEncoder(w).Encode(fileCache)
	if err != nil {
		return fmt.Errorf("encode cache index: %w", err)
	}
	return nil
}

func readFileCacheIndex() error {
	r, err := os.Open("cache/index.yaml")
	if err != nil {
		return fmt.Errorf("read cache index: %w", err)
	}
	defer r.Close()

	err = yaml.NewDecoder(r).Decode(&fileCache)
	if err != nil {
		return fmt.Errorf("decode cache index: %w", err)
	}
	return nil
}

func Close() error {
	return sqliteClose()
}
