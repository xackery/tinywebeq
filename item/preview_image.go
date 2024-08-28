package item

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/xackery/tinywebeq/image"
	"github.com/xackery/tinywebeq/store"
	"github.com/xackery/tinywebeq/tlog"
)

// Preview handles item preview requests
func PreviewImage(w http.ResponseWriter, r *http.Request) {
	var err error
	var id int64
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	tlog.Debugf("previewImage: %s", r.URL.String())

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

	tlog.Debugf("previewImageRender: id: %d", id)

	err = previewImageRender(ctx, id, w)
	if err != nil {
		tlog.Errorf("previewImageRender: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tlog.Debugf("previewImageRender: id: %d done", id)
}

func previewImageRender(ctx context.Context, id int64, w http.ResponseWriter) error {
	if id <= 1000 {
		return fmt.Errorf("invalid id")
	}

	item, err := store.ItemByItemID(ctx, id)
	if err != nil {
		return fmt.Errorf("store.ItemByItemID: %w", err)
	}

	itemQuest, err := store.ItemQuestByItemID(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		tlog.Debugf("Ignoring err store.ItemQuestByItemID: %v", err)
	}

	itemRecipe, err := store.ItemRecipeByItemID(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		tlog.Debugf("Ignoring err store.ItemRecipeByItemID: %v", err)
	}

	data, err := image.GenerateItemPreview(item, itemQuest, itemRecipe)
	if err != nil {
		return fmt.Errorf("GenerateItemPreview: %w", err)
	}

	_, err = w.Write(data)
	if err != nil {
		return fmt.Errorf("w.Write: %w", err)
	}

	return nil
}
