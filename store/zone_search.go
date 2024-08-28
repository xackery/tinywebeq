package store

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/model"
)

var (
	zoneSearchMux sync.RWMutex
	zoneSearch    = map[string]*model.ZoneSearch{}
)

func initZoneSearch(ctx context.Context) error {
	if !config.Get().Zone.Search.IsEnabled {
		return nil
	}

	if !config.Get().Zone.Search.IsMemorySearchEnabled {
		return nil
	}

	zoneSearchMux.Lock()
	defer zoneSearchMux.Unlock()

	zoneSearch = make(map[string]*model.ZoneSearch)

	rows, err := db.Mysql.ZonesAll(ctx)
	if err != nil {
		return fmt.Errorf("zonesAll: %w", err)
	}

	for _, row := range rows {
		zoneSearch[row.ShortName] = &model.ZoneSearch{
			ID:        int64(row.ID),
			ShortName: row.ShortName,
			LongName:  row.LongName,
			Level:     int64(row.Expansion),
		}
	}
	return nil
}

func ZoneSearchByName(ctx context.Context, name string) ([]*model.ZoneSearch, error) {
	if !config.Get().Zone.Search.IsEnabled {
		return nil, fmt.Errorf("zone search is disabled")
	}

	if !config.Get().Zone.Search.IsMemorySearchEnabled {
		results := []*model.ZoneSearch{}

		rows, err := db.Mysql.ZoneSearchByName(ctx, name)
		if err != nil {
			return nil, fmt.Errorf("zone search by name: %w", err)
		}
		for _, row := range rows {
			zone := &model.ZoneSearch{
				ID:        int64(row.ID),
				ShortName: row.ShortName,
				LongName:  row.LongName,
				Level:     int64(row.Expansion),
			}
			results = append(results, zone)
		}
		return results, nil
	}

	zoneSearchMux.RLock()
	defer zoneSearchMux.RUnlock()

	var zones []*model.ZoneSearch

	zone, ok := zoneSearch[name]
	if ok {
		zones = append(zones, zone)
		return zones, nil
	}

	names := strings.Split(name, " ")
	for _, zone := range zoneSearch {
		for _, n := range names {
			if strings.Contains(strings.ToLower(zone.ShortName), strings.ToLower(n)) {
				zones = append(zones, zone)
				break
			}
			if strings.Contains(strings.ToLower(zone.LongName), strings.ToLower(n)) {
				zones = append(zones, zone)
				break
			}
		}
	}

	return zones, nil
}
