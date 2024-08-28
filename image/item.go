package image

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/xackery/tinywebeq/model"
	"github.com/xackery/tinywebeq/store"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/math/fixed"
)

var (
	itemFGColor  = color.RGBA{0xdb, 0xde, 0xe1, 0xff}
	itemBGColor  = color.RGBA{0x31, 0x33, 0x38, 0xff}
	itemFGImage  = image.NewUniform(itemFGColor)
	itemBGImage  = image.NewUniform(itemBGColor)
	itemFont     = goregular.TTF
	itemFontBold = gobold.TTF
)

type ItemPreview struct {
	c           *freetype.Context
	cb          *freetype.Context
	cbFont      *truetype.Font
	cTitle      *freetype.Context
	item        *model.Item
	itemQuest   *model.ItemQuest
	itemRecipe  *model.ItemRecipe
	fontSize    float64
	pt          fixed.Point26_6
	lineStart   int
	lineCurrent int
	lineMax     int
	maxHeight   int
	maxWidth    int
}

func (e *ItemPreview) write(field string, value string) {
	var newPos fixed.Point26_6

	newPos, _ = e.cb.DrawString(field, e.pt)
	if newPos.X.Ceil() > e.maxWidth {
		e.maxWidth = newPos.X.Ceil()
	}

	if len(value) > 0 {
		baseX := e.pt.X
		textWidth := 0
		face := truetype.NewFace(e.cbFont, &truetype.Options{Size: e.fontSize, DPI: 72})
		for _, x := range value {
			awidth, ok := face.GlyphAdvance(rune(x))
			if !ok {
				continue
			}
			iwidthf := int(float64(awidth) / 64)
			textWidth += iwidthf
		}

		rightAlignX := e.pt.X + e.c.PointToFixed(150) - e.c.PointToFixed(float64(textWidth))

		// Update position to right-aligned position
		e.pt.X = rightAlignX
		newPos, _ = e.c.DrawString(value, e.pt)
		if newPos.X.Ceil() > e.maxWidth {
			e.maxWidth = newPos.X.Ceil()
		}

		e.pt.X = baseX
	}
}

