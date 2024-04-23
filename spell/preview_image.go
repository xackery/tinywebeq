package spell

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/xackery/tinywebeq/image"
	"github.com/xackery/tinywebeq/tlog"

	"github.com/xackery/tinywebeq/library"
)

// Preview handles spell preview requests
func PreviewImage(w http.ResponseWriter, r *http.Request) {
	var err error
	var id int
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	tlog.Debugf("previewImage: %s", r.URL.String())

	strID := r.URL.Query().Get("id")
	if len(strID) > 0 {
		id, err = strconv.Atoi(strID)
		if err != nil {
			tlog.Errorf("strconv.Atoi: %v", err)
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tlog.Debugf("previewImageRender: id: %d", id)

	err = previewImageRender(ctx, id, w)
	if err != nil {
		tlog.Errorf("previewImageRender: %v", err)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	tlog.Debugf("previewImageRender: id: %d done", id)
}

func previewImageRender(ctx context.Context, id int, w http.ResponseWriter) error {
	spellIcon, lines := library.SpellInfo(id)
	if len(lines) == 0 {
		return fmt.Errorf("no spell info found")
	}

	data, err := image.GenerateSpellPreview(spellIcon, lines)
	if err != nil {
		return fmt.Errorf("GenerateSpellPreview: %w", err)
	}

	_, err = w.Write(data)
	if err != nil {
		return fmt.Errorf("w.Write: %w", err)
	}

	return nil
}
