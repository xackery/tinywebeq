package library

import (
	"fmt"

	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/tlog"
)

var (
	spells = map[int]*Spell{}
)

type Spell struct {
	ID           int    `db:"id"`
	Name         string `db:"name"`
	Attribs      []int  // effectid
	Bases        []int  // effect_base_value
	Calcs        []int  // formula
	Limits       []int  // effect_limit_value
	Maxes        []int  // max
	Classes      []int  // classes
	Range        int
	DurationCap  int
	DurationCalc int
	MaxTargets   int
	TargetType   int
	Skill        int
	RecoveryTime int
	RecastTime   int
	Pushback     int
	TeleportZone string
	Mana         int
}

func initSpells() error {
	spells = map[int]*Spell{}

	query := "SELECT id, name, TargetType, maxtargets, buffduration, skill, "
	for i := 1; i < 13; i++ {
		query += fmt.Sprintf("effectid%d, ", i)
	}
	for i := 1; i < 13; i++ {
		query += fmt.Sprintf("effect_base_value%d, ", i)
	}
	for i := 1; i < 13; i++ {
		query += fmt.Sprintf("effect_limit_value%d, ", i)
	}
	for i := 1; i < 13; i++ {
		query += fmt.Sprintf("max%d, ", i)
	}
	for i := 1; i < 13; i++ {
		query += fmt.Sprintf("formula%d, ", i)
	}
	for i := 1; i < 17; i++ {
		query += fmt.Sprintf("classes%d, ", i)
	}
	query += "`range`, recovery_time, recast_time, buffduration, pushback, teleport_zone, mana"

	query += " FROM spells_new"

	//fmt.Println(query)
	rows, err := db.Instance.Query(query)
	if err != nil {
		return fmt.Errorf("query spells: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		se := &Spell{
			Attribs: make([]int, 12),
			Bases:   make([]int, 12),
			Calcs:   make([]int, 12),
			Limits:  make([]int, 12),
			Maxes:   make([]int, 12),
			Classes: make([]int, 16),
		}

		err = rows.Scan(&se.ID, &se.Name, &se.TargetType, &se.MaxTargets, &se.DurationCap, &se.Skill,
			&se.Attribs[0], &se.Attribs[1], &se.Attribs[2], &se.Attribs[3], &se.Attribs[4], &se.Attribs[5], &se.Attribs[6], &se.Attribs[7], &se.Attribs[8], &se.Attribs[9], &se.Attribs[10], &se.Attribs[11],
			&se.Bases[0], &se.Bases[1], &se.Bases[2], &se.Bases[3], &se.Bases[4], &se.Bases[5], &se.Bases[6], &se.Bases[7], &se.Bases[8], &se.Bases[9], &se.Bases[10], &se.Bases[11],
			&se.Limits[0], &se.Limits[1], &se.Limits[2], &se.Limits[3], &se.Limits[4], &se.Limits[5], &se.Limits[6], &se.Limits[7], &se.Limits[8], &se.Limits[9], &se.Limits[10], &se.Limits[11],
			&se.Maxes[0], &se.Maxes[1], &se.Maxes[2], &se.Maxes[3], &se.Maxes[4], &se.Maxes[5], &se.Maxes[6], &se.Maxes[7], &se.Maxes[8], &se.Maxes[9], &se.Maxes[10], &se.Maxes[11],
			&se.Calcs[0], &se.Calcs[1], &se.Calcs[2], &se.Calcs[3], &se.Calcs[4], &se.Calcs[5], &se.Calcs[6], &se.Calcs[7], &se.Calcs[8], &se.Calcs[9], &se.Calcs[10], &se.Calcs[11],
			&se.Classes[0], &se.Classes[1], &se.Classes[2], &se.Classes[3], &se.Classes[4], &se.Classes[5], &se.Classes[6], &se.Classes[7], &se.Classes[8], &se.Classes[9], &se.Classes[10], &se.Classes[11], &se.Classes[12], &se.Classes[13], &se.Classes[14], &se.Classes[15],
			&se.Range, &se.RecoveryTime, &se.RecastTime, &se.DurationCalc, &se.Pushback, &se.TeleportZone, &se.Mana,
		)
		if err != nil {
			return fmt.Errorf("rows.Scan: %w", err)
		}
		spells[se.ID] = se
	}
	tlog.Debugf("Loaded %d spells", len(spells))
	return nil
}

func SpellName(id int) string {
	mu.Lock()
	defer mu.Unlock()
	se, ok := spells[id]
	if !ok {
		return fmt.Sprintf("Unknown Spell (%d)", id)
	}
	return se.Name
}

func SpellByID(id int) *Spell {
	mu.RLock()
	defer mu.RUnlock()
	se, ok := spells[id]
	if !ok {
		return nil
	}
	return se
}
