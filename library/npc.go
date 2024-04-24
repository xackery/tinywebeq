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
	"github.com/blevesearch/bleve/mapping"
	"github.com/blevesearch/bleve/search/query"
	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/tlog"
)

var (
	npcIndex bleve.Index
	npcIcons = make(map[int]image.Image)
)

type NpcIndexData struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}

func initNpcs() error {
	var err error
	start := time.Now()

	err = os.MkdirAll("assets", os.ModePerm)
	if err != nil {
		return fmt.Errorf("os.MkdirAll: %w", err)
	}

	err = initNpcIcons()
	if err != nil {
		if os.IsNotExist(err) {
			tlog.Warnf("initNpcIcons: %v", err)
		}
		tlog.Infof("To add npc icons, copy uifiles/default/npcs0*.tga to the assets folder")
	}

	if !config.Get().Npc.IsEnabled {
		return nil
	}
	isSearchEnabled := config.Get().Npc.Search.IsEnabled

	totalCount := 0
	isNewIndex := false

	if isSearchEnabled {
		npcIndex, err = bleve.Open("cache/npc.bleve")
		if err != nil {
			if err != bleve.ErrorIndexPathDoesNotExist {
				return fmt.Errorf("bleve.Open: %w", err)
			}
			tlog.Infof("No cache/npc.bleve found, creating new index")

			englishTextFieldMapping := bleve.NewTextFieldMapping()
			englishTextFieldMapping.Analyzer = en.AnalyzerName

			npcMapping := bleve.NewDocumentMapping()
			npcMapping.AddFieldMappingsAt("name", englishTextFieldMapping)
			npcMapping.AddFieldMappingsAt("level", mapping.NewNumericFieldMapping())

			index := bleve.NewIndexMapping()
			index.AddDocumentMapping("npc", npcMapping)
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

			npcIndex, err = bleve.New("cache/npc.bleve", index)
			if err != nil {
				return fmt.Errorf("bleve.New: %w", err)
			}
			isNewIndex = true
		}
	}

	query := "SELECT id, name, level FROM npc_types"

	//fmt.Println(query)
	rows, err := db.Instance.Query(query)
	if err != nil {
		return fmt.Errorf("query npcs: %w", err)
	}
	defer rows.Close()

	var batch *bleve.Batch
	if isSearchEnabled {
		batch = npcIndex.NewBatch()
	}

	for rows.Next() {
		totalCount++
		ne := &NpcIndexData{}

		err = rows.Scan(&ne.ID, &ne.Name, &ne.Level)
		if err != nil {
			return fmt.Errorf("rows.Scan: %w", err)
		}

		out := ne.Name
		out = strings.ReplaceAll(out, "_", " ")
		out = strings.ReplaceAll(out, "-", "`")
		out = strings.ReplaceAll(out, "#", "")
		out = strings.ReplaceAll(out, "!", "")
		out = strings.ReplaceAll(out, "~", "")
		ne.Name = out

		if isSearchEnabled && isNewIndex {
			if totalCount%1000 == 0 {
				if totalCount%10000 == 0 {
					tlog.Infof("Indexed %d out of ~67000 npcs", totalCount)
				}
				err = npcIndex.Batch(batch)
				if err != nil {
					return fmt.Errorf("npcIndex.Batch: %w", err)
				}
				batch = npcIndex.NewBatch()
			}

			err = batch.Index(fmt.Sprintf("%d", ne.ID), ne)
			if err != nil {
				return fmt.Errorf("npcIndex.Index: %w", err)
			}
		}
	}

	if isSearchEnabled && isNewIndex {
		err = npcIndex.Batch(batch)
		if err != nil {
			return fmt.Errorf("npcIndex.Batch: %w", err)
		}
		tlog.Debugf("Loaded %d npcs in %s", totalCount, time.Since(start).String())
		return nil
	}
	return nil
}

func NpcIcon(id int) image.Image {
	return nil
}

func NpcSearchByName(ctx context.Context, name string) ([]NpcIndexData, error) {

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

	searchResults, err := npcIndex.SearchInContext(ctx, searchRequest)
	if err != nil {
		return nil, fmt.Errorf("npcIndex.Search: %w", err)
	}
	results := []NpcIndexData{}

	for _, hit := range searchResults.Hits {
		id, err := strconv.Atoi(hit.ID)
		if err != nil {
			return nil, fmt.Errorf("strconv.Atoi: %w", err)
		}
		name, ok := hit.Fields["name"].(string)
		if !ok {
			name = fmt.Sprintf("Unknown Npc (%d)", id)
		}
		level := 255
		levelField, ok := hit.Fields["level"].(float64)
		if !ok {
			tlog.Warnf("npc %d has no level", id)
			level = 255
		} else {
			level = int(levelField)
		}
		results = append(results,
			NpcIndexData{
				ID:    id,
				Name:  name,
				Level: level,
			})
	}
	tlog.Debugf("Search found %d results, %d hits", len(results), len(searchResults.Hits))
	return results, nil
}

func initNpcIcons() error {
	files := []string{
		"assets/npcs01.tga",
		"assets/npcs02.tga",
		"assets/npcs03.tga",
		"assets/npcs04.tga",
		"assets/npcs05.tga",
		"assets/npcs06.tga",
		"assets/npcs07.tga",
	}

	index := 0
	isLoaded := false
	for _, file := range files {
		img, err := fetchTGA(file)
		if err != nil {
			if isLoaded {
				return nil
			}
			return fmt.Errorf("fetchTGA: %w", err)
		}

		isEmpty := false
		for i := 0; i < 6; i++ {
			if isEmpty {
				break
			}
			for j := 0; j < 6; j++ {
				//subImg := img.(*image.NRGBA).SubImage(image.Rect(j*40, i*41, j*40+40, i*41+41))
				// move subimg pixels to 0,0
				iconImg := image.NewNRGBA(image.Rect(0, 0, 40, 40))
				//draw.Draw(iconImg, iconImg.Bounds(), subImg, image.Pt(0, 0), draw.Src)
				draw.Draw(iconImg, iconImg.Bounds(), img, image.Pt(j*40, i*40), draw.Src)

				if iconImg.At(0, 0) == image.Transparent {
					isEmpty = true
					break
				}

				npcIcons[index] = iconImg
				index++
			}
		}
		isLoaded = true
	}

	tlog.Debugf("Loaded %d npc icons", len(npcIcons))

	return nil
}
