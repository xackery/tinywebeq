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
	questSearchMux sync.RWMutex
	questSearch    = map[string]*model.QuestSearch{}
)

func initQuestSearch(ctx context.Context) error {
	if !config.Get().Quest.Search.IsEnabled {
		return nil
	}

	if !config.Get().Quest.Search.IsMemorySearchEnabled {
		return nil
	}

	questSearchMux.Lock()
	defer questSearchMux.Unlock()

	questSearch = make(map[string]*model.QuestSearch)

	rows, err := db.BBolt.QuestsAll(ctx)
	if err != nil {
		return fmt.Errorf("questsAll: %w", err)
	}

	for _, row := range rows {
		questSearch[row.Name] = &model.QuestSearch{
			ID:    int64(row.ID),
			Name:  row.Name,
			Level: int64(row.Level),
		}
	}
	return nil
}

func QuestSearchByName(ctx context.Context, name string) ([]*model.QuestSearch, error) {
	if !config.Get().Quest.Search.IsEnabled {
		return nil, fmt.Errorf("quest search is disabled")
	}

	if !config.Get().Quest.Search.IsMemorySearchEnabled {
		results := []*model.QuestSearch{}

		rows, err := db.BBolt.QuestSearchByName(ctx, name)
		if err != nil {
			return nil, fmt.Errorf("quest search by name: %w", err)
		}
		for _, row := range rows {
			quest := &model.QuestSearch{
				ID:    int64(row.ID),
				Name:  row.Name,
				Level: int64(row.Level),
			}
			results = append(results, quest)
		}
		return results, nil
	}

	questSearchMux.RLock()
	defer questSearchMux.RUnlock()

	var quests []*model.QuestSearch

	quest, ok := questSearch[name]
	if ok {
		quests = append(quests, quest)
		return quests, nil
	}

	names := strings.Split(name, " ")
	for _, quest := range questSearch {
		for _, n := range names {
			if strings.Contains(strings.ToLower(quest.Name), strings.ToLower(n)) {
				quests = append(quests, quest)
				break
			}
		}
	}

	return quests, nil
}
