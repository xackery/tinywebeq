package quest

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/xackery/tinywebeq/cache"
	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/library"
	"github.com/xackery/tinywebeq/model"
	"github.com/xackery/tinywebeq/npc"
	"github.com/xackery/tinywebeq/tlog"
)

var (
	regNumbers  = regexp.MustCompile("(?m)([0-9]+)")
	failureChan = make(chan error, 9999)
	totalCount  int
	mu          sync.RWMutex
	start       time.Time
	chunkStart  time.Time
	wg          sync.WaitGroup
	jobCount    int
)

func Init() error {

	if !config.Get().Quest.IsEnabled {
		tlog.Debugf("Quest scanning are disabled")
		return nil
	}

	if !config.Get().Quest.IsBackgroundScanningEnabled {
		tlog.Debugf("Quest background scanning is disabled")
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
	path := config.Get().Quest.Path
	if path == "" {
		return fmt.Errorf("quest path is empty")
	}

	err := cache.TruncateSqliteCache(ctx, "item_quest")
	if err != nil {
		return fmt.Errorf("truncate sqlite cache: %w", err)
	}

	totalCount = 0
	err = filepath.Walk(path, questWalk)
	if err != nil {
		return fmt.Errorf("walk: %w", err)
	}

	wg.Wait()

	if config.Get().IsDebugEnabled {
		tlog.SetLevel(zerolog.DebugLevel)
	}

	close(failureChan)
	for err := range failureChan {
		tlog.Errorf("%s", err)
	}

	tlog.Infof("Parsed quests in %s", time.Since(start).String())
	return nil
}

func questWalk(path string, info os.FileInfo, err error) error {

	totalCount++
	if totalCount%jobCount == 0 {
		wg.Wait()
		if totalCount%1000 == 0 {
			tlog.Infof("Parsed %d quests in %s (%s total)", totalCount, time.Since(chunkStart).String(), time.Since(start).String())
			chunkStart = time.Now()
		}
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		err = questParse(path, info, err)
		if err != nil {
			failureChan <- err
			err = nil
		}
	}()
	return nil
}

func questParse(path string, info os.FileInfo, err error) error {
	if err != nil {
		return fmt.Errorf("walk %q: %w", path, err)
	}
	if info.IsDir() {
		return nil
	}
	language := strings.ToLower(filepath.Ext(path))
	if language != ".lua" && language != ".pl" {
		return nil
	}
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	r := bufio.NewReader(f)

	lineNumber := 0
	itemsCommentLineNumber := -1
	lineType := 0
	items := []string{}
	isCommentBlock := false
	for {
		lineNumber++
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}

		if line[len(line)-1] == '\n' {
			line = line[0 : len(line)-1]
		}

		if strings.Contains(line, "--[[") {
			isCommentBlock = true
		}
		if isCommentBlock && strings.Contains(line, "]]---") {
			isCommentBlock = false
		}

		if !isCommentBlock && lineType == 0 && itemsLineType(line, language) > 0 {
			itemsCommentLineNumber = lineNumber
			lineType = itemsLineType(line, language)
		}

		if !isCommentBlock && itemsCommentLineNumber == -1 && firstCharacter(line) != commentMarker(language) {
			itemsCommentLineNumber = lineNumber
		}

		possibleItems := findItems(line, language)
		for _, newItem := range possibleItems {
			isNew := true
			for _, oldItem := range items {
				if oldItem != newItem {
					continue
				}
				isNew = false
				break
			}
			if !isNew {
				continue
			}
			items = append(items, newItem)
		}
	}
	if len(items) == 0 {
		return nil
	}

	dir := filepath.Dir(path)
	zoneShortName := filepath.Base(dir)
	npcName := strings.TrimSuffix(filepath.Base(path), language)

	zoneID := library.ZoneIDByShortName(zoneShortName)

	tlog.Debugf("path: %s npcName: %s zoneShortName: %s zoneID: %d items: %s", path, npcName, zoneShortName, zoneID, items)
	ctx := context.Background()

	var npcInfo *model.Npc
	npcID, err := strconv.Atoi(npcName)
	if err != nil {
		npcInfo, err = npc.FetchNpcByName(ctx, npcName, zoneID)
		if err != nil {
			return fmt.Errorf("fetch npc %s zoneID %d by name: %w", npcName, zoneID, err)
		}
	} else {
		npcInfo, err = npc.FetchNpc(ctx, npcID)
		if err != nil {
			return fmt.Errorf("fetch npc: %w", err)
		}
	}

	err = writeItemQuest(ctx, zoneID, npcInfo.ID, npcName, items)
	if err != nil {
		return fmt.Errorf("write item quest: %w", err)
	}

	return nil
}

