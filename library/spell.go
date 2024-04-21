package library

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/analysis/lang/en"
	"github.com/blevesearch/bleve/analysis/token/lowercase"
	"github.com/blevesearch/bleve/analysis/token/porter"
	"github.com/blevesearch/bleve/mapping"
	"github.com/blevesearch/bleve/search/query"
	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/tlog"
)

var (
	spells     = map[int]*Spell{}
	spellIndex bleve.Index
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

type SpellIndexData struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}

func initSpells() error {
	var err error
	start := time.Now()

	isSearchEnabled := config.Get().Spell.IsSearchEnabled

	if !isSearchEnabled {
		return nil
	}

	totalCount := 0
	isNewIndex := false

	spellIndex, err = bleve.Open("cache/spell.bleve")
	if err != nil {
		if err != bleve.ErrorIndexPathDoesNotExist {
			return fmt.Errorf("bleve.Open: %w", err)
		}
		tlog.Infof("No cache/spell.bleve found, creating new index")

		englishTextFieldMapping := bleve.NewTextFieldMapping()
		englishTextFieldMapping.Analyzer = en.AnalyzerName

		spellMapping := bleve.NewDocumentMapping()
		spellMapping.AddFieldMappingsAt("name", englishTextFieldMapping)
		spellMapping.AddFieldMappingsAt("level", mapping.NewNumericFieldMapping())

		index := bleve.NewIndexMapping()
		index.AddDocumentMapping("spell", spellMapping)
		index.TypeField = "type"
		index.DefaultAnalyzer = "en"
		index.AddCustomAnalyzer("en", map[string]interface{}{
			"type":      "standard",
			"tokenizer": "whitespace",
			"token_filters": []string{
				en.PossessiveName,
				lowercase.Name,
				en.StopName,
				porter.Name,
			},
		})

		spellIndex, err = bleve.New("cache/spell.bleve", index)
		if err != nil {
			return fmt.Errorf("bleve.New: %w", err)
		}
		isNewIndex = true
	}

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

	batch := spellIndex.NewBatch()

	for rows.Next() {
		totalCount++
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
		if isNewIndex {
			if totalCount%1000 == 0 {
				if totalCount%10000 == 0 {
					tlog.Infof("Indexed %d out of ~40000 spells", totalCount)
				}
				err = spellIndex.Batch(batch)
				if err != nil {
					return fmt.Errorf("spellIndex.Batch: %w", err)
				}
				batch = spellIndex.NewBatch()
			}

			spellData := SpellIndexData{
				Name:  se.Name,
				Level: 255,
			}
			for i := 0; i < 16; i++ {
				if se.Classes[i] > 0 && se.Classes[i] < 255 {
					newLevel := se.Classes[i]
					if newLevel >= spellData.Level {
						continue
					}
					spellData.Level = newLevel
				}
			}

			if config.Get().Spell.IsSearchOnlyPlayerSpells {
				if spellData.Level == 255 {
					continue
				}
				if spellData.Level > config.Get().MaxLevel {
					continue
				}
			}

			err = batch.Index(fmt.Sprintf("%d", se.ID), spellData)
			if err != nil {
				return fmt.Errorf("spellIndex.Index: %w", err)
			}
		}
	}

	if isNewIndex {
		err = spellIndex.Batch(batch)
		if err != nil {
			return fmt.Errorf("spellIndex.Batch: %w", err)
		}
		tlog.Debugf("Loaded %d spells in %s", totalCount, time.Since(start).String())
		return nil
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

func SpellSearchByName(ctx context.Context, name string) ([]SpellIndexData, error) {

	searches := []query.Query{}

	terms := strings.Split(name, " ")
	for _, term := range terms {
		search := bleve.NewFuzzyQuery(term)
		search.SetField("name")
		search.SetFuzziness(1)
		searches = append(searches, search)
	}
	multiQuery := bleve.NewConjunctionQuery(searches...)

	searchRequest := bleve.NewSearchRequestOptions(multiQuery, 10, 0, true)
	searchRequest.Fields = []string{"name", "level"}
	searchRequest.SortBy([]string{"level"})

	searchResults, err := spellIndex.SearchInContext(ctx, searchRequest)
	if err != nil {
		return nil, fmt.Errorf("spellIndex.Search: %w", err)
	}
	results := []SpellIndexData{}

	for _, hit := range searchResults.Hits {
		id, err := strconv.Atoi(hit.ID)
		if err != nil {
			return nil, fmt.Errorf("strconv.Atoi: %w", err)
		}
		name, ok := hit.Fields["name"].(string)
		if !ok {
			name = fmt.Sprintf("Unknown Spell (%d)", id)
		}
		level := 255
		levelField, ok := hit.Fields["level"].(float64)
		if !ok {
			tlog.Warnf("spell %d has no level", id)
			level = 255
		} else {
			level = int(levelField)
		}
		results = append(results,
			SpellIndexData{
				ID:    id,
				Name:  name,
				Level: level,
			})
	}
	tlog.Debugf("Search found %d results, %d hits", len(results), len(searchResults.Hits))
	return results, nil
}
