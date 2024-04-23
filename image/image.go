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
	if focus != "" && focus != "#FFFFFF" {
		itemFGColor, err = hexColor(focus)
		if err != nil {
			return fmt.Errorf("item.Preview.FGColor %s: %w", focus, err)
		}
		itemFGImage = image.NewUniform(itemFGColor)
	}

	focus = strings.ToUpper(config.Get().Item.Preview.BGColor)
	if focus != "" && focus != "#303030" {
		itemBGColor, err = hexColor(focus)
		if err != nil {
			return fmt.Errorf("item.Preview.FGColor %s: %w", focus, err)
		}
		itemBGImage = image.NewUniform(itemBGColor)
	}

	focus = strings.ToUpper(config.Get().Spell.Preview.FGColor)
	if focus != "" && focus != "#FFFFFF" {
		spellBGColor, err = hexColor(focus)
		if err != nil {
			return fmt.Errorf("spell.PreviewFGColor %s: %w", focus, err)
		}
		spellBGImage = image.NewUniform(spellBGColor)
	}

	focus = strings.ToUpper(config.Get().Spell.Preview.BGColor)
	if focus != "" && focus != "#303030" {
		spellFGColor, err = hexColor(focus)
		if err != nil {
			return fmt.Errorf("spell.PreviewFGColor %s: %w", focus, err)
		}
		spellFGImage = image.NewUniform(spellFGColor)
	}

	focus = config.Get().Spell.Preview.FontNormal
	if focus != "" && focus != "goregular" {
		tlog.Debugf("Loading spellFont: %s", focus)
		spellFont, err = os.ReadFile(focus)
		if err != nil {
			return fmt.Errorf("read spellFont: %w", err)
		}
	}

	focus = config.Get().Spell.Preview.FontBold
	if focus != "" && focus != "gobold" {
		tlog.Debugf("Loading spellFontBold: %s", focus)
		spellFontBold, err = os.ReadFile(focus)
		if err != nil {
			return fmt.Errorf("read spellFontBold: %w", err)
		}
	}

	focus = config.Get().Item.Preview.FontNormal
	if focus != "" && focus != "goregular" {
		tlog.Debugf("Loading itemFont: %s", focus)
		itemFont, err = os.ReadFile(focus)
		if err != nil {
			return fmt.Errorf("read itemFont: %w", err)
		}
	}

	focus = config.Get().Item.Preview.FontBold
	if focus != "" && focus != "gobold" {
		tlog.Debugf("Loading itemFontBold: %s", focus)
		itemFontBold, err = os.ReadFile(focus)
		if err != nil {
			return fmt.Errorf("read itemFontBold: %w", err)
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
