package library

import (
	"context"
	"fmt"
	"image"
	"image/draw"
	"os"
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
	"github.com/xypwn/filediver/dds"
)

var (
	itemIndex bleve.Index
	itemIcons = make(map[int]image.Image)
)

type ItemIndexData struct {
	ID    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Level int    `json:"level"`
}

func initItems() error {
	var err error
	start := time.Now()

	err = os.MkdirAll("assets", os.ModePerm)
	if err != nil {
		return fmt.Errorf("os.MkdirAll: %w", err)
	}

	err = initItemIcons()
	if err != nil {
		if os.IsNotExist(err) {
			tlog.Warnf("initItemIcons: %v", err)
		}
		tlog.Warnf("%v", err)
		tlog.Infof("To add item icons, copy uifiles/default/dragitem*.dds to the assets folder")
	}

	isSearchEnabled := config.Get().Item.Search.IsEnabled

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

func initItemIcons() error {
	files := []string{}
	for i := 1; i < 179; i++ {
		files = append(files, fmt.Sprintf("assets/dragitem%d.dds", i))
	}

	index := 500
	for _, file := range files {
		img, err := fetchDDS(file)
		if err != nil {
			return fmt.Errorf("fetchDDS: %w", err)
		}

		for x := 0; x+40 <= img.Bounds().Dx(); x += 40 {
			for y := 0; y+40 <= img.Bounds().Dy(); y += 40 {
				//subImg := img.(*image.NRGBA).SubImage(image.Rect(j*40, i*41, j*40+40, i*41+41))
				// move subimg pixels to 0,0
				iconImg := image.NewRGBA(image.Rect(0, 0, 40, 40))
				//draw.Draw(iconImg, iconImg.Bounds(), subImg, image.Pt(0, 0), draw.Src)
				draw.Draw(iconImg, iconImg.Bounds(), img, image.Pt(x, y), draw.Src)

				itemIcons[index] = iconImg

				index++
			}
		}
	}

	tlog.Debugf("Loaded %d item icons", len(itemIcons))

	return nil
}

func fetchDDS(path string) (image.Image, error) {
	r, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("os.Open: %w", err)
	}
	defer r.Close()

	img, err := dds.Decode(r, false)
	if err != nil {
		return nil, fmt.Errorf("dds.Decode: %w", err)
	}

	return img, nil
}

func ItemIcon(id int) image.Image {
	mu.RLock()
	defer mu.RUnlock()
	img, ok := itemIcons[id]
	if !ok {
		return nil
	}
	return img
}
