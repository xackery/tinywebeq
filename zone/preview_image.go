package zone

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/image"
	"github.com/xackery/tinywebeq/store"
	"github.com/xackery/tinywebeq/tlog"
)

// Preview handles zone preview rezones
func PreviewImage(w http.ResponseWriter, r *http.Request) {
	var err error
	var id int
	if !config.Get().Zone.IsEnabled {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
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
	zone, err := store.ZoneByZoneIDNumber(ctx, int64(id))
	if err != nil {
		return fmt.Errorf("store.ZoneByZoneIDNumber: %w", err)
	}

	lines := []string{
		zone.LongName,
		fmt.Sprintf("ID: %d", zone.Zoneidnumber),
	}

	data, err := image.GenerateZonePreview(zone.Icon, lines)
	if err != nil {
		return fmt.Errorf("GenerateZonePreview: %w", err)
	}

	_, err = w.Write(data)
	if err != nil {
		return fmt.Errorf("w.Write: %w", err)
	}

	return nil
}
