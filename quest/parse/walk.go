package parse

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/xackery/tinywebeq/tlog"
)

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

	questNames := []string{}

	lineNumber := 0
	itemsCommentLineNumber := -1
	lineType := 0
	items := []*itemEntry{}
	isCommentBlock := false
	spawnNpcs := []int64{}
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
			questName := line
			questName = strings.ReplaceAll(questName, "--", "")
			questName = strings.ReplaceAll(questName, "#", "")
			questName = strings.TrimSpace(questName)
			if strings.HasPrefix(strings.ToLower(questName), "quest: ") {
				questName = questName[6:]
			}
			questName = strings.TrimSpace(questName)
			isNew := true
			for _, oldQuestName := range questNames {
				if oldQuestName != questName {
					continue
				}
				isNew = false
				break
			}
			if !isNew {
				continue
			}
			questNames = append(questNames, questName)
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

	dir := filepath.Dir(path)
	zoneShortName := filepath.Base(dir)
	npcName := strings.TrimSuffix(filepath.Base(path), language)

	resultChan <- &result{
		questNames:    questNames,
		zoneShortName: zoneShortName,
		npcName:       npcName,
		spawnNpcs:     spawnNpcs,
		items:         items,
		fileName:      path,
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

func findSpawn(line string, language string) int64 {
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
	id, err := strconv.ParseInt(line, 10, 64)
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
			return findItemIDs(int64(idx), line, "success")
		}
		idx = strings.Index(strings.ToLower(line), "check_turn_in(")
		if idx != -1 {
			idx += 14
			return findItemIDs(int64(idx), line, "component")
		}

	}
	if language == ".pl" {
		idx = strings.Index(strings.ToLower(line), "summonitem(")
		if idx != -1 {
			idx += 11
			return findItemIDs(int64(idx), line, "success")
		}
		idx = strings.Index(strings.ToLower(line), "check_handin(")
		if idx != -1 {
			idx += 14
			return findItemIDs(int64(idx), line, "component")
		}
	}
	return []*itemEntry{}
}

func findItemIDs(index int64, line string, useCase string) []*itemEntry {
	items := []*itemEntry{}
	itemMatches := regNumbers.FindAllStringSubmatch(line[index:], -1)
	for _, groups := range itemMatches {
		for _, match := range groups {
			id, err := strconv.ParseInt(match, 10, 64)
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
