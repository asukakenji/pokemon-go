package main

import (
	"math"
	"sort"

	"github.com/asukakenji/pokemon-go/lv"
	"github.com/asukakenji/pokemon-go/pokemon"
)

type Spec struct {
	iv  pokemon.IndividualValues
	lvl lv.Level
}

type PokemonFilter struct {
	pkm   pokemon.Pokemon
	specs []Spec
}

func NewPokemonFilter() *PokemonFilter {
	filter := &PokemonFilter{pokemon.None, nil}
	filter.Reset()
	return filter
}

func (filter *PokemonFilter) Reset() {
	ivs := make([]pokemon.IndividualValues, 0, 16*16*16)
	for stamina := 0; stamina <= 15; stamina++ {
		for attack := 0; attack <= 15; attack++ {
			for defense := 0; defense <= 15; defense++ {
				ivs = append(ivs, pokemon.IndividualValues{stamina, attack, defense})
			}
		}
	}

	lvls := make([]lv.Level, 0, (40-1)*2)
	for lvl := lv.Level(float32(1.0)); lvl <= 40.0; lvl += 0.5 {
		lvls = append(lvls, lvl)
	}

	specs := make([]Spec, 0, len(ivs)*len(lvls))
	for _, iv := range ivs {
		for _, lvl := range lvls {
			specs = append(specs, Spec{iv, lvl})
		}
	}
	filter.specs = specs
}

func (filter *PokemonFilter) IncrementLvl() {
	for i := range filter.specs {
		filter.specs[i].lvl += 0.5
	}
}

func (filter *PokemonFilter) SetPokemon(pkm pokemon.Pokemon) {
	filter.pkm = pkm
}

func (filter *PokemonFilter) SetCp(cp int) {
	specs := []Spec{}
	for _, spec := range filter.specs {
		cpMin, cpMax, _, _ := filter.pkm.CombatPowerAndHitPoints(spec.iv, spec.lvl)
		if match(cp, cpMin, cpMax) {
			specs = append(specs, spec)
		}
	}
	filter.specs = specs
}

func (filter *PokemonFilter) SetHp(hp int) {
	specs := []Spec{}
	for _, spec := range filter.specs {
		_, _, hpMin, hpMax := filter.pkm.CombatPowerAndHitPoints(spec.iv, spec.lvl)
		if match(hp, hpMin, hpMax) {
			specs = append(specs, spec)
		}
	}
	filter.specs = specs
}

func (filter *PokemonFilter) SetSd(sd int) {
	specs := []Spec{}
	for _, spec := range filter.specs {
		if spec.lvl.StardustToPowerUp() == sd {
			specs = append(specs, spec)
		}
	}
	filter.specs = specs
}

func (filter *PokemonFilter) SetCd(cd int) {
	specs := []Spec{}
	for _, spec := range filter.specs {
		if spec.lvl.CandyToPowerUp() == cd {
			specs = append(specs, spec)
		}
	}
	filter.specs = specs
}

func (filter *PokemonFilter) SetIsWild(isWild bool) {
	if !isWild {
		return
	}
	specs := []Spec{}
	for _, spec := range filter.specs {
		if spec.lvl.IsInteger() {
			specs = append(specs, spec)
		}
	}
	filter.specs = specs
}

func (filter *PokemonFilter) Pokemon() pokemon.Pokemon {
	return filter.pkm
}

func (filter *PokemonFilter) MinMaxCpHp() (minCp, maxCp, minHp, maxHp int) {
	if len(filter.specs) == 0 {
		return 0, 0, 0, 0
	}

	minCp, maxCp, minHp, maxHp = 2147483647, 0, 2147483647, 0
	for _, spec := range filter.specs {
		cpMin, cpMax, hpMin, hpMax := filter.pkm.CombatPowerAndHitPoints(spec.iv, spec.lvl)
		if cpMin < minCp {
			minCp = cpMin
		}
		if cpMax > maxCp {
			maxCp = cpMax
		}
		if hpMin < minHp {
			minHp = hpMin
		}
		if hpMax > maxHp {
			maxHp = hpMax
		}
	}
	return
}

func (filter *PokemonFilter) MinMaxSdCd() (minSd, maxSd, minCd, maxCd int, sds, cds []int) {
	if len(filter.specs) == 0 {
		return 0, 0, 0, 0, []int{}, []int{}
	}

	minSd, maxSd, minCd, maxCd = 2147483647, 0, 2147483647, 0
	sds, cds = []int{}, []int{}
	sdMap, cdMap := map[int]bool{}, map[int]bool{}
	for _, spec := range filter.specs {
		sd, cd := spec.lvl.StardustAndCandy()
		if !sdMap[sd] {
			sds = append(sds, sd)
			sdMap[sd] = true
		}
		if !cdMap[cd] {
			cds = append(cds, cd)
			cdMap[cd] = true
		}
		if sd < minSd {
			minSd = sd
		}
		if sd > maxSd {
			maxSd = sd
		}
		if cd < minCd {
			minCd = cd
		}
		if cd > maxCd {
			maxCd = cd
		}
	}
	sort.Ints(sds)
	sort.Ints(cds)
	return
}

func (filter *PokemonFilter) MinMaxLvl() (minLvl, maxLvl lv.Level) {
	if len(filter.specs) == 0 {
		return lv.Level(0.0), lv.Level(0.0)
	}

	minLvl, maxLvl = lv.Level(math.Inf(1)), lv.Level(0)
	for _, spec := range filter.specs {
		if spec.lvl < minLvl {
			minLvl = spec.lvl
		}
		if spec.lvl > maxLvl {
			maxLvl = spec.lvl
		}
	}
	return
}

func (filter *PokemonFilter) ForEach(consumer func(pokemon.IndividualValues, lv.Level)) {
	for _, spec := range filter.specs {
		consumer(spec.iv, spec.lvl)
	}
}

func match(value, min, max int) bool {
	return min <= value && value <= max
}
