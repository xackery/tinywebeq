package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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

func (h *Handlers) IndexZones() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var name string
		var err error
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if !config.Get().Zone.Search.IsEnabled {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		tlog.Debugf("search: %s", r.URL.String())

		type Response struct {
			Code    int                  `json:"code"`
			Message string               `json:"message"`
			Zones   []*models.ZoneSearch `json:"zones,omitempty"`
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
		results, err := store.ZoneSearchByName(ctx, name)
		if err != nil {
			tlog.Errorf("library.ZoneSearchByName: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		resp := Response{
			Code:    200,
			Message: fmt.Sprintf("There are %d zones found", len(results)),
			Zones:   results,
		}

		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			tlog.Errorf("json.NewEncoder: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handlers) ViewZone() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		zone := r.Context().Value(ContextKeyZone).(*models.Zone)

		data := struct {
			Site                site.BaseData
			Zone                *models.Zone
			Library             *library.Library
			ZoneInfo            []string
			IsZoneSearchEnabled bool
		}{
			Site:                site.BaseDataInit(zone.LongName),
			Zone:                zone,
			Library:             library.Instance(),
			IsZoneSearchEnabled: config.Get().Zone.Search.IsEnabled,
		}

		h.render(w, "zone", "view.go.tmpl", "content.go.tmpl", data)
	}
}

func (h *Handlers) GenerateZoneImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		zone := r.Context().Value(ContextKeyZone).(*models.Zone)

		lines := []string{
			zone.LongName,
			fmt.Sprintf("ID: %d", zone.Zoneidnumber),
		}

		data, err := image.GenerateZonePreview(zone.Icon, lines)
		if err != nil {
			h.serverErrorResponse(w, fmt.Errorf("GenerateZonePreview: %w", err))
			return
		}

		_, err = w.Write(data)
		if err != nil {
			h.logger.Error("error writing image", zap.Error(err))
		}
	}
}
