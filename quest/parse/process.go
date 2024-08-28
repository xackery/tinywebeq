package parse

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/model"
	"github.com/xackery/tinywebeq/store"
	"github.com/xackery/tinywebeq/tlog"
)

var (
	successQuests = 0
	successNPCs   = 0
	successItems  = 0
	zonesSkipped  = []string{}
)

func processResults() error {
	tlog.Infof("Processing results")
	close(resultChan)
	count := 0
	startChunk := time.Now()
	w, err := os.Create("quest_parse.log")
	if err != nil {
		return fmt.Errorf("create quest parse log: %w", err)
	}
	defer w.Close()

	failureCount := 0

	for result := range resultChan {
		count++

		if count%1000 == 0 {
			tlog.Infof("Processed %d results in %s", count, time.Since(startChunk).String())
			startChunk = time.Now()
		}
		err := processResult(result)
		if err != nil {

			w.WriteString(fmt.Sprintf("zone: %s, file: %s, %s\n", result.zoneShortName, filepath.Base(result.fileName), err))
			failureCount++
			continue
		}
	}
	tlog.Infof("Process success: %d quests, %d npcs, %d items", successQuests, successNPCs, successItems)
	tlog.Infof("Process failure: %d written to quest_parse.log", failureCount)
	return nil
}

func processResult(result *result) error {
	questNames := result.questNames

	npcName := result.npcName
	zoneShortName := result.zoneShortName

	zone, ok := zones[zoneShortName]
	if !ok {
		for _, zoneName := range zonesSkipped {
			if zoneName == zoneShortName {
				return nil
			}
		}

		zonesSkipped = append(zonesSkipped, zoneShortName)
		return fmt.Errorf("zone %s not found", zoneShortName)
	}
	zoneID := zone.Zoneidnumber
	expansion := zone.Expansion
	spawnNpcs := result.spawnNpcs
	items := result.items

	//	tlog.Infof("npcName: %s zoneShortName: %s zoneID: %d items: %s", npcName, zoneShortName, zoneID, items)
	ctx := context.Background()

	npcID, err := strconv.ParseInt(npcName, 10, 64)
	if err != nil {
		npcName = strings.ReplaceAll(npcName, "-", "`")
		npcName = strings.ReplaceAll(npcName, " ", "_")

		if npcName == "player" {
			return nil
		}
		if npcName == "global" {
			return nil
		}

		npcIDs, ok := npcs[npcName]
		if !ok {
			return fmt.Errorf("npc %s not found", npcName)
		}

		isFound := false
		for _, npcID := range npcIDs {
			if npcID < int64(zoneID)*1000 || npcID > int64(zoneID)*1000+999 {
				continue
			}
			isFound = true
			break
		}
		if !isFound {
			return fmt.Errorf("npc %s not found in zone %s (%d) (candidates: %v)", npcName, zoneShortName, zoneID, npcIDs)
		}
	}

	if len(questNames) == 0 {
		questName := npcName
		questNames = append(questNames, questName)
	}

	for _, questName := range questNames {
		questID := int64(0)

		questName = strings.ReplaceAll(questName, "-", "`")
		questName = strings.ReplaceAll(questName, "_", " ")
		questName = strings.ReplaceAll(questName, "#", "")
		questName = strings.TrimSpace(questName)

		questID = oldQuestIndex[questName]

		if questID == 0 {
			questID, err = store.QuestNextQuestID(ctx)
			if err != nil {
				tlog.Warnf("quest next id: %s", err)
				os.Exit(1)
				return fmt.Errorf("quest next id: %w", err)
			}
		}

		oldQuestIndex[questName] = questID

		for _, spawnNpcID := range spawnNpcs {
			//tlog.Infof("Writing npc quest %s for npc %s in zone %s", questName, npcName, zoneShortName)
			err = writeNpcQuest(ctx, questID, zoneID, spawnNpcID, questName)
			if err != nil {
				return fmt.Errorf("write npc quest: %w", err)
			}
			successNPCs++
		}

		for _, item := range items {
			var ok bool
			item.ItemName, ok = itemNames[item.ItemID]
			if !ok {
				return fmt.Errorf("item name %d not found", item.ItemID)
			}
			//tlog.Infof("Writing quest %s for npc %s in zone %s", questName, npcName, zoneShortName)
			err = writeQuest(ctx, questID, zoneID, npcID, questName, npcName, result.fileName, expansion, item)
			if err != nil {
				return fmt.Errorf("write quest: %w", err)
			}

			successQuests++

			err = writeItemQuest(ctx, questID, zoneID, npcID, questName, npcName, expansion, item)
			if err != nil {
				return fmt.Errorf("write item quest: %w", err)
			}
			successItems++
		}
	}
	return nil
}

