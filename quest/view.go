package quest

import (
	"context"
	"fmt"
	"io/fs"
	"net/http"
	"strconv"
	"time"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/library"
	"github.com/xackery/tinywebeq/models"
	"github.com/xackery/tinywebeq/site"
	"github.com/xackery/tinywebeq/store"
	"github.com/xackery/tinywebeq/template"
	"github.com/xackery/tinywebeq/tlog"
)

// View handles quest view requests
func View(templates fs.FS) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var id int64

		if !config.Get().Quest.IsEnabled {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		tlog.Debugf("view: %s", r.URL.String())

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

		tlog.Debugf("viewRender: id: %d", id)

		err = viewRender(ctx, templates, id, w)
		if err != nil {
			if err.Error() == "quest not found" {
				http.Error(w, "Not Found", http.StatusNotFound)
				return
			}
			tlog.Errorf("viewRender: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		tlog.Debugf("viewRender: id: %d done", id)
	}
}

func viewRender(ctx context.Context, templates fs.FS, id int64, w http.ResponseWriter) error {

	quest, err := store.QuestByQuestID(ctx, id)
	if err != nil {
		return fmt.Errorf("store.QuestByQuestID: %w", err)
	}

	data := struct {
		Site                 site.BaseData
		Quest                *models.Quest
		Library              *library.Library
		QuestInfo            []string
		IsQuestSearchEnabled bool
	}{
		Site:                 site.BaseDataInit(quest.Name),
		Quest:                quest,
		Library:              library.Instance(),
		IsQuestSearchEnabled: config.Get().Quest.Search.IsEnabled,
	}

	if config.Get().Quest.Preview.IsEnabled {
		data.Site.ImageURL = fmt.Sprintf("/quest/preview.png?id=%d", id)
	}

	view, err := template.Compile("quest", "quest/view.go.tmpl", templates)
	if err != nil {
		return fmt.Errorf("template.Compile: %w", err)
	}

	if err = view.ExecuteTemplate(w, "content.go.tmpl", data); err != nil {
		return fmt.Errorf("viewTemplate.Execute: %w", err)
	}

	return nil
}
