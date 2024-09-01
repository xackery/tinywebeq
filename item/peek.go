package item

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/jbsmith7741/toml"
	"gopkg.in/yaml.v3"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/library"
	"github.com/xackery/tinywebeq/model"
	"github.com/xackery/tinywebeq/site"
	"github.com/xackery/tinywebeq/store"
	"github.com/xackery/tinywebeq/tlog"
)

var (
	peekTemplate *template.Template
)

func peekInit() error {
	var err error
	peekTemplate = template.New("peek")
	peekTemplate, err = peekTemplate.ParseFS(site.TemplateFS(),
		"item/peek.go.tpl", // data
	)
	if err != nil {
		return fmt.Errorf("template.ParseFS: %w", err)
	}

	return nil
}

// Peek handles item peek requests
func Peek(w http.ResponseWriter, r *http.Request) {
	var err error
	var id int64

	if !config.Get().Item.IsEnabled {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	tlog.Debugf("peek: %s", r.URL.String())

	strID := r.URL.Query().Get("id")
	if len(strID) > 0 {
		id, err = strconv.ParseInt(strID, 10, 64)
		if err != nil {
			tlog.Errorf("strconv.Atoi: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tlog.Debugf("peekRender: id: %d", id)

	err = peekRender(ctx, id, w)
	if err != nil {
		if err.Error() == "item not found" {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		tlog.Errorf("peekRender: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func peekRender(ctx context.Context, id int64, w http.ResponseWriter) error {

	item, err := store.ItemByItemID(ctx, id)
	if err != nil {
		return fmt.Errorf("store.ItemByItemID: %w", err)
	}

	itemQuest, err := store.ItemQuestByItemID(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("store.ItemQuestByItemID: %w", err)
	}

	itemRecipe, err := store.ItemRecipeByItemID(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("store.ItemRecipeByItemID: %w", err)
	}

	type TemplateData struct {
		Site                site.BaseData
		Library             *library.Library
		ItemInfo            []string
		IsItemSearchEnabled bool
		Store               *store.Store
		Item                *model.Item
		ItemRecipe          *model.ItemRecipe
		ItemQuest           *model.ItemQuest
	}

	data := TemplateData{
		Site:                site.BaseDataInit("ITEM"),
		Library:             library.Instance(),
		IsItemSearchEnabled: config.Get().Item.Search.IsEnabled,
		Store:               store.Instance(),
		Item:                item,
		ItemRecipe:          itemRecipe,
		ItemQuest:           itemQuest,
	}

	//buf := &bytes.Buffer{}
	err = peekTemplate.ExecuteTemplate(w, "peek.go.tpl", data)
	if err != nil {
		return fmt.Errorf("peekTemplate.Execute: %w", err)
	}
	/*
		var tomlMap map[string][]string

		err = yaml.NewDecoder(buf).Decode(&tomlMap)
		if err != nil {
			return fmt.Errorf("yaml.NewDecoder: %w", err)
		}

		for k, lines := range tomlMap {
			newLines := []string{}
			for _, line := range lines {
				if line == "" {
					continue
				}
				newLines = append(newLines, line)
			}
			tomlMap[k] = newLines
		}

		err = toml.NewEncoder(w).Encode(tomlMap)
		if err != nil {
			return fmt.Errorf("toml.NewEncoder: %w", err)
		}
	*/
	return nil
}
