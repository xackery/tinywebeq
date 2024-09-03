package parse

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sync"
	"time"

	"github.com/rs/zerolog"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/models"
	"github.com/xackery/tinywebeq/tlog"
)

var (
	regNumbers    = regexp.MustCompile("(?m)([0-9]+)")
	resultChan    = make(chan *result, 9999)
	failureChan   = make(chan error, 9999)
	totalCount    int
	mu            sync.RWMutex
	start         time.Time
	chunkStart    time.Time
	wg            sync.WaitGroup
	jobCount      int
	oldQuestIndex map[string]int64
	itemNames     map[int64]string
	zones         map[string]*models.Zone
	npcs          map[string][]int64
)

type itemEntry struct {
	ItemID   int64
	ItemName string
	UseCase  string
}

type result struct {
	questNames    []string
	zoneShortName string
	npcName       string
	spawnNpcs     []int64
	items         []*itemEntry
	fileName      string
}

func Init() error {
	if !config.Get().Quest.IsEnabled {
		tlog.Debugf("Quest scanning are disabled")
		return nil
	}

	if !config.Get().Quest.IsBackgroundScanningEnabled {
		tlog.Debugf("Quest background scanning is disabled")
		return nil
	}

	path := config.Get().Quest.Path
	if path == "" {
		return fmt.Errorf("quest path is empty")
	}

	fi, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("stat quest path: %s: %w", path, err)
	}
	if !fi.IsDir() {
		return fmt.Errorf("quest path is not a directory: %s", path)
	}

	go maintain()
	return nil
}

func maintain() {
	time.Sleep(10 * time.Second)

	tickerQuestCache := time.NewTicker(time.Duration(config.Get().Quest.ScanSchedule) * time.Second)

	for {
		select {
		case <-tickerQuestCache.C:
			err := Parse(context.Background(), config.Get().Quest.BackgroundScanConcurrency)
			if err != nil {
				tlog.Errorf("quest walk: %s", err)
			}
		}
	}
}

func Parse(ctx context.Context, concurrency int) error {
	if !config.Get().Quest.IsEnabled {
		tlog.Debugf("Quests are disabled")
		return nil
	}
	mu.Lock()
	defer mu.Unlock()

	tlog.Debugf("Setting log level to info")

	jobCount = concurrency

	tlog.SetLevel(zerolog.InfoLevel)

	tlog.Infof("Parsing quests at %s (this will take a while)", config.Get().Quest.Path)

	start = time.Now()
	chunkStart = time.Now()
	oldQuestIndex = map[string]int64{}
	resultChan = make(chan *result, 9999)
	path := config.Get().Quest.Path
	if path == "" {
		return fmt.Errorf("quest path is empty")
	}

	itemNames = make(map[int64]string)
	itemRows, err := db.Mysql.ItemsAll(ctx)
	if err != nil {
		return fmt.Errorf("items all: %w", err)
	}
	for _, row := range itemRows {
		itemNames[int64(row.ID)] = row.Name
	}

	quests, err := db.BBolt.QuestsAll(ctx)
	if err != nil {
		return fmt.Errorf("quest all: %w", err)
	}

	for _, quest := range quests {
		oldQuestIndex[quest.Name] = quest.ID
	}

	zones = make(map[string]*models.Zone)
	zoneRows, err := db.Mysql.ZonesAll(ctx)
	if err != nil {
		return fmt.Errorf("zones all: %w", err)
	}
	for i := range zoneRows {
		zones[zoneRows[i].ShortName] = zoneRows[i]
	}

	npcs = make(map[string][]int64)
	npcRows, err := db.Mysql.NpcsAll(ctx)
	if err != nil {
		return fmt.Errorf("npcs all: %w", err)
	}
	for _, row := range npcRows {
		npcs[row.Name] = append(npcs[row.Name], int64(row.ID))
	}

	tlog.Infof("Truncating quest cache")

	err = db.BBolt.QuestTruncate(ctx)
	if err != nil {
		return fmt.Errorf("quest truncate: %w", err)
	}

	err = db.BBolt.ItemQuestTruncate(ctx)
	if err != nil {
		return fmt.Errorf("item quest truncate: %w", err)
	}

	err = db.BBolt.NpcQuestTruncate(ctx)
	if err != nil {
		return fmt.Errorf("npc quest truncate: %w", err)
	}

	totalCount = 0

	// fi, err := os.Stat(path + "/dreadlands/Baldric_Slezaf.pl")
	// questParse(path+"/dreadlands/Baldric_Slezaf.pl", fi, nil)
	// os.Exit(0)

	tlog.Infof("Walking quests")

	err = filepath.Walk(path, questWalk)
	if err != nil {
		return fmt.Errorf("walk: %w", err)
	}

	wg.Wait()

	err = processResults()
	if err != nil {
		return fmt.Errorf("process results: %w", err)
	}

	if config.Get().IsDebugEnabled {
		tlog.SetLevel(zerolog.DebugLevel)
	}

	tlog.Infof("Parsed quests in %s", time.Since(start).String())
	return nil
}
