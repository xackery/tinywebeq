package handlers

import (
	"net/http"

	"github.com/xackery/tinywebeq/models"
	"github.com/xackery/tinywebeq/site"
)

func (h *Handlers) ViewPlayer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		player := r.Context().Value(ContextKeyPlayer).(*models.Player)

		data := struct {
			Site   site.BaseData
			Player *models.Player
		}{
			Site:   site.BaseDataInit("Player View"),
			Player: player,
		}

		h.render(w, "player", "view.go.tmpl", "content.go.tmpl", data)
	}
}
