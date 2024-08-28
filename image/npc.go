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
	"github.com/xackery/tinywebeq/store"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/math/fixed"
)

var (
	npcFGColor  = color.RGBA{0xdb, 0xde, 0xe1, 0xff}
	npcBGColor  = color.RGBA{0x31, 0x33, 0x38, 0xff}
	npcFGImage  = image.NewUniform(npcFGColor)
	npcBGImage  = image.NewUniform(npcBGColor)
	npcFont     = goregular.TTF
	npcFontBold = gobold.TTF
)

type NpcPreview struct {
	c           *freetype.Context
	cb          *freetype.Context
	cbFont      *truetype.Font
	cTitle      *freetype.Context
	fontSize    float64
	pt          fixed.Point26_6
	lineCurrent int
	maxHeight   int
	maxWidth    int
}

func (e *NpcPreview) writeNoAlign(field string, value string) {
	var newPos fixed.Point26_6
	if len(field) == 0 { // non-bold value only
		newPos, _ = e.c.DrawString(value, e.pt)
		if newPos.X.Ceil() > e.maxWidth {
			e.maxWidth = newPos.X.Ceil()
		}
		return
	}
	newPos, _ = e.cb.DrawString(field, e.pt) // field key
	if newPos.X.Ceil() > e.maxWidth {
		e.maxWidth = newPos.X.Ceil()
	}
	if len(value) > 0 { // if field/value pairing non-bold value
		newPos, _ = e.c.DrawString(" "+value, newPos)
		if newPos.X.Ceil() > e.maxWidth {
			e.maxWidth = newPos.X.Ceil()
		}
	}
}

func (e *NpcPreview) writeNoAlignLn(field string, value string) {
	e.writeNoAlign(field, value)
	e.newLine(1)
}

func GenerateNpcPreview(npcIcon int32, lines []string) ([]byte, error) {
	mu.RLock()
	defer mu.RUnlock()

	var newPos fixed.Point26_6

	font, err := truetype.Parse(npcFont)
	if err != nil {
		return nil, fmt.Errorf("parse font: %w", err)
	}

	fontSize := float64(14)

	fontBold, err := truetype.Parse(npcFontBold)
	if err != nil {
		return nil, fmt.Errorf("parse fontBold: %w", err)
	}

	rgba := image.NewRGBA(image.Rect(0, 0, 700, 600))
	draw.Draw(rgba, rgba.Bounds(), npcBGImage, image.Pt(0, 0), draw.Src)
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(font)
	c.SetFontSize(fontSize)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(npcFGImage)

	cb := freetype.NewContext()
	cb.SetDPI(72)
	cb.SetFont(fontBold)
	cb.SetFontSize(fontSize)
	cb.SetClip(rgba.Bounds())
	cb.SetDst(rgba)
	cb.SetSrc(npcFGImage)

	cTitle := freetype.NewContext()
	cTitle.SetDPI(72)
	cTitle.SetFont(fontBold)
	cTitle.SetFontSize(fontSize * 2)
	cTitle.SetClip(rgba.Bounds())
	cTitle.SetDst(rgba)
	cTitle.SetSrc(npcFGImage)

	e := &NpcPreview{
		c:        c,
		cb:       cb,
		cbFont:   fontBold,
		cTitle:   cTitle,
		fontSize: fontSize,
	}

	//c.SetHinting(font.H)

	// title

	isSlot := false
	for i, line := range lines {
		if i == 0 {
			line = strings.ReplaceAll(line, "Npc Info for Effect: ", "")
			iconImg := store.NpcIcon(npcIcon)
			titleOffset := 10
			// draw a 40x41 image
			if iconImg != nil {
				draw.Draw(rgba, image.Rect(10, 10, 700, 600), iconImg, image.Pt(0, 0), draw.Src)
				titleOffset = 60
			}
			e.setCursor(titleOffset, 40)
			newPos, _ = e.cTitle.DrawString(line, e.pt)
			if newPos.X.Ceil() > e.maxWidth {
				e.maxWidth = newPos.X.Ceil()
			}

			e.setCursor(10, 80)
			continue
		}
		if strings.HasPrefix(line, "Slot") {
			isSlot = true
		}
		if isSlot && !strings.HasPrefix(line, "Slot") {
			e.newLine(1)
		}
		e.writeNoAlignLn("", line)
	}

	//e.maxHeight -= int(e.fontSize * 1)
	e.maxWidth += 10

	if e.maxWidth != 700 || e.maxHeight != 600 {
		rgba2 := image.NewRGBA(image.Rect(0, 0, e.maxWidth, e.maxHeight))
		draw.Draw(rgba2, rgba2.Bounds(), npcBGImage, image.Pt(0, 0), draw.Src)

		draw.Draw(rgba2, rgba2.Bounds(), rgba, image.Pt(0, 0), draw.Src)
		rgba = rgba2
	}

	// ratio := 0.60
	// rgba2 := resize.Resize(uint(ratio*700), uint(ratio*float64(e.maxHeight)), rgba, resize.Lanczos3)

	b := new(bytes.Buffer)

	err = png.Encode(b, rgba) //&jpeg.Options{Quality: 100});
	if err != nil {
		log.Println("unable to encode image.")
		return nil, err
	}
	return b.Bytes(), nil
}

func (e *NpcPreview) newLine(lineCount int) {
	e.lineCurrent += lineCount
	e.shiftLn(lineCount)
}

func (e *NpcPreview) shiftLn(lineCount int) {
	e.pt.Y += e.c.PointToFixed(float64(lineCount) * e.fontSize * 1.5)
	if e.pt.Y.Ceil() > e.maxHeight {
		e.maxHeight = e.pt.Y.Ceil()
	}

}

func (e *NpcPreview) setCursor(x int, y int) {
	e.pt = freetype.Pt(x, y)
}
