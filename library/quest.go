package library

import (
	"context"
	"fmt"
	"image"
	"image/draw"
	"os"
	"strconv"
	"strings"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/search/query"
	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/tlog"
)

var (
	questIndex bleve.Index
	questIcons = make(map[int]image.Image)
)

type QuestIndexData struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}

func initQuests() error {
	var err error

	err = os.MkdirAll("assets", os.ModePerm)
	if err != nil {
		return fmt.Errorf("os.MkdirAll: %w", err)
	}

	err = initQuestIcons()
	if err != nil {
		if os.IsNotExist(err) {
			tlog.Warnf("initQuestIcons: %v", err)
		}
		tlog.Infof("To add quest icons, copy uifiles/default/quests0*.tga to the assets folder")
	}

	if !config.Get().Quest.IsEnabled {
		return nil
	}
	return nil
}

func QuestIcon(id int) image.Image {
	return nil
}

func QuestSearchByName(ctx context.Context, name string) ([]QuestIndexData, error) {

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

	searchResults, err := questIndex.SearchInContext(ctx, searchRequest)
	if err != nil {
		return nil, fmt.Errorf("questIndex.Search: %w", err)
	}
	results := []QuestIndexData{}

	for _, hit := range searchResults.Hits {
		id, err := strconv.Atoi(hit.ID)
		if err != nil {
			return nil, fmt.Errorf("strconv.Atoi: %w", err)
		}
		name, ok := hit.Fields["name"].(string)
		if !ok {
			name = fmt.Sprintf("Unknown Quest (%d)", id)
		}
		level := 255
		levelField, ok := hit.Fields["level"].(float64)
		if !ok {
			tlog.Warnf("quest %d has no level", id)
			level = 255
		} else {
			level = int(levelField)
		}
		results = append(results,
			QuestIndexData{
				ID:    id,
				Name:  name,
				Level: level,
			})
	}
	tlog.Debugf("Search found %d results, %d hits", len(results), len(searchResults.Hits))
	return results, nil
}

func initQuestIcons() error {
	files := []string{
		"assets/quests01.tga",
		"assets/quests02.tga",
		"assets/quests03.tga",
		"assets/quests04.tga",
		"assets/quests05.tga",
		"assets/quests06.tga",
		"assets/quests07.tga",
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

				questIcons[index] = iconImg
				index++
			}
		}
		isLoaded = true
	}

	tlog.Debugf("Loaded %d quest icons", len(questIcons))

	return nil
}
