package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/db/mysql/storage/mysqlc"
	"github.com/xackery/tinywebeq/models"
)

func (b *Mysql) ZonesAll(ctx context.Context) ([]*models.Zone, error) {
	zones := []*models.Zone{}
	rows, err := b.query.ZonesAll(ctx, int8(config.Get().CurrentExpansion))
	if err != nil {
		return nil, fmt.Errorf("zones all: %w", err)
	}
	for _, row := range rows {
		zone := &models.Zone{
			ID:           int32(row.ID),
			Zoneidnumber: int32(row.Zoneidnumber),
			ShortName:    row.ShortName.String,
			LongName:     row.LongName,
			Expansion:    row.Expansion,
		}
		zones = append(zones, zone)
	}
	return zones, nil
}

func (b *Mysql) ZoneByShortName(ctx context.Context, shortName string) (*models.Zone, error) {
	zone := &models.Zone{}
	row, err := b.query.ZoneByShortName(ctx, mysqlc.ZoneByShortNameParams{
		ShortName: sql.NullString{String: shortName, Valid: true},
	})
	if err != nil {
		return nil, fmt.Errorf("zone by short name: %w", err)
	}
	zone.ID = int32(row.ID)
	zone.ShortName = row.ShortName.String
	zone.LongName = row.LongName
	zone.Expansion = row.Expansion
	return zone, nil
}

func (b *Mysql) ZoneSearchByName(ctx context.Context, name string) ([]*models.ZoneSearch, error) {
	zones := []*models.ZoneSearch{}

	rows, err := b.query.ZoneSearchByName(ctx, mysqlc.ZoneSearchByNameParams{
		ShortName: sql.NullString{String: name, Valid: true},
		LongName:  name,
	})
	if err != nil {
		return nil, fmt.Errorf("zone search by name: %w", err)
	}
	for _, row := range rows {
		zone := &models.ZoneSearch{
			ID:        int64(row.ID),
			ShortName: row.ShortName.String,
			LongName:  row.LongName,
			Level:     int64(row.Expansion),
		}
		zones = append(zones, zone)
	}
	return zones, nil
}

func (b *Mysql) ZoneByZoneIDNumber(ctx context.Context, zoneID int64) (*models.Zone, error) {
	zone := &models.Zone{}
	row, err := b.query.ZoneByZoneIDNumber(ctx, mysqlc.ZoneByZoneIDNumberParams{
		Zoneidnumber: int32(zoneID),
		Expansion:    int8(config.Get().CurrentExpansion),
	})
	if err != nil {
		return nil, fmt.Errorf("query zone: %w", err)
	}
	zone.ID = int32(row.ID)
	zone.ShortName = row.ShortName.String
	zone.LongName = row.LongName
	zone.Expansion = row.Expansion
	return zone, nil
}
