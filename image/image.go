package image

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"strings"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/xackery/tinywebeq/item"
	"golang.org/x/image/font/gofont/goregular"
)

func GenerateItemPreview(item *item.Table) ([]byte, error) {
	fgColor := color.RGBA{0xff, 0xff, 0xff, 0xff}

	bgColor := color.RGBA{0x30, 0x0a, 0x24, 0xff}
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return nil, fmt.Errorf("parse font: %w", err)
	}
	fontSize := float64(14)

	fg := image.NewUniform(fgColor)
	bg := image.NewUniform(bgColor)
	rgba := image.NewRGBA(image.Rect(0, 0, 500, 300))
	draw.Draw(rgba, rgba.Bounds(), bg, image.Pt(0, 0), draw.Src)
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(font)
	c.SetFontSize(fontSize)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)
	//c.SetHinting(font.H)

	textXOffset := 50
	textYOffset := 10 + int(c.PointToFixed(fontSize)>>6) // Note shift/truncate 6 bits first

	text := []string{"test", "test2"}
	pt := freetype.Pt(textXOffset, textYOffset)
	for _, s := range text {
		_, err = c.DrawString(strings.Replace(s, "\r", "", -1), pt)
		if err != nil {
			return nil, err
		}
		pt.Y += c.PointToFixed(fontSize * 1.5)
	}

	b := new(bytes.Buffer)
	if err := png.Encode(b, rgba); err != nil {
		log.Println("unable to encode image.")
		return nil, err
	}
	return b.Bytes(), nil
}
