package item

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/xackery/tinywebeq/cache"
	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/tlog"

	"github.com/xackery/tinywebeq/db"
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

	path := fmt.Sprintf("item/view_%d.html", id)

	ok, cacheData := cache.Read(path)
	if ok {
		_, err := w.Write(cacheData)
		if err != nil {
			return fmt.Errorf("from cache write: %w", err)
		}
		return nil
	}

	query := "SELECT * FROM items WHERE id=:id LIMIT 1"
	if config.Get().IsDiscoveredOnly {
		query += "SELECT * FROM items, discovered items WHERE items.id=:id AND discovered_items.item_id=:id LIMIT 1"
	}

	rows, err := db.Query(ctx,
		query,
		map[string]interface{}{
			"id": id,
		})
	if err != nil {
		return fmt.Errorf("query items: %w", err)
	}
	defer rows.Close()

	type TemplateData struct {
		Site site.BaseData
		Item *Table
	}

	data := TemplateData{Site: site.BaseDataInit("Item View")}
	item := &Table{}

	if !rows.Next() {
		http.Error(w, "Not Found", http.StatusNotFound)
		return nil
	}

	err = rows.StructScan(item)
	if err != nil {
		return fmt.Errorf("rows.StructScan: %w", err)
	}
	data.Item = item

	buf := &bytes.Buffer{}

	err = viewTemplate.ExecuteTemplate(buf, "content.go.tpl", data)
	if err != nil {
		return fmt.Errorf("viewTemplate.Execute: %w", err)
	}

	err = cache.Write(path, buf.Bytes())
	if err != nil {
		return fmt.Errorf("cache.Write: %w", err)
	}

	_, err = w.Write(buf.Bytes())
	if err != nil {
		return fmt.Errorf("w.Write: %w", err)
	}

	return nil
}
