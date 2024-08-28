package recipe

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/tlog"
)

var (
	mu    sync.RWMutex
	start time.Time
)

func Init() error {

	if !config.Get().Recipe.IsEnabled {
		tlog.Debugf("Recipe scanning are disabled")
		return nil
	}

	if !config.Get().Recipe.IsBackgroundScanningEnabled {
		tlog.Debugf("Recipe background scanning is disabled")
	}

	go maintain()
	return nil
}

func maintain() {
	time.Sleep(10 * time.Second)

	tickerRecipeCache := time.NewTicker(time.Duration(config.Get().Recipe.ScanSchedule) * time.Second)

	for {
		select {
		case <-tickerRecipeCache.C:
			err := Parse(context.Background())
			if err != nil {
				tlog.Errorf("recipe walk: %s", err)
			}
		}
	}
}

func Parse(ctx context.Context) error {
	if !config.Get().Recipe.IsEnabled {
		tlog.Debugf("Recipes are disabled")
		return nil
	}
	mu.Lock()
	defer mu.Unlock()

	tlog.Debugf("Setting log level to info")

	tlog.SetLevel(zerolog.InfoLevel)

	tlog.Infof("Parsing recipes (this will take a while)")

	start = time.Now()
	chunkStart := time.Now()
	err := db.BBolt.ItemRecipeTruncate(ctx)
	if err != nil {
		return fmt.Errorf("truncate bbolt cache: %w", err)
	}

	if config.Get().IsDebugEnabled {
		tlog.SetLevel(zerolog.DebugLevel)
	}

	itemRecipes, err := db.Mysql.ItemRecipeAll(ctx)
	if err != nil {
		return fmt.Errorf("itemRecipeAll: %w", err)
	}

	totalCount := 0

	for _, itemRecipe := range itemRecipes {
		totalCount++
		if totalCount%10000 == 0 {
			tlog.Infof("Parsed %d recipes in %s (%s total)", totalCount, time.Since(chunkStart).String(), time.Since(start).String())
			chunkStart = time.Now()
		}

		err = db.BBolt.ItemRecipeReplace(ctx, itemRecipe.ItemID, itemRecipe)
		if err != nil {
			return fmt.Errorf("bbolt replace: %w", err)
		}
	}

	tlog.Infof("Parsed %d recipes in %s", totalCount, time.Since(start).String())
	return nil
}
