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
	"github.com/blevesearch/bleve/search/query"
	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/tlog"
)

var (
	itemIndex bleve.Index
)

type ItemIndexData struct {
	ID    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Level int    `json:"level"`
}

func initItems() error {
	var err error
	start := time.Now()

	isSearchEnabled := config.Get().Item.IsSearchEnabled

	if !isSearchEnabled {
		return nil
	}

	totalCount := 0
	isNewIndex := false

	itemIndex, err = bleve.Open("cache/item.bleve")
	if err != nil {
		if err != bleve.ErrorIndexPathDoesNotExist {
			return fmt.Errorf("bleve.Open: %w", err)
		}
		tlog.Infof("No cache/item.bleve found, creating new index")

		englishTextFieldMapping := bleve.NewTextFieldMapping()
		englishTextFieldMapping.Analyzer = en.AnalyzerName

		itemMapping := bleve.NewDocumentMapping()
		itemMapping.AddFieldMappingsAt("name", englishTextFieldMapping)

		index := bleve.NewIndexMapping()
		index.AddDocumentMapping("item", itemMapping)
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

		itemIndex, err = bleve.New("cache/item.bleve", index)
		if err != nil {
			return fmt.Errorf("bleve.New: %w", err)
		}
		isNewIndex = true
	}

	query := "SELECT id, name, ac, reqlevel, reclevel, hp, damage, delay, mana FROM items"

	rows, err := db.Instance.Query(query)
	if err != nil {
		return fmt.Errorf("query items: %w", err)
	}
	defer rows.Close()

	batch := itemIndex.NewBatch()

	for rows.Next() {
		totalCount++
		ie := &ItemIndexData{}

		recLevel, reqLevel, hp, mana, damage, delay := 0, 0, 0, 0, 0, 0
		ac := 0
		err = rows.Scan(&ie.ID, &ie.Name, &ac, &reqLevel, &recLevel, &hp, &damage, &delay, &mana)
		if err != nil {
			return fmt.Errorf("rows.Scan: %w", err)
		}
		if isNewIndex {
			if totalCount%1000 == 0 {
				if totalCount%10000 == 0 {
					tlog.Infof("Indexed %d out of ~117000 items", totalCount)
				}
				err = itemIndex.Batch(batch)
				if err != nil {
					return fmt.Errorf("itemIndex.Batch: %w", err)
				}
				batch = itemIndex.NewBatch()
			}

			level := 0
			if hp > 5 && level < 5 {
				level = 5
			}
			ratio := float64(damage) / float64(delay)
			if ratio > 1.5 && level < 10 {
				level = 10
			}
			if mana > 5 && level < 15 {
				level = 15
			}

			if reqLevel > 0 && level < reqLevel {
				level = reqLevel
			}
			if recLevel > 0 && level < recLevel {
				level = recLevel
			}
			if ac > 0 {
				if ac < 5 && level < 5 {
					level = 5
				}
				if ac < 10 && level < 10 {
					level = 10
				}
				if ac < 15 && level < 15 {
					level = 15
				}
				if ac < 20 && level < 20 {
					level = 20
				}
			}

			if level > config.Get().MaxLevel {
				continue
			}

			itemData := ItemIndexData{
				Name:  ie.Name,
				Level: level,
			}

			err = batch.Index(fmt.Sprintf("%d", ie.ID), itemData)
			if err != nil {
				return fmt.Errorf("itemIndex.Index: %w", err)
			}
		}
	}

	if isNewIndex {
		err = itemIndex.Batch(batch)
		if err != nil {
			return fmt.Errorf("itemIndex.Batch: %w", err)
		}
		tlog.Debugf("Loaded %d items in %s", totalCount, time.Since(start).String())
		return nil
	}
	return nil
}

func ItemName(id int) string {
	return fmt.Sprintf("Unknown Item (%d)", id)
}

func ItemSearchByName(ctx context.Context, name string) ([]ItemIndexData, error) {
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
	searchRequest.SortBy([]string{"-level"})

	searchResults, err := itemIndex.SearchInContext(ctx, searchRequest)
	if err != nil {
		return nil, fmt.Errorf("itemIndex.Search: %w", err)
	}
	results := []ItemIndexData{}

	for _, hit := range searchResults.Hits {
		id, err := strconv.Atoi(hit.ID)
		if err != nil {
			return nil, fmt.Errorf("strconv.Atoi: %w", err)
		}
		name, ok := hit.Fields["name"].(string)
		if !ok {
			name = fmt.Sprintf("Unknown Item (%d)", id)
		}
		level := 0
		levelField, ok := hit.Fields["level"].(float64)
		if !ok {
			tlog.Warnf("spell %d has no level", id)
			level = 0
		} else {
			level = int(levelField)
		}
		results = append(results,
			ItemIndexData{
				ID:    id,
				Name:  name,
				Level: level,
			})
	}
	tlog.Debugf("Search found %d results, %d hits", len(results), len(searchResults.Hits))
	return results, nil
}
