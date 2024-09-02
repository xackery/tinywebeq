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
	SlotPowerSource
	SlotAmmo
)

var slotToString = map[Slot]string{
	SlotCharm:       "Charm",
	SlotEarL:        "Ear",
	SlotHead:        "Head",
	SlotFace:        "Face",
	SlotEarR:        "Ear",
	SlotNeck:        "Neck",
	SlotShoulders:   "Shoulders",
	SlotArms:        "Arms",
	SlotBack:        "Back",
	SlotWristL:      "Wrist",
	SlotWristR:      "Wrist",
	SlotRange:       "Range",
	SlotHands:       "Hands",
	SlotSecondary:   "Secondary",
	SlotPrimary:     "Primary",
	SlotFingerL:     "Finger",
	SlotFingerR:     "Finger",
	SlotChest:       "Chest",
	SlotLegs:        "Legs",
	SlotFeet:        "Feet",
	SlotWaist:       "Waist",
	SlotPowerSource: "Power Source",
	SlotAmmo:        "Ammo",
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
