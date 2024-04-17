package util

import "fmt"

func (e *Util) SpellName(id int) string {
	name, ok := spellNames[id]
	if !ok {
		return fmt.Sprintf("Unknown Spell (%d)", id)
	}
	return name
}
