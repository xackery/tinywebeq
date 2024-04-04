package cache

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/tlog"
)

var (
	mu            sync.RWMutex
	memCache      map[string]*cacheEntry = make(map[string]*cacheEntry)
	memCacheSize  int
	fileCache     map[string]int = make(map[string]int)
	fileCacheSize int
)

type cacheEntry struct {
	expiration int64
	data       []byte
}

func init() {
	readFileCacheIndex()
	go maintain()
}

// Write writes data to cache
func Write(path string, data []byte) error {
	mu.Lock()
	defer mu.Unlock()

	err := writeMemoryCache(path, data)
	if err != nil {
		return fmt.Errorf("write memory cache: %w", err)
	}

	err = writeFileCache(path, data)
	if err != nil {
		return fmt.Errorf("write file cache: %w", err)
	}

	return nil
}

func writeMemoryCache(path string, data []byte) error {
	if !config.Get().MemCache.IsEnabled {
		return nil
	}

	oldCache, ok := memCache[path]
	if ok {
		memCacheSize -= len(oldCache.data)
		memCache[path] = &cacheEntry{
			expiration: time.Now().Add(time.Minute * time.Duration(config.Get().MemCache.Expiration)).Unix(),
			data:       data,
		}
		memCacheSize += len(data)
		tlog.Debugf("memcache overwrite: %s, expiration: %d", path, memCache[path].expiration)
		return nil
	}

	if memCacheSize+len(data) > config.Get().MemCache.Max {
		tlog.Debugf("memcache full, skipping: %s", path)
		return nil
	}

	memCache[path] = &cacheEntry{
		expiration: time.Now().Add(time.Minute * time.Duration(config.Get().MemCache.Expiration)).Unix(),
		data:       data,
	}
	memCacheSize += len(data)
	tlog.Debugf("memcache write: %s, expiration: %d (%d total size)", path, memCache[path].expiration, memCacheSize)
	return nil
}

func writeFileCache(path string, data []byte) error {
	if !config.Get().FileCache.IsEnabled {
		return nil
	}

	if fileCacheSize+len(data) > config.Get().FileCache.Max {
		tlog.Debugf("filecache full, skipping: %s", path)
		return nil
	}

	fileCache[path] = int(time.Now().Add(time.Minute * time.Duration(config.Get().FileCache.Expiration)).Unix())
	fileCacheSize += len(data)

	err := writeFileCacheIndex()
	if err != nil {
		tlog.Errorf("write file cache index: %v", err)
	}

	basePath := filepath.Dir(path)
	err = os.MkdirAll(basePath, 0755)
	if err != nil {
		return fmt.Errorf("make %s: %w", basePath, err)
	}

	cw, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create %s: %w", path, err)
	}
	defer cw.Close()
	_, err = cw.Write(data)
	if err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	tlog.Debugf("filecache write: %s", path)
	return nil
}

// Read reads data from cache
func Read(path string) (bool, []byte) {
	mu.Lock()
	defer mu.Unlock()
	if !config.Get().FileCache.IsEnabled && !config.Get().MemCache.IsEnabled {
		return false, nil
	}
	if config.Get().MemCache.IsEnabled {
		entry, ok := memCache[path]
		if ok {
			if entry.expiration > time.Now().Unix() {
				tlog.Debugf("memcache read: %s, expiration: %s", path, entry.expiration)
				return true, entry.data
			}
			tlog.Debugf("memcache expired: %s", path)
			delete(memCache, path)
			memCacheSize -= len(entry.data)
		}
	}

	if config.Get().FileCache.IsEnabled {
		expiration, ok := fileCache[path]
		if ok {
			if expiration < int(time.Now().Unix()) {
				tlog.Debugf("filecache expired: %s", path)
				delete(fileCache, path)
				err := os.Remove("cache/" + path)
				if err != nil {
					tlog.Errorf("remove file: %v", err)
				}
				err = writeFileCacheIndex()
				if err != nil {
					tlog.Errorf("write file cache index: %v", err)
				}
				return false, nil
			}

			tlog.Debugf("filecache read: %s, expiration: %d", path, expiration)
			data, err := os.ReadFile("cache/" + path)
			if err != nil {
				tlog.Errorf("read file: %v", err)
				return false, nil
			}

			err = writeMemoryCache(path, data)
			if err != nil {
				tlog.Errorf("write memory cache: %v", err)
			}
			return true, data
		}
	}

	tlog.Debugf("cache miss: %s", path)
	return false, nil
}

func maintain() {

	time.Sleep(5 * time.Second)

	tickerMemCache := time.NewTicker(time.Duration(config.Get().MemCache.TruncateSchedule) * time.Second)
	defer tickerMemCache.Stop()
	tickerFileCache := time.NewTicker(time.Duration(config.Get().FileCache.TruncateSchedule) * time.Second)
	defer tickerFileCache.Stop()

	for {
		select {
		case <-tickerMemCache.C:
			truncateMemCache()
		case <-tickerFileCache.C:
			truncateFileCache()
		}
	}
}

func truncateMemCache() {
	if !config.Get().MemCache.IsEnabled {
		return
	}

	time.Sleep(time.Duration(config.Get().MemCache.TruncateSchedule))
	tlog.Debug("memcache truncate schedule running...")
	start := time.Now()
	mu.Lock()
	for path, entry := range memCache {
		if entry.expiration > time.Now().Unix() {
			continue
		}
		tlog.Debugf("memcache expired: %s", path)
		memCacheSize -= len(entry.data)
		delete(memCache, path)
	}
	mu.Unlock()
	tlog.Debugf("memcache truncate schedule complete in %s", time.Since(start))
}

func truncateFileCache() {
	if !config.Get().FileCache.IsEnabled {
		return
	}
	time.Sleep(time.Duration(config.Get().FileCache.TruncateSchedule))
	tlog.Debug("filecache truncate schedule running...")
	start := time.Now()
	mu.Lock()
	for path, expiration := range fileCache {
		if expiration > int(time.Now().Unix()) {
			continue
		}
		tlog.Debugf("filecache expired: %s", path)
		fileCacheSize -= 0
		delete(fileCache, path)
		err := os.Remove("cache/" + path)
		if err != nil {
			tlog.Errorf("remove file: %v", err)
		}
	}
	err := writeFileCacheIndex()
	if err != nil {
		tlog.Errorf("write file cache index: %v", err)
	}

	mu.Unlock()
	tlog.Debugf("filecache truncate schedule complete in %s", time.Since(start))
}

func writeFileCacheIndex() error {
	err := os.MkdirAll("cache", 0755)
	if err != nil {
		return fmt.Errorf("make cache: %w", err)
	}
	w, err := os.Create("cache/index.json")
	if err != nil {
		return fmt.Errorf("write cache index: %w", err)
	}
	defer w.Close()

	err = json.NewEncoder(w).Encode(fileCache)
	if err != nil {
		return fmt.Errorf("encode cache index: %w", err)
	}
	return nil
}

func readFileCacheIndex() error {
	r, err := os.Open("cache/index.json")
	if err != nil {
		return fmt.Errorf("read cache index: %w", err)
	}
	defer r.Close()

	err = json.NewDecoder(r).Decode(&fileCache)
	if err != nil {
		return fmt.Errorf("decode cache index: %w", err)
	}
	return nil
}
