package cache

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/model"
	"github.com/xackery/tinywebeq/tlog"
	"gopkg.in/yaml.v3"
)

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
