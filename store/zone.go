package store

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/model"
	"github.com/xackery/tinywebeq/tlog"
)

func ZoneLongNameByShortName(shortName string) string {
	zone := ZoneByShortName(shortName)
	if zone == nil {
		return fmt.Sprintf("Unknown zone (%s)", shortName)
	}
	return zone.LongName
}

func ZoneLongNameByZoneIDNumber(id int32) string {
	zone, err := ZoneByZoneIDNumber(context.Background(), int64(id))
	if err != nil {
		tlog.Debugf("ZoneLongNameByID %d: %v", id, err)
		return fmt.Sprintf("Unknown zone (%d)", id)
	}

	return zone.LongName
}

func ZoneByShortName(shortName string) *model.Zone {
	zone, err := db.Mysql.ZoneByShortName(context.Background(), shortName)
	if err != nil {
		tlog.Debugf("ZoneByShortName %s: %v", shortName, err)
		return nil
	}

	return zone
}

func ZoneByZoneIDNumber(ctx context.Context, zoneID int64) (*model.Zone, error) {
	zone, err := db.Mysql.ZoneByZoneIDNumber(ctx, zoneID)
	if err != nil {
		return nil, fmt.Errorf("zoneByZoneIDNumber: %w", err)
	}
	return zone, nil
}
