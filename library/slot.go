package library

import (
	"bytes"
	"strings"
)

type (
	//go:generate stringer -type=Slot
	Slot int

	Slots []Slot
)

const (
	SlotCharm Slot = 1 << iota
	SlotEarL
	SlotHead
	SlotFace
	SlotEarR
	SlotNeck
	SlotShoulders
	SlotArms
	SlotBack
	SlotWristL
	SlotWristR
	SlotRange
	SlotHands
	SlotSecondary
	SlotPrimary
	SlotFingerL
	SlotFingerR
	SlotChest
	SlotLegs
	SlotFeet
	SlotWaist
	SlotAmmo
	SlotPowerSource
)

var slotToString = map[Slot]string{
	SlotCharm:       "CHARM",
	SlotEarL:        "EAR",
	SlotHead:        "HEAD",
	SlotFace:        "FACE",
	SlotEarR:        "EAR",
	SlotNeck:        "NECK",
	SlotShoulders:   "SHOULDERS",
	SlotArms:        "ARMS",
	SlotBack:        "BACK",
	SlotWristL:      "WRIST",
	SlotWristR:      "WRIST",
	SlotRange:       "RANGE",
	SlotHands:       "HANDS",
	SlotSecondary:   "SECONDARY",
	SlotPrimary:     "PRIMARY",
	SlotFingerL:     "FINGER",
	SlotFingerR:     "FINGER",
	SlotChest:       "CHEST",
	SlotLegs:        "LEGS",
	SlotFeet:        "FEET",
	SlotWaist:       "WAIST",
	SlotAmmo:        "AMMO",
	SlotPowerSource: "POWER SOURCE",
}

func (s Slot) String() string {
	return slotToString[s]
}

func (s Slot) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBufferString(`"`)
	buf.WriteString(s.String())
	buf.WriteString(`"`)
	return buf.Bytes(), nil
}

func (s Slots) String() string {
	type unique map[string]struct{}

	// There are duplicate entries for things like ears, wrists that don't matter for item display.
	// Create a unique map of slots by name, then concatenate them.
	tmp := make(unique)
	for _, slot := range s {
		tmp[slotToString[slot]] = struct{}{}
	}

	var sl []string
	for slot := range tmp {
		sl = append(sl, slot)
	}

	str := &strings.Builder{}
	for i, slot := range sl {
		str.WriteString(slot)
		if i != 0 {
			str.WriteString(" ")
		}
	}

	return str.String()
}

func (s Slots) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}

func SlotsFromBitmask(mask int32) Slots {
	var slots Slots
	var i int32
	for i = 1; i <= mask; i <<= 1 {
		if i&mask != 0 {
			slots = append(slots, Slot(i))
		}
	}
	return slots
}
