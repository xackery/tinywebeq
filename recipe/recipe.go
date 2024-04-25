package recipe

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/xackery/tinywebeq/cache"
	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/model"
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

	err := cache.TruncateSqliteCache(ctx, "item_recipe")
	if err != nil {
		return fmt.Errorf("truncate sqlite cache: %w", err)
	}

	if config.Get().IsDebugEnabled {
		tlog.SetLevel(zerolog.DebugLevel)
	}

	query := `SELECT tr.id recipe_id, tr.name recipe_name, 
tr.tradeskill, tr.trivial, tre.item_id, tre.iscontainer is_container,
tre.componentcount component_count, tre.successcount success_count
FROM tradeskill_recipe tr, tradeskill_recipe_entries tre
WHERE tr.id = tre.recipe_id
AND tr.enabled = 1
AND tre.componentcount > 0
ORDER by tre.item_id ASC`

	rows, err := db.Instance.QueryxContext(ctx, query)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	totalCount := 0

	itemRecipe := &model.ItemRecipe{}
	itemID := 0
	for rows.Next() {
		totalCount++
		if totalCount%10000 == 0 {
			tlog.Infof("Parsed %d recipes in %s (%s total)", totalCount, time.Since(chunkStart).String(), time.Since(start).String())
			chunkStart = time.Now()
		}
		ire := &model.ItemRecipeEntry{}
		err = rows.StructScan(ire)
		if err != nil {
			return fmt.Errorf("scan: %w", err)
		}
		if itemID == 0 {
			itemID = ire.ItemID
		}

		if itemID == ire.ItemID {
			itemRecipe.Entries = append(itemRecipe.Entries, ire)
			continue
		}
		path := fmt.Sprintf("item_recipe/%d.yaml", itemID)

		err := cache.WriteSqlite(ctx, path, itemRecipe)
		if err != nil {
			return fmt.Errorf("write: %w", err)
		}
		itemID = ire.ItemID
		itemRecipe = &model.ItemRecipe{}
		itemRecipe.Entries = append(itemRecipe.Entries, ire)
	}

	tlog.Infof("Parsed %d recipes in %s", totalCount, time.Since(start).String())
	return nil
}
