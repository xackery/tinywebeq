package spell

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/go-jose/go-jose/v4/json"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/model"
	"github.com/xackery/tinywebeq/store"
	"github.com/xackery/tinywebeq/tlog"
)

// Search handles spell search requests
func Search(w http.ResponseWriter, r *http.Request) {
	var name string
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if !config.Get().Spell.Search.IsEnabled {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	tlog.Debugf("search: %s", r.URL.String())

	type Response struct {
		Code    int                  `json:"code"`
		Message string               `json:"message"`
		Spells  []*model.SpellSearch `json:"spells,omitempty"`
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
	results, err := store.SpellSearchByName(ctx, name)
	if err != nil {
		tlog.Errorf("library.SpellSearchByName: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	resp := Response{
		Code:    200,
		Message: fmt.Sprintf("There are %d spells found", len(results)),
		Spells:  results,
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		tlog.Errorf("json.NewEncoder: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
