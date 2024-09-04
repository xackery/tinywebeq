package types

import (
	"strings"
)

type Race int32

type RacesBitmask int32

const (
	RaceHuman Race = 1 << iota
	RaceBarbarian
	RaceErudite
	RaceWoodElf
	RaceHighElf
	RaceDarkElf
	RaceHalfElf
	RaceDwarf
	RaceTroll
	RaceOgre
	RaceHalfling
	RaceGnome
	RaceIksar
	RaceVahShir
	RaceFroglok
	RaceDrakkin
)

var raceToShortString = map[Race]string{
	RaceHuman:     "HUM",
	RaceBarbarian: "BAR",
	RaceErudite:   "ERU",
	RaceWoodElf:   "ELF",
	RaceHighElf:   "HIE",
	RaceDarkElf:   "DEF",
	RaceHalfElf:   "HEF",
	RaceDwarf:     "DWF",
	RaceTroll:     "TRL",
	RaceOgre:      "OGR",
	RaceHalfling:  "HFL",
	RaceGnome:     "GNM",
	RaceIksar:     "IKS",
	RaceVahShir:   "VAH",
	RaceFroglok:   "FRG",
	RaceDrakkin:   "DRK",
}

func (r RacesBitmask) String() string {
	races := make([]string, 0)

	var i int32
	for i = 1; i <= int32(r); i <<= 1 {
		if i&int32(r) != 0 {
			races = append(races, raceToShortString[Race(i)])
		}
	}

	if len(races) == 16 {
		return "ALL"
	}

	return strings.Join(races, " ")
}

func (r RacesBitmask) MarshalJSON() ([]byte, error) {
	return []byte(`"` + r.String() + `"`), nil
}
