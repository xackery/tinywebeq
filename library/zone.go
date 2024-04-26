package library

import (
	"fmt"

	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/tlog"
)

var (
	zones = map[int]*Zone{}
)

type Zone struct {
	ZoneIDNumber int    `db:"zoneidnumber"`
	ShortName    string `db:"short_name"`
	LongName     string `db:"long_name"`
	Expansion    int    `db:"expansion"`
	MinExpansion int    `db:"min_expansion"`
	MaxExpansion int    `db:"max_expansion"`
}

func initZones() error {
	zones = map[int]*Zone{}

	query := "SELECT zoneidnumber, short_name, long_name, expansion, min_expansion, max_expansion FROM zone"
	rows, err := db.Instance.Query(query)
	if err != nil {
		return fmt.Errorf("query spells: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		ze := &Zone{}
		err = rows.Scan(&ze.ZoneIDNumber, &ze.ShortName, &ze.LongName, &ze.Expansion, &ze.MinExpansion, &ze.MaxExpansion)
		if err != nil {
			return fmt.Errorf("rows.Scan: %w", err)
		}
		zones[ze.ZoneIDNumber] = ze
	}
	tlog.Debugf("Loaded %d zones", len(zones))
	return nil
}

func ZoneByID(id int) *Zone {
	mu.RLock()
	defer mu.RUnlock()

	if zone, ok := zones[id]; ok {
		return zone
	}
	return nil
}

func ZoneLongNameByID(id int) string {
	mu.RLock()
	defer mu.RUnlock()

	if zone, ok := zones[id]; ok {
		return zone.LongName
	}
	return fmt.Sprintf("Unknown Zone %d", id)
}

func ZoneLongNameByShortName(shortName string) string {
	mu.RLock()
	defer mu.RUnlock()
	for _, zone := range zones {
		if zone.ShortName == shortName {
			return zone.LongName
		}
	}
	return fmt.Sprintf("Unknown Zone %s", shortName)
}

func ZoneIDByShortName(shortName string) int {
	mu.RLock()
	defer mu.RUnlock()

	for _, zone := range zones {
		if zone.ShortName == shortName {
			return zone.ZoneIDNumber
		}
	}
	return 0
}
