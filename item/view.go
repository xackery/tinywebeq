package item

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/library"
	"github.com/xackery/tinywebeq/model"
	"github.com/xackery/tinywebeq/tlog"

	"github.com/xackery/tinywebeq/site"
)

var (
	viewTemplate *template.Template
)

// https://allaclone.wayfarershaven.com/?a=item&id=1004
func viewInit() error {
	var err error
	viewTemplate = template.New("view")
	viewTemplate, err = viewTemplate.ParseFS(site.TemplateFS(),
		"item/view.go.tpl",      // data
		"head.go.tpl",           // head
		"header.go.tpl",         // header
		"footer.go.tpl",         // footer
		"layout/content.go.tpl", // layout (requires footer, header, head, data)
	)
	if err != nil {
		return fmt.Errorf("template.ParseFS: %w", err)
	}
	return nil
}

// View handles item view requests
func View(w http.ResponseWriter, r *http.Request) {
	var err error
	var id int

	tlog.Debugf("view: %s", r.URL.String())

	strID := r.URL.Query().Get("id")
	if len(strID) > 0 {
		id, err = strconv.Atoi(strID)
		if err != nil {
			tlog.Errorf("strconv.Atoi: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tlog.Debugf("viewRender: id: %d", id)

	err = viewRender(ctx, id, w)
	if err != nil {
		tlog.Errorf("viewRender: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tlog.Debugf("viewRender: id: %d done", id)
}

func viewRender(ctx context.Context, id int, w http.ResponseWriter) error {
	if id <= 1000 {
		return fmt.Errorf("id too low")
	}

	item, err := fetchItem(ctx, id)
	if err != nil {
		return fmt.Errorf("fetchItem: %w", err)
	}

	itemQuest, err := fetchItemQuest(ctx, id)
	if err != nil {
		tlog.Debugf("Ignoring err fetchItemQuest: %v", err)
	}

	itemRecipe, err := fetchItemRecipe(ctx, id)
	if err != nil {
		tlog.Debugf("Ignoring err fetchItemRecipe: %v", err)
	}

	type TemplateData struct {
		Site                site.BaseData
		Item                *model.Item
		Library             *library.Library
		IsItemSearchEnabled bool
		ItemQuest           *model.ItemQuest
		ItemRecipe          *model.ItemRecipe
	}

	data := TemplateData{
		Site:                site.BaseDataInit(item.Name),
		Item:                item,
		Library:             library.Instance(),
		IsItemSearchEnabled: config.Get().Item.Search.IsEnabled,
		ItemQuest:           itemQuest,
		ItemRecipe:          itemRecipe,
	}

	err = viewTemplate.ExecuteTemplate(w, "content.go.tpl", data)
	if err != nil {
		return fmt.Errorf("viewTemplate.Execute: %w", err)
	}

	return nil
}
