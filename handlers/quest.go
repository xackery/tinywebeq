package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"go.uber.org/zap"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/image"
	"github.com/xackery/tinywebeq/library"
	"github.com/xackery/tinywebeq/models"
	"github.com/xackery/tinywebeq/site"
	"github.com/xackery/tinywebeq/store"
	"github.com/xackery/tinywebeq/tlog"
)

func (h *Handlers) IndexQuests() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var name string
		var err error
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if !config.Get().Quest.Search.IsEnabled {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		tlog.Debugf("search: %s", r.URL.String())

		type Response struct {
			Code    int                   `json:"code"`
			Message string                `json:"message"`
			Quests  []*models.QuestSearch `json:"quests,omitempty"`
		}

		name = r.URL.Query().Get("name")
		if len(name) < 1 {
			resp := Response{
				Code:    400,
				Message: "Invalid name",
			}
			err = json.NewEncoder(w).Encode(resp)
			if err != nil {
				tlog.Errorf("json.NewEncoder: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			return

		}

		// turn %20 and other HTML to normal string formatting
		name, err = url.QueryUnescape(name)
		if err != nil {
			tlog.Errorf("url.QueryUnescape: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		tlog.Debugf("search: name: %s", name)

		rows, err := store.QuestSearchByName(ctx, name)
		if err != nil {
			tlog.Errorf("library.QuestSearchByName: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		resp := Response{
			Code:    200,
			Message: fmt.Sprintf("There are %d quests found", len(rows)),
			Quests:  rows,
		}

		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			tlog.Errorf("json.NewEncoder: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handlers) ViewQuest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		quest := r.Context().Value(ContextKeyQuest).(*models.Quest)

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

		h.render(w, "quest", "view.go.tmpl", "content.go.tmpl", data)
	}
}

func (h *Handlers) GenerateQuestImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		quest := r.Context().Value(ContextKeyQuest).(*models.Quest)

		lines := []string{
			quest.Name,
			"Level: " + strconv.Itoa(quest.Level),
		}

		data, err := image.GenerateQuestPreview(quest.Icon, lines)
		if err != nil {
			h.serverErrorResponse(w, fmt.Errorf("GenerateQuestPreview: %w", err))
		}

		_, err = w.Write(data)
		if err != nil {
			h.logger.Error("error writing image", zap.Error(err))
		}
	}
}