func writeNpcQuest(ctx context.Context, questID int64, zoneID int32, npcID int64, questName string) error {

	npcQuest, err := db.BBolt.NpcQuestByNpcID(ctx, npcID)
	if err != nil {
		if !strings.Contains(err.Error(), "no rows in result set") {
			return fmt.Errorf("npc quest by npc id: %w", err)
		}
		npcQuest = &model.NpcQuest{
			ID: npcID,
		}
	}

	entry := &model.NpcQuestEntry{
		QuestID:   questID,
		QuestName: questName,
		ZoneID:    zoneID,
	}
	npcQuest.SetExpiration(time.Now().Add(365 * 24 * time.Hour).Unix())

	npcQuest.Entries = append(npcQuest.Entries, entry)

	err = db.BBolt.NpcQuestReplace(ctx, npcID, npcQuest)
	if err != nil {
		return fmt.Errorf("npc quest replace: %w", err)
	}

	return nil
}

func writeQuest(ctx context.Context, questID int64, zoneID int32, npcID int64, questName string, npcName string, fileName string, expansion int8, item *itemEntry) error {
	if item.ItemID < 1001 {
		return nil
	}

	quest, err := db.BBolt.QuestByQuestID(ctx, questID)
	if err != nil {
		if !strings.Contains(err.Error(), "no rows in result set") {
			return fmt.Errorf("quest by id: %w", err)
		}
		quest = &model.Quest{
			ID:   questID,
			Name: questName,
		}
	}

	entry := &model.QuestEntry{
		ItemID:    item.ItemID,
		ItemName:  item.ItemName,
		Score:     0,
		FileName:  fileName,
		NpcID:     npcID,
		NpcName:   npcName,
		ZoneID:    zoneID,
		UseCase:   item.UseCase,
		Expansion: expansion,
	}
	quest.Entries = append(quest.Entries, entry)

	err = db.BBolt.QuestReplace(ctx, questID, quest)
	if err != nil {
		return fmt.Errorf("quest replace: %w", err)
	}

	return nil
}

func writeItemQuest(ctx context.Context, questID int64, zoneID int32, npcID int64, questName string, npcName string, expansion int8, item *itemEntry) error {
	if item.ItemID < 1001 {
		return nil
	}

	itemQuest, err := db.BBolt.ItemQuestByItemID(ctx, item.ItemID)
	if err != nil {
		if !strings.Contains(err.Error(), "no rows in result set") {
			return fmt.Errorf("item quest by item id: %w", err)
		}
		itemQuest = &model.ItemQuest{
			ItemID: item.ItemID,
		}
	}

	entry := &model.ItemQuestEntry{
		QuestID:   questID,
		QuestName: questName,
		ItemID:    item.ItemID,
		ItemName:  item.ItemName,
		NpcID:     npcID,
		NpcName:   npcName,
		ZoneID:    zoneID,
		UseCase:   item.UseCase,
	}

	itemQuest.Entries = append(itemQuest.Entries, entry)

	err = db.BBolt.ItemQuestReplace(ctx, item.ItemID, itemQuest)
	if err != nil {
		return fmt.Errorf("item quest replace: %w", err)
	}

	return nil
}
