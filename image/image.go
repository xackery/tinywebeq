package image

import (
	"context"
	"fmt"
	"image"
	"image/color"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/tlog"
)

var (
	mu sync.RWMutex
)

func Init(ctx context.Context) error {
	mu.Lock()
	defer mu.Unlock()
	var err error

	focus := strings.ToUpper(config.Get().Item.Preview.FGColor)
	if focus != "" && focus != "#DBDEE1" {
		itemFGColor, err = hexColor(focus)
		if err != nil {
			return fmt.Errorf("item.Preview.FGColor %s: %w", focus, err)
		}
		itemFGImage = image.NewUniform(itemFGColor)
	}

	focus = strings.ToUpper(config.Get().Item.Preview.BGColor)
	if focus != "" && focus != "#313338" {
		itemBGColor, err = hexColor(focus)
		if err != nil {
			return fmt.Errorf("item.Preview.FGColor %s: %w", focus, err)
		}
		itemBGImage = image.NewUniform(itemBGColor)
	}

	focus = strings.ToUpper(config.Get().Spell.Preview.BGColor)
	if focus != "" && focus != "#DBDEE1" {
		spellBGColor, err = hexColor(focus)
		if err != nil {
			return fmt.Errorf("spell.BGColor %s: %w", focus, err)
		}
		spellBGImage = image.NewUniform(spellBGColor)
	}

	focus = strings.ToUpper(config.Get().Spell.Preview.FGColor)
	if focus != "" && focus != "#313338" {
		spellFGColor, err = hexColor(focus)
		if err != nil {
			return fmt.Errorf("spell.PreviewFGColor %s: %w", focus, err)
		}
		spellFGImage = image.NewUniform(spellFGColor)
	}

	focus = strings.ToUpper(config.Get().Npc.Preview.BGColor)
	if focus != "" && focus != "#DBDEE1" {
		npcBGColor, err = hexColor(focus)
		if err != nil {
			return fmt.Errorf("npc.Preview.BGColor %s: %w", focus, err)
		}
		npcBGImage = image.NewUniform(npcBGColor)
	}

	focus = strings.ToUpper(config.Get().Npc.Preview.FGColor)
	if focus != "" && focus != "#313338" {
		npcFGColor, err = hexColor(focus)
		if err != nil {
			return fmt.Errorf("npc.Preview.FGColor %s: %w", focus, err)
		}
		npcFGImage = image.NewUniform(npcFGColor)
	}

	focus = config.Get().Spell.Preview.FontNormal
	if focus != "" && focus != "goregular.ttf" {
		tlog.Debugf("Loading spellFont: %s", focus)
		spellFont, err = os.ReadFile(focus)
		if err != nil {
			return fmt.Errorf("read spellFont: %w", err)
		}
	}

	focus = config.Get().Spell.Preview.FontBold
	if focus != "" && focus != "gobold.ttf" {
		tlog.Debugf("Loading spellFontBold: %s", focus)
		spellFontBold, err = os.ReadFile(focus)
		if err != nil {
			return fmt.Errorf("read spellFontBold: %w", err)
		}
	}

	focus = config.Get().Item.Preview.FontNormal
	if focus != "" && focus != "goregular.ttf" {
		tlog.Debugf("Loading itemFont: %s", focus)
		itemFont, err = os.ReadFile(focus)
		if err != nil {
			return fmt.Errorf("read itemFont: %w", err)
		}
	}

	focus = config.Get().Item.Preview.FontBold
	if focus != "" && focus != "gobold.ttf" {
		tlog.Debugf("Loading itemFontBold: %s", focus)
		itemFontBold, err = os.ReadFile(focus)
		if err != nil {
			return fmt.Errorf("read itemFontBold: %w", err)
		}
	}

	focus = config.Get().Npc.Preview.FontNormal
	if focus != "" && focus != "goregular.ttf" {
		tlog.Debugf("Loading npcFont: %s", focus)
		npcFont, err = os.ReadFile(focus)
		if err != nil {
			return fmt.Errorf("read npcFont: %w", err)
		}
	}

	focus = config.Get().Npc.Preview.FontBold
	if focus != "" && focus != "gobold.ttf" {
		tlog.Debugf("Loading npcFontBold: %s", focus)
		npcFontBold, err = os.ReadFile(focus)
		if err != nil {
			return fmt.Errorf("read npcFontBold: %w", err)
		}
	}

	return nil
}

func hexColor(hex string) (color.RGBA, error) {
	values, err := strconv.ParseUint(string(hex[1:]), 16, 32)
	if err != nil {
		return color.RGBA{}, err
	}

	return color.RGBA{R: uint8(values >> 16), G: uint8((values >> 8) & 0xFF), B: uint8(values & 0xFF), A: 255}, nil
}
