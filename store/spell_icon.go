package store

import (
	"context"
	"errors"
	"fmt"
	"image"
	"image/draw"
	"os"
	"sync"

	"github.com/ftrvxmtrx/tga"
	"github.com/xackery/tinywebeq/tlog"
)

var (
	spellIcons = sync.Map{} // map[int32]image.Image
)

func initSpellIcon(ctx context.Context) error {
	files := []string{
		"assets/spells01.tga",
		"assets/spells02.tga",
		"assets/spells03.tga",
		"assets/spells04.tga",
		"assets/spells05.tga",
		"assets/spells06.tga",
		"assets/spells07.tga",
	}

	index := 0
	isLoaded := false
	for _, file := range files {
		img, err := fetchTGA(file)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return nil
			}
			if isLoaded {
				return nil
			}
			return fmt.Errorf("fetchTGA: %w", err)
		}

		isEmpty := false
		for i := 0; i < 6; i++ {
			if isEmpty {
				break
			}
			for j := 0; j < 6; j++ {
				//subImg := img.(*image.NRGBA).SubImage(image.Rect(j*40, i*41, j*40+40, i*41+41))
				// move subimg pixels to 0,0
				iconImg := image.NewNRGBA(image.Rect(0, 0, 40, 40))
				//draw.Draw(iconImg, iconImg.Bounds(), subImg, image.Pt(0, 0), draw.Src)
				draw.Draw(iconImg, iconImg.Bounds(), img, image.Pt(j*40, i*40), draw.Src)

				if iconImg.At(0, 0) == image.Transparent {
					isEmpty = true
					break
				}

				spellIcons.Store(int32(index), iconImg)
				index++
			}
		}
		isLoaded = true
	}

	tlog.Debugf("Loaded %d spell icons", index)

	return nil
}

func fetchTGA(path string) (image.Image, error) {
	r, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("os.Open: %w", err)
	}
	defer r.Close()
	img, err := tga.Decode(r)
	if err != nil {
		return nil, fmt.Errorf("tga.Decode: %w", err)
	}

	return img, nil
}

// SpellIcon returns an image.Image for a spell icon
func SpellIcon(id int32) image.Image {
	rawImg, ok := spellIcons.Load(id)
	if !ok {
		return nil
	}
	img, ok := rawImg.(image.Image)
	if !ok {
		return nil
	}
	return img
}
