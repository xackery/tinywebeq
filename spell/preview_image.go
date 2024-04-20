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
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tlog.Debugf("previewImageRender: id: %d", id)

	err = previewImageRender(ctx, id, w)
	if err != nil {
		tlog.Errorf("previewImageRender: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tlog.Debugf("previewImageRender: id: %d done", id)
}

func previewImageRender(ctx context.Context, id int, w http.ResponseWriter) error {

	lines := library.SpellInfo(id)

	data, err := image.GenerateSpellPreview(lines)
	if err != nil {
		return fmt.Errorf("GenerateSpellPreview: %w", err)
	}

	_, err = w.Write(data)
	if err != nil {
		return fmt.Errorf("w.Write: %w", err)
	}

	return nil
}
