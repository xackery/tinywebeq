package store

import (
	"context"
	"fmt"
	"image"
	"image/draw"
	"os"
	"sync"

	"github.com/xackery/tinywebeq/tlog"
)

var (
	npcIcons = sync.Map{} // map[int32]image.Image
)

func initNpcIcon(ctx context.Context) error {
	var err error

	err = os.MkdirAll("assets", os.ModePerm)
	if err != nil {
		return fmt.Errorf("os.MkdirAll: %w", err)
	}

	err = initNpcIcons()
	if err != nil {
		if os.IsNotExist(err) {
			tlog.Warnf("initNpcIcons: %v", err)
		}
		tlog.Warnf("%v", err)
		tlog.Infof("To add npc icons, copy uifiles/default/dragnpc*.dds to the assets folder")
	}
	return nil
}

func initNpcIcons() error {
	files := []string{}
	for i := 1; i < 179; i++ {
		files = append(files, fmt.Sprintf("assets/dragnpc%d.dds", i))
	}

	count := 0
	index := int32(500)
	for _, file := range files {
		img, err := fetchDDS(file)
		if err != nil {
			return fmt.Errorf("fetchDDS: %w", err)
		}

		for x := 0; x+40 <= img.Bounds().Dx(); x += 40 {
			for y := 0; y+40 <= img.Bounds().Dy(); y += 40 {
				//subImg := img.(*image.NRGBA).SubImage(image.Rect(j*40, i*41, j*40+40, i*41+41))
				// move subimg pixels to 0,0
				iconImg := image.NewRGBA(image.Rect(0, 0, 40, 40))
				//draw.Draw(iconImg, iconImg.Bounds(), subImg, image.Pt(0, 0), draw.Src)
				draw.Draw(iconImg, iconImg.Bounds(), img, image.Pt(x, y), draw.Src)

				npcIcons.Store(index, iconImg)
				index++
				count++
			}
		}
	}

	tlog.Debugf("Loaded %d npc icons", count)
	return nil
}

func NpcIcon(id int32) image.Image {
	rawImg, ok := npcIcons.Load(id)
	if !ok {
		return nil
	}
	img, ok := rawImg.(image.Image)
	if !ok {
		return nil
	}
	return img
}
