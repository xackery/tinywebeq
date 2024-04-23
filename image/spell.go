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
	"github.com/nfnt/resize"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/math/fixed"
)

var (
	spellBGColor  = color.RGBA{0x30, 0x30, 0x30, 0xff}
	spellFGColor  = color.RGBA{0xff, 0xff, 0xff, 0xff}
	spellFGImage  = image.NewUniform(spellFGColor)
	spellBGImage  = image.NewUniform(spellBGColor)
	spellFont     = goregular.TTF
	spellFontBold = gobold.TTF
)

type SpellPreview struct {
	c           *freetype.Context
	cb          *freetype.Context
	cbFont      *truetype.Font
	cTitle      *freetype.Context
	fontSize    float64
	pt          fixed.Point26_6
	lineCurrent int
	maxHeight   int
}

func (e *SpellPreview) writeNoAlign(field string, value string) {
	if len(field) == 0 { // non-bold value only
		e.c.DrawString(value, e.pt)
		return
	}
	newPos, _ := e.cb.DrawString(field, e.pt) // field key
	if len(value) > 0 {                       // if field/value pairing non-bold value
		e.c.DrawString(" "+value, newPos)
	}
}

func (e *SpellPreview) writeNoAlignLn(field string, value string) {
	e.writeNoAlign(field, value)
	e.newLine(1)
}

func GenerateSpellPreview(lines []string) ([]byte, error) {
	mu.RLock()
	defer mu.RUnlock()

	font, err := truetype.Parse(spellFont)
	if err != nil {
		return nil, fmt.Errorf("parse font: %w", err)
	}

	fontSize := float64(16)

	fontBold, err := truetype.Parse(spellFontBold)
	if err != nil {
		return nil, fmt.Errorf("parse fontBold: %w", err)
	}

	rgba := image.NewRGBA(image.Rect(0, 0, 700, 600))
	draw.Draw(rgba, rgba.Bounds(), spellBGImage, image.Pt(0, 0), draw.Src)
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(font)
	c.SetFontSize(fontSize)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(spellFGImage)

	cb := freetype.NewContext()
	cb.SetDPI(72)
	cb.SetFont(fontBold)
	cb.SetFontSize(fontSize)
	cb.SetClip(rgba.Bounds())
	cb.SetDst(rgba)
	cb.SetSrc(spellFGImage)

	cTitle := freetype.NewContext()
	cTitle.SetDPI(72)
	cTitle.SetFont(fontBold)
	cTitle.SetFontSize(fontSize * 2)
	cTitle.SetClip(rgba.Bounds())
	cTitle.SetDst(rgba)
	cTitle.SetSrc(spellFGImage)

	e := &SpellPreview{
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
			line = strings.ReplaceAll(line, "Spell Info for Effect: ", "")
			e.setCursor(10, 30)
			e.cTitle.DrawString(line, e.pt)
			e.setCursor(10, 60)
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

	e.maxHeight -= int(e.fontSize * 1)

	if e.maxHeight != 600 {
		rgba2 := image.NewRGBA(image.Rect(0, 0, 700, e.maxHeight))
		draw.Draw(rgba2, rgba2.Bounds(), spellBGImage, image.Pt(0, 0), draw.Src)
		draw.Draw(rgba2, rgba2.Bounds(), rgba, image.Pt(0, 0), draw.Src)
		rgba = rgba2
	}

	ratio := 0.60
	rgba2 := resize.Resize(uint(ratio*700), uint(ratio*float64(e.maxHeight)), rgba, resize.Bicubic)
	//resize image to half

	b := new(bytes.Buffer)

	err = png.Encode(b, rgba2) //&jpeg.Options{Quality: 100});
	if err != nil {
		log.Println("unable to encode image.")
		return nil, err
	}
	return b.Bytes(), nil
}

func (e *SpellPreview) newLine(lineCount int) {
	e.lineCurrent += lineCount
	e.shiftLn(lineCount)
}

func (e *SpellPreview) shiftLn(lineCount int) {
	e.pt.Y += e.c.PointToFixed(float64(lineCount) * e.fontSize * 1.5)
	if e.pt.Y.Ceil() > e.maxHeight {
		e.maxHeight = e.pt.Y.Ceil()
	}

}

func (e *SpellPreview) setCursor(x int, y int) {
	e.pt = freetype.Pt(x, y)
}
