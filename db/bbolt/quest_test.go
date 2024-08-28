package bbolt

import (
	"context"
	"os"
	"testing"

	"github.com/xackery/tinywebeq/tlog"
)

func TestQuest(t *testing.T) {
	if os.Getenv("SINGLE_TEST") != "1" {
		return
	}
	b := New("../../bin/cache/bolt.db")
	ctx := context.Background()
	quest, err := b.QuestByQuestID(ctx, 110)
	if err != nil {
		t.Fatalf("quest by quest id: %s", err)
	}
	if quest == nil {
		t.Fatalf("quest by quest id: nil")
	}

	tlog.Infof("quest: %+v", quest)

}
