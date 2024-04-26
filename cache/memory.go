package cache

import (
	"context"
	"time"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/model"
	"github.com/xackery/tinywebeq/tlog"
)

var (
	memCache     map[string]*cacheEntry = make(map[string]*cacheEntry)
	memCacheSize int
)

func WriteMemoryCache(ctx context.Context, path string, data model.CacheIdentifier) error {
	if !config.Get().MemCache.IsEnabled {
		return nil
	}

	size := 4000
	_, ok := memCache[path]
	if ok {
		memCacheSize -= size
		memCache[path] = &cacheEntry{
			CacheExpiration: time.Now().Add(time.Minute * time.Duration(config.Get().MemCache.Expiration)).Unix(),
			data:            data,
		}
		memCacheSize += size
		tlog.Debugf("Memcache overwrite: %s, expiration: %d", path, memCache[path].CacheExpiration)
		return nil
	}

	if memCacheSize+size > config.Get().MemCache.MaxMemory {
		tlog.Debugf("Memcache full, skipping: %s", path)
		return nil
	}

	memCache[path] = &cacheEntry{
		CacheExpiration: time.Now().Add(time.Minute * time.Duration(config.Get().MemCache.Expiration)).Unix(),
		data:            data,
	}
	memCacheSize += size
	//tlog.Debugf("Memcache write: %s, expiration: %d (%d total size)", path, memCache[path].expiration, memCacheSize)
	return nil
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
		if entry.CacheExpiration > time.Now().Unix() {
			continue
		}
		tlog.Debugf("Memcache expired: %s", path)
		memCacheSize -= size
		delete(memCache, path)
	}
	mu.Unlock()
	tlog.Debugf("Memcache truncate schedule complete in %s", time.Since(start))
}