func writeItemQuest(ctx context.Context, zoneID int, npcID int, npcName string, items []string) error {

	for _, item := range items {
		itemID, err := strconv.Atoi(item)
		if err != nil {
			return fmt.Errorf("strconv.Atoi: %w", err)
		}
		path := fmt.Sprintf("item_quest/%d.yaml", itemID)

		var itemQuest *model.ItemQuest
		itemQuestCache, ok := cache.ReadSqlite(path)
		if ok {
			itemQuest = itemQuestCache.(*model.ItemQuest)
		}
		if itemQuest == nil {
			itemQuest = &model.ItemQuest{}
		}

		entry := &model.ItemQuestEntry{
			ItemID:  itemID,
			NpcID:   npcID,
			NpcName: npcName,
			ZoneID:  zoneID,
			UseCase: "quest",
		}

		itemQuest.Entries = append(itemQuest.Entries, entry)

		err = cache.WriteSqlite(ctx, path, itemQuest)
		if err != nil {
			return fmt.Errorf("cache write: %w", err)
		}
	}

	return nil
}

// itemsLineType has 3 possible values: 0 (none), 1: (items: generated line), 2: (!items: manually edited line to skip)
func itemsLineType(line string, language string) int {
	if !isComment(line, language) {
		return 0
	}
	commentMark := commentMarker(language)
	idx := strings.Index(line, string(commentMark))
	if idx == -1 {
		return 0
	}
	line = strings.ReplaceAll(line, string(commentMark), "")
	line = strings.TrimSpace(line)
	idx = strings.Index(line, "items:")
	if idx == 0 {
		return 1
	}
	idx = strings.Index(line, "!items:")
	if idx == 0 {
		return 2
	}
	return 0
}

func isComment(line string, language string) bool {
	return firstCharacter(line) == commentMarker(language)
}

// commentMarker returns either // or -- based on the language provided
func commentMarker(language string) string {
	if language == ".lua" {
		return "-"
	}
	if language == ".pl" {
		return "#"
	}
	return "-"
}

// firstCharacter returns the first valid character detected on a line
func firstCharacter(line string) string {
	line = strings.TrimSpace(line)
	if len(line) == 0 {
		return ""
	}
	return line[0:1]
}

func findItems(line string, language string) []string {

	idx := -1
	if language == ".lua" {
		idx = strings.Index(strings.ToLower(line), "summonitem(")
		if idx != -1 {
			idx += 11
			return findItemIDs(idx, line)
		}
		idx = strings.Index(strings.ToLower(line), "check_turn_in(")
		if idx != -1 {
			idx += 14
			return findItemIDs(idx, line)
		}

	}
	if language == ".pl" {
		idx = strings.Index(strings.ToLower(line), "summonitem(")
		if idx != -1 {
			idx += 11
			return findItemIDs(idx, line)
		}
		idx = strings.Index(strings.ToLower(line), "check_handin(")
		if idx != -1 {
			idx += 14
			return findItemIDs(idx, line)
		}
	}
	return []string{}
}

func findItemIDs(index int, line string) []string {
	items := []string{}
	itemMatches := regNumbers.FindAllStringSubmatch(line[index:], -1)
	for _, groups := range itemMatches {
		for _, match := range groups {
			id, err := strconv.Atoi(match)
			if err != nil {
				continue
			}
			if id < 1000 {
				continue
			}
			items = append(items, match)
		}
	}
	return items
}
