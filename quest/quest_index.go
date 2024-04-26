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
	regNumbers    = regexp.MustCompile("(?m)([0-9]+)")
	failureChan   = make(chan error, 9999)
	totalCount    int
	mu            sync.RWMutex
	start         time.Time
	chunkStart    time.Time
	wg            sync.WaitGroup
	jobCount      int
	oldQuestIndex = map[string]int{} // to help assist with retaining unique id's
)

type itemEntry struct {
	ItemID  int
	UseCase string
}

func questInit() error {

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
	oldQuestIndex = make(map[string]int)
	path := config.Get().Quest.Path
	if path == "" {
		return fmt.Errorf("quest path is empty")
	}

	tlog.Infof("Truncating quest cache")
	oldQuests, err := cache.DumpSqlite(ctx, "quest")
	if err != nil {
		if err.Error() != cache.ErrCacheNotFound {
			return fmt.Errorf("dump quest cache: %w", err)
		}
	} else {
		for _, oldQuest := range oldQuests {
			oldQuestIndex[oldQuest.(*model.Quest).Name] = oldQuest.(*model.Quest).ID
		}
		err = cache.TruncateSqliteCache(ctx, "quest")
		if err != nil {
			return fmt.Errorf("truncate sqlite cache quest: %w", err)
		}
	}

	err = cache.TruncateSqliteCache(ctx, "item_quest")
	if err != nil {
		return fmt.Errorf("truncate sqlite cache item_quest: %w", err)
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

	questName := ""

	lineNumber := 0
	itemsCommentLineNumber := -1
	lineType := 0
	items := []*itemEntry{}
	isCommentBlock := false
	spawnNpcs := []int{}
	for {
		lineNumber++
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}

		if line[len(line)-1] == '\n' {
			line = line[0 : len(line)-1]
		}

		line = strings.TrimSpace(line)

		isComment := false
		if strings.Contains(line, "--[[") {
			isCommentBlock = true
			isComment = false
		}
		if isCommentBlock && strings.Contains(line, "]]---") {
			isCommentBlock = false
		}
		if strings.HasPrefix(line, "#") {
			isComment = true
		}

		if isCommentBlock {
			isComment = true
		}

		if !isCommentBlock && lineType == 0 && itemsLineType(line, language) > 0 {
			itemsCommentLineNumber = lineNumber
			lineType = itemsLineType(line, language)
		}

		if !isCommentBlock && itemsCommentLineNumber == -1 && firstCharacter(line) != commentMarker(language) {
			itemsCommentLineNumber = lineNumber
		}

		if isComment && strings.Contains(strings.ToLower(line), "quest: ") {
			questName = line
			questName = strings.ReplaceAll(questName, "--[[", "")
			questName = strings.ReplaceAll(questName, "]]---", "")
			questName = strings.ReplaceAll(questName, "--", "")
			questName = strings.ReplaceAll(questName, "#", "")
			questName = strings.TrimSpace(questName)
			if strings.HasPrefix(strings.ToLower(questName), "quest: ") {
				questName = questName[6:]
			}
			questName = strings.TrimSpace(questName)
		}

		if !isComment {
			spawnNpcID := findSpawn(line, language)
			if spawnNpcID > 0 {
				isNew := true
				for _, oldNpcID := range spawnNpcs {
					if oldNpcID != spawnNpcID {
						continue
					}
					isNew = false
					break
				}
				if isNew {
					spawnNpcs = append(spawnNpcs, spawnNpcID)
				}
			}
		}

		possibleItems := findItems(line, language)
		for _, newItem := range possibleItems {
			isNew := true
			for _, oldItem := range items {
				if oldItem.ItemID != newItem.ItemID {
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

	zone := library.ZoneByID(library.ZoneIDByShortName(zoneShortName))
	if zone == nil {
		return fmt.Errorf("zone %s not found", zoneShortName)
	}
	zoneID := zone.ZoneIDNumber
	expansion := zone.Expansion

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

	if questName == "" {
		questName = npcName
		questName = strings.ReplaceAll(questName, "-", "`")
		questName = strings.ReplaceAll(questName, "_", " ")
		questName = strings.ReplaceAll(questName, "#", "")
	}

	questID := 0
	if oldQuestIndex[questName] > 0 {
		questID = oldQuestIndex[questName]
	}

	if questID == 0 {
		questID, err = cache.NextIDSqlite(ctx, "quest")
		if err != nil {
			return fmt.Errorf("next id: %w", err)
		}
	}

	oldQuestIndex[questName] = questID

	for _, spawnNpcID := range spawnNpcs {
		err = writeNpcQuest(ctx, questID, zoneID, spawnNpcID, questName)
		if err != nil {
			return fmt.Errorf("write npc quest: %w", err)
		}
	}

	err = writeQuest(ctx, questID, zoneID, npcInfo.ID, questName, npcName, path, expansion, items)
	if err != nil {
		return fmt.Errorf("write quest: %w", err)
	}

	err = writeItemQuest(ctx, questID, zoneID, npcInfo.ID, questName, npcName, expansion, items)
	if err != nil {
		return fmt.Errorf("write item quest: %w", err)
	}

	return nil
}

func writeNpcQuest(ctx context.Context, questID int, zoneID int, npcID int, questName string) error {
	path := fmt.Sprintf("npc_quest/%d.yaml", npcID)

	var npcQuest *model.NpcQuest
	npcQuestCache, ok := cache.ReadSqlite(path)
	if ok {
		npcQuest = npcQuestCache.(*model.NpcQuest)
	}
	if npcQuest == nil {
		npcQuest = &model.NpcQuest{
			ID: npcID,
		}
	}
	entry := &model.NpcQuestEntry{
		QuestID:   questID,
		QuestName: questName,
		ZoneID:    zoneID,
	}
	npcQuest.Entries = append(npcQuest.Entries, entry)
	err := cache.WriteSqlite(ctx, path, npcQuest)
	if err != nil {
		return fmt.Errorf("cache write: %w", err)
	}

	return nil
}

func writeQuest(ctx context.Context, questID int, zoneID int, npcID int, questName string, npcName string, fileName string, expansion int, items []*itemEntry) error {
	path := fmt.Sprintf("quest/%d.yaml", questID)

	var quest *model.Quest
	questCache, ok := cache.ReadSqlite(path)
	if ok {
		quest = questCache.(*model.Quest)
	}
	if quest == nil {
		quest = &model.Quest{
			ID:   questID,
			Name: questName,
		}
	}
	for _, item := range items {
		if item.ItemID < 1001 {
			continue
		}
		entry := &model.QuestEntry{
			ItemID:    item.ItemID,
			Score:     0,
			FileName:  fileName,
			NpcID:     npcID,
			NpcName:   npcName,
			ZoneID:    zoneID,
			UseCase:   item.UseCase,
			Expansion: expansion,
		}
		quest.Entries = append(quest.Entries, entry)
	}
	err := cache.WriteSqlite(ctx, path, quest)
	if err != nil {
		return fmt.Errorf("cache write: %w", err)
	}

	return nil
}

func writeItemQuest(ctx context.Context, questID int, zoneID int, npcID int, questName string, npcName string, expansion int, items []*itemEntry) error {

	for _, item := range items {
		if item.ItemID < 1001 {
			continue
		}
		path := fmt.Sprintf("item_quest/%d.yaml", item.ItemID)

		var itemQuest *model.ItemQuest
		itemQuestCache, ok := cache.ReadSqlite(path)
		if ok {
			itemQuest = itemQuestCache.(*model.ItemQuest)
		}
		if itemQuest == nil {
			itemQuest = &model.ItemQuest{
				ID: questID,
			}
		}

		entry := &model.ItemQuestEntry{
			QuestID:   questID,
			QuestName: questName,
			ItemID:    item.ItemID,
			NpcID:     npcID,
			NpcName:   npcName,
			ZoneID:    zoneID,
			UseCase:   item.UseCase,
		}

		itemQuest.Entries = append(itemQuest.Entries, entry)

		err := cache.WriteSqlite(ctx, path, itemQuest)
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

func findSpawn(line string, language string) int {
	line = strings.ToLower(line)
	if !strings.Contains(line, "spawn2(") {
		return 0
	}

	line = line[strings.Index(line, "spawn2(")+7:]
	line = line[:strings.Index(line, ")")]
	if strings.Contains(line, ",") {
		line = line[:strings.Index(line, ",")]
	}

	line = strings.TrimSpace(line)
	id, err := strconv.Atoi(line)
	if err != nil {
		return 0
	}
	return id
}

func findItems(line string, language string) []*itemEntry {

	idx := -1
	if language == ".lua" {
		idx = strings.Index(strings.ToLower(line), "summonitem(")
		if idx != -1 {
			idx += 11
			return findItemIDs(idx, line, "success")
		}
		idx = strings.Index(strings.ToLower(line), "check_turn_in(")
		if idx != -1 {
			idx += 14
			return findItemIDs(idx, line, "component")
		}

	}
	if language == ".pl" {
		idx = strings.Index(strings.ToLower(line), "summonitem(")
		if idx != -1 {
			idx += 11
			return findItemIDs(idx, line, "success")
		}
		idx = strings.Index(strings.ToLower(line), "check_handin(")
		if idx != -1 {
			idx += 14
			return findItemIDs(idx, line, "component")
		}
	}
	return []*itemEntry{}
}

func findItemIDs(index int, line string, useCase string) []*itemEntry {
	items := []*itemEntry{}
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

			items = append(items, &itemEntry{
				ItemID:  id,
				UseCase: useCase,
			})
		}
	}
	return items
}
