package quest

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

// Preview handles quest preview requests
func PreviewImage(w http.ResponseWriter, r *http.Request) {
	var err error
	var id int64
	if !config.Get().Quest.IsEnabled {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	tlog.Debugf("previewImage: %s", r.URL.String())

	strID := r.URL.Query().Get("id")
	if len(strID) > 0 {
		id, err = strconv.ParseInt(strID, 10, 64)
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

func previewImageRender(ctx context.Context, id int64, w http.ResponseWriter) error {
	quest, err := store.QuestByQuestID(ctx, id)
	if err != nil {
		return fmt.Errorf("store.QuestByID: %w", err)
	}

	lines := []string{
		quest.Name,
		"Level: " + strconv.Itoa(quest.Level),
	}

	data, err := image.GenerateQuestPreview(quest.Icon, lines)
	if err != nil {
		return fmt.Errorf("GenerateQuestPreview: %w", err)
	}

	_, err = w.Write(data)
	if err != nil {
		return fmt.Errorf("w.Write: %w", err)
	}

	return nil
}