func (e *ItemPreview) writeNoAlign(field string, value string) {
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

func (e *ItemPreview) writeNoAlignLn(field string, value string) {
	e.writeNoAlign(field, value)
	e.newLine(1)
}

func GenerateItemPreview(item *model.Item, itemQuest *model.ItemQuest, itemRecipe *model.ItemRecipe) ([]byte, error) {
	mu.RLock()
	defer mu.RUnlock()
	var newPos fixed.Point26_6

	font, err := truetype.Parse(itemFont)
	if err != nil {
		return nil, fmt.Errorf("parse font: %w", err)
	}

	fontSize := float64(16)

	fontBold, err := truetype.Parse(itemFontBold)
	if err != nil {
		return nil, fmt.Errorf("parse fontBold: %w", err)
	}

	rgba := image.NewRGBA(image.Rect(0, 0, 700, 600))
	draw.Draw(rgba, rgba.Bounds(), itemBGImage, image.Pt(0, 0), draw.Src)
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(font)
	c.SetFontSize(fontSize)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(itemFGImage)

	cb := freetype.NewContext()
	cb.SetDPI(72)
	cb.SetFont(fontBold)
	cb.SetFontSize(fontSize)
	cb.SetClip(rgba.Bounds())
	cb.SetDst(rgba)
	cb.SetSrc(itemFGImage)

	cTitle := freetype.NewContext()
	cTitle.SetDPI(72)
	cTitle.SetFont(fontBold)
	cTitle.SetFontSize(fontSize * 2)
	cTitle.SetClip(rgba.Bounds())
	cTitle.SetDst(rgba)
	cTitle.SetSrc(itemFGImage)

	e := &ItemPreview{
		c:          c,
		cb:         cb,
		cbFont:     fontBold,
		cTitle:     cTitle,
		item:       item,
		itemQuest:  itemQuest,
		itemRecipe: itemRecipe,
		fontSize:   fontSize,
	}

	//c.SetHinting(font.H)

	// title
	titleOffset := 10

	iconImg := store.ItemIcon(item.Icon)
	if iconImg != nil {
		draw.Draw(rgba, image.Rect(10, 10, 700, 600), iconImg, image.Pt(0, 0), draw.Over)
		titleOffset = 60
	}

	e.setCursor(titleOffset, 40)
	newPos, _ = e.cTitle.DrawString(item.Name, e.pt)
	if newPos.X.Ceil() > e.maxWidth {
		e.maxWidth = newPos.X.Ceil()
	}

	e.lineStart = 2
	e.render1Left()
	e.lineStart = e.lineMax + 1
	e.render2Left()
	e.render2Center()
	e.render2Right()
	e.lineStart = e.lineMax + 1
	e.render3Left()
	e.render3Center()
	e.render3Right()
	e.lineStart = e.lineMax + 1
	e.render4Left()
	e.lineStart = e.lineMax + 1
	e.render5Left()

	//resize rgba image to maxHeight

	//rgba2 := image.NewRGBA(image.Rect(0, 0, 700, maxHeight))
	//draw.Draw(rgba2, rgba2.Bounds(), bg, image.Pt(0, 0), draw.Src)
	//draw.Draw(rgba2, rgba2.Bounds(), rgba, image.Pt(0, 0), draw.Src)
	//ratio := float64(maxHeight) / 600 / 1.5
	//rgba2 := resize.Resize(uint(ratio*700), uint(ratio*600), rgba, resize.Lanczos3)
	//resize image to half

	//e.maxHeight -= int(e.fontSize * 1)
	e.maxWidth += 10

	if e.maxHeight != 600 || e.maxWidth != 700 {
		rgba2 := image.NewRGBA(image.Rect(0, 0, e.maxWidth, e.maxHeight))
		draw.Draw(rgba2, rgba2.Bounds(), itemBGImage, image.Pt(0, 0), draw.Src)
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

func (e *ItemPreview) newLine(lineCount int) {
	e.lineCurrent += lineCount
	e.shiftLn(lineCount)
}

func (e *ItemPreview) shiftLn(lineCount int) {
	e.pt.Y += e.c.PointToFixed(float64(lineCount) * e.fontSize * 1.5)
	if e.pt.Y.Ceil() > e.maxHeight {
		e.maxHeight = e.pt.Y.Ceil()
	}

}

func (e *ItemPreview) writeLn(field string, value string) {
	e.write(field, value)
	e.newLine(1)
}

func (e *ItemPreview) setCursor(x int, y int) {
	e.pt = freetype.Pt(x, y)
}

func (e *ItemPreview) render1Left() {
	item := e.item
	e.setCursor(10, 20)
	e.shiftLn(e.lineStart)
	e.lineCurrent = e.lineStart

	e.writeNoAlignLn("", item.TagStr())
	if item.Classes > 0 {
		e.writeNoAlignLn("Class:", item.ClassesStr())
	}
	if item.Races > 0 {
		e.writeNoAlignLn("Race:", item.RaceStr())
	}
	if item.Deity > 0 {
		e.writeNoAlignLn("Deity:", item.DeityStr())
	}
	if item.Slots > 0 {
		e.writeNoAlignLn(item.SlotStr(), "")
	} else {
		if false {
			e.writeNoAlignLn("NONE", "")
		}
	}

	if item.Bagslots > 0 {
		e.writeNoAlignLn("Item Type:", "Container")
		e.writeNoAlignLn("Number of Slots:", fmt.Sprintf("%d", item.Bagslots))
		if item.Bagtype > 0 {
			e.writeNoAlignLn("Trade Skill Container:", item.BagTypeStr())
		}
		if item.Bagwr > 0 {
			e.writeNoAlignLn("Weight Reduction:", fmt.Sprintf("%d%%", item.Bagwr))
		}
		e.writeNoAlignLn("", fmt.Sprintf("This can hold %s and smaller items", item.BagsizeStr()))
	}
	if e.lineCurrent > e.lineMax {
		e.lineMax = e.lineCurrent
	}
}

func (e *ItemPreview) render2Left() {
	item := e.item
	e.setCursor(10, 20)
	e.shiftLn(e.lineStart)
	e.lineCurrent = e.lineStart

	e.writeLn("Size:", item.SizeStr())
	if item.Weight > 0 {
		e.writeLn("Weight:", fmt.Sprintf("%0.1f", float64(item.Weight/10)))
	}
	if item.Slots > 0 {
		e.writeLn(item.TypeStr(), "Inventory")
	} else {
		e.writeLn(item.TypeStr(), item.ItemTypeStr())
	}
	if item.Reclevel > 0 {
		e.writeLn("Rec Level:", fmt.Sprintf("%d", item.Reclevel))
	}
	if item.Reqlevel > 0 {
		e.writeLn("Req Level:", fmt.Sprintf("%d", item.Reqlevel))
	}

	if e.lineCurrent > e.lineMax {
		e.lineMax = e.lineCurrent
	}
}

func (e *ItemPreview) render3Left() {
	item := e.item
	e.setCursor(10, 20)
	e.shiftLn(e.lineStart)
	e.lineCurrent = e.lineStart
	if item.Astr != 0 {
		e.writeLn("Strength:", fmt.Sprintf("%d", item.Astr))
	}
	if item.Asta != 0 {
		e.writeLn("Stamina:", fmt.Sprintf("%d", item.Asta))
	}
	if item.Aint != 0 {
		e.writeLn("Intelligence:", fmt.Sprintf("%d", item.Aint))
	}
	if item.Awis != 0 {
		e.writeLn("Wisdom:", fmt.Sprintf("%d", item.Awis))
	}
	if item.Aagi != 0 {
		e.writeLn("Agility:", fmt.Sprintf("%d", item.Aagi))
	}
	if item.Adex != 0 {
		e.writeLn("Dexterity:", fmt.Sprintf("%d", item.Adex))
	}
	if item.Acha != 0 {
		e.writeLn("Charisma:", fmt.Sprintf("%d", item.Acha))
	}

	if e.lineCurrent > e.lineMax {
		e.lineMax = e.lineCurrent
	}
}

func (e *ItemPreview) render4Left() {
	itemQuest := e.itemQuest
	itemRecipe := e.itemRecipe
	e.setCursor(10, 20)
	e.shiftLn(e.lineStart)
	e.lineCurrent = e.lineStart
	counter := 0
	if itemQuest != nil {
		for _, entry := range itemQuest.Entries {
			counter++
			if counter > 3 {
				break
			}
			e.writeNoAlignLn("Quest", fmt.Sprintf("%s in %s", entry.NpcCleanName(), store.ZoneLongNameByZoneIDNumber(int32(entry.ZoneID))))
		}
	}
	if itemRecipe != nil {
		for _, entry := range itemRecipe.Entries {

			if entry.ComponentCount > 0 {
				counter++
				if counter > 3 {
					break
				}
				e.writeNoAlignLn("Recipe Component", fmt.Sprintf("%s x%d", entry.RecipeName, entry.ComponentCount))
			}
			if entry.SuccessCount > 0 {
				counter++
				if counter > 3 {
					break
				}
				e.writeNoAlignLn("Recipe Result", fmt.Sprintf("%s x%d", entry.RecipeName, entry.SuccessCount))
			}
		}
	}

	if e.lineCurrent > e.lineMax {
		e.lineMax = e.lineCurrent
	}

}

func (e *ItemPreview) render5Left() {
	item := e.item
	e.setCursor(10, 20)
	e.shiftLn(e.lineStart)
	e.lineCurrent = e.lineStart
	if item.Extradmgamt != 0 {
		e.writeLn(fmt.Sprintf("%s Damage:", item.ExtraDamageSkillStr()), fmt.Sprintf("%d", item.Extradmgamt))
	}
	if item.Skillmodtype != 0 && item.Skillmodvalue != 0 {
		e.writeLn(fmt.Sprintf("Skill Mod %s:", item.SkillModTypeStr()), fmt.Sprintf("%d", item.Skillmodvalue))
	}
	if item.Augslot1type > 0 {
		e.writeNoAlignLn("Slot 1:", fmt.Sprintf("Type %d", item.Augslot1type))
	}
	if item.Augslot2type > 0 {
		e.writeNoAlignLn("Slot 2:", fmt.Sprintf("Type %d", item.Augslot2type))
	}
	if item.Augslot3type > 0 {
		e.writeNoAlignLn("Slot 3:", fmt.Sprintf("Type %d", item.Augslot3type))
	}
	if item.Augslot4type > 0 {
		e.writeNoAlignLn("Slot 4:", fmt.Sprintf("Type %d", item.Augslot4type))
	}
	if item.Augslot5type > 0 {
		e.writeNoAlignLn("Slot 5:", fmt.Sprintf("Type %d", item.Augslot5type))
	}

	e.newLine(1)

	if item.Proceffect > 0 && item.Proceffect != 65535 {
		e.writeNoAlignLn("Combat Effects:", fmt.Sprintf("%s (%d)", store.SpellName(item.Proceffect), item.Proceffect))
		if item.Proclevel2 > 0 {
			e.writeNoAlignLn("Level for effect:", fmt.Sprintf("%d", item.Proclevel2))
		}
		e.writeNoAlignLn("Effect chance modifier:", fmt.Sprintf("%d%%", item.Procrate+100))
		e.renderSpellInfo(item.Proceffect)
	}

	if item.Worneffect > 0 && item.Worneffect != 65535 {
		e.writeNoAlignLn("Worn Effect:", fmt.Sprintf("%s (%d)", store.SpellName(item.Worneffect), item.Worneffect))
		e.renderSpellInfo(item.Worneffect)
	}
	if item.Wornlevel > 0 {
		e.writeNoAlignLn("Level for effect:", fmt.Sprintf("%d", item.Wornlevel))
		e.renderSpellInfo(item.Worneffect)
	}
	if item.Focuseffect > 0 && item.Focuseffect != 65535 {
		e.writeNoAlignLn("Focus Effect:", fmt.Sprintf("%s (%d)", store.SpellName(item.Focuseffect), item.Focuseffect))
		e.renderSpellInfo(item.Focuseffect)
	}
	if item.Focuslevel > 0 {
		e.writeNoAlignLn("Level for effect:", fmt.Sprintf("%d", item.Focuslevel))
	}

	if item.Clickeffect > 0 && item.Clickeffect != 65535 {
		details := store.SpellName(item.Clickeffect) + fmt.Sprintf(" (%d)", item.Clickeffect) + " ("
		if item.Clicktype == 4 {
			details += "Must Equip. "
		}
		if item.Casttime > 0 {
			details += fmt.Sprintf("Casting Time: %0.1f sec", float64(item.Casttime/1000))
		} else {
			details += "Casting Time: Instant"
		}
		details += ")"
		e.writeNoAlignLn("Click Effect:", details)
		if item.Clicklevel > 0 {
			e.writeNoAlignLn("Level for effect:", fmt.Sprintf("%d", item.Clicklevel))
		}
		if item.Maxcharges > 0 {
			e.writeNoAlignLn("Charges:", fmt.Sprintf("%d", item.Maxcharges))
		} else if item.Maxcharges < 0 {
			e.writeNoAlignLn("Charges:", "Unlimited")
		} else {
			e.writeNoAlignLn("Charges:", "None")
		}
		e.renderSpellInfo(item.Clickeffect)
	}

	if item.Scrolleffect > 0 && item.Scrolleffect != 65535 {
		e.writeNoAlignLn("Spell Scroll Effect:", fmt.Sprintf("%s (%d)", store.SpellName(item.Scrolleffect), item.Scrolleffect))
		e.renderSpellInfo(item.Scrolleffect)
	}

	if item.Bardtype > 22 && item.Bardtype < 65535 {
		e.writeNoAlignLn("Bard Skill:", fmt.Sprintf("%s (%d%%)", item.BardTypeStr(), item.Bardvalue*10-100))
	}

	details := item.AugSlotStr()
	if len(details) > 0 {
		e.writeNoAlignLn("", details)
	}
	details = item.AugRestrictStr()
	if len(details) > 0 {
		e.writeNoAlignLn("", details)
	}

	pp := 0
	gp := 0
	sp := 0
	cp := 0

	out := ""

	if int(item.Price) > 1000 {
		pp = int(int(item.Price) / 1000)
		if pp > 0 {
			out += fmt.Sprintf("%dp ", pp)
		}
	}
	if int(item.Price)-(pp*1000) > 100 {
		gp = int((int(item.Price) - (pp * 1000)) / 100)
		if gp > 0 {
			out += fmt.Sprintf("%dg ", gp)
		}
	}

	if int(item.Price)-(pp*1000)-(gp*100) > 10 {
		sp = int((int(item.Price) - (pp * 1000) - (gp * 100)) / 10)
		if sp > 0 {
			out += fmt.Sprintf("%ds ", sp)
		}
	}

	if int(item.Price)-(pp*1000)-(gp*100)-(sp*10) > 0 {
		cp = int(item.Price) - (pp * 1000) - (gp * 100) - (sp * 10)
		if cp > 0 {
			out += fmt.Sprintf("%dc ", cp)
		}
	}
	if len(out) > 0 {
		out = out[:len(out)-1]
	}

	if int(item.Price) > 0 {
		e.writeNoAlignLn("Value:", out)
	}

	if e.lineCurrent > e.lineMax {
		e.lineMax = e.lineCurrent
	}
}

func (e *ItemPreview) render2Center() {
	item := e.item
	e.setCursor(200, 20)
	e.shiftLn(e.lineStart)
	e.lineCurrent = e.lineStart
	if item.Ac != 0 {
		e.writeLn("AC:", fmt.Sprintf("%d", item.Ac))
	}
	if item.Hp != 0 {
		e.writeLn("HP:", fmt.Sprintf("%d", item.Hp))
	}
	if item.Mana != 0 {
		e.writeLn("Mana:", fmt.Sprintf("%d", item.Mana))
	}
	if item.Endur != 0 {
		e.writeLn("Endurance:", fmt.Sprintf("%d", item.Endur))
	}
	if item.Haste != 0 {
		e.writeLn("Haste:", fmt.Sprintf("%d%%", item.Haste))
	}

	if e.lineCurrent > e.lineMax {
		e.lineMax = e.lineCurrent
	}
}

func (e *ItemPreview) render3Center() {
	item := e.item
	e.setCursor(200, 20)
	e.shiftLn(e.lineStart)
	e.lineCurrent = e.lineStart

	if item.Mr != 0 {
		e.writeLn("Magic Resist:", fmt.Sprintf("%d", item.Mr))
	}
	if item.Fr != 0 {
		e.writeLn("Fire Resist:", fmt.Sprintf("%d", item.Fr))
	}
	if item.Cr != 0 {
		e.writeLn("Cold Resist:", fmt.Sprintf("%d", item.Cr))
	}
	if item.Dr != 0 {
		e.writeLn("Disease Resist:", fmt.Sprintf("%d", item.Dr))
	}
	if item.Pr != 0 {
		e.writeLn("Poison Resist:", fmt.Sprintf("%d", item.Pr))
	}

	if e.lineCurrent > e.lineMax {
		e.lineMax = e.lineCurrent
	}

}

func (e *ItemPreview) render2Right() {
	item := e.item
	e.setCursor(400, 20)
	e.shiftLn(e.lineStart)
	e.lineCurrent = e.lineStart

	if item.Damage > 0 {
		e.writeLn("Base Damage:", fmt.Sprintf("%d", item.Damage))
	}
	if item.Delay > 0 {
		e.writeLn("Delay:", "24")
	}
	if item.DamageBonus() > 0 {
		e.writeLn("Damage Bonus:", fmt.Sprintf("%d", item.DamageBonus()))
	}
	if item.Range > 0 {
		e.writeLn("Range:", fmt.Sprintf("%d", item.Range))
	}

	if e.lineCurrent > e.lineMax {
		e.lineMax = e.lineCurrent
	}
}

func (e *ItemPreview) render3Right() {
	item := e.item
	e.setCursor(400, 20)
	e.shiftLn(e.lineStart)
	e.lineCurrent = e.lineStart

	if item.Attack != 0 {
		e.writeLn("Attack:", fmt.Sprintf("%d", item.Attack))
	}
	if item.Regen > 0 {
		e.writeLn("HP Regen:", fmt.Sprintf("%d", item.Regen))
	}
	if item.Manaregen != 0 {
		e.writeLn("Mana Regen:", fmt.Sprintf("%d", item.Manaregen))
	}
	if item.Enduranceregen != 0 {
		e.writeLn("Endurance Regen:", fmt.Sprintf("%d", item.Enduranceregen))
	}
	if item.Spellshield != 0 {
		e.writeLn("Spell Shield:", fmt.Sprintf("%d", item.Spellshield))
	}
	if item.Dotshielding != 0 {
		e.writeLn("Dot Shielding:", fmt.Sprintf("%d", item.Dotshielding))
	}
	if item.Avoidance != 0 {
		e.writeLn("Avoidance:", fmt.Sprintf("%d", item.Avoidance))
	}
	if item.Accuracy != 0 {
		e.writeLn("Accuracy:", fmt.Sprintf("%d", item.Accuracy))
	}
	if item.Stunresist != 0 {
		e.writeLn("Stun Resist:", fmt.Sprintf("%d", item.Stunresist))
	}
	if item.Strikethrough != 0 {
		e.writeLn("Strikethrough:", fmt.Sprintf("%d", item.Strikethrough))
	}
	if item.Damageshield != 0 {
		e.writeLn("Damage Shield:", fmt.Sprintf("%d", item.Damageshield))
	}

	if e.lineCurrent > e.lineMax {
		e.lineMax = e.lineCurrent
	}
}

func (e *ItemPreview) renderSpellInfo(id int32) {
	_, info := store.SpellInfo(id, 0)
	for _, line := range info {
		e.writeNoAlignLn("", line)
	}
}
