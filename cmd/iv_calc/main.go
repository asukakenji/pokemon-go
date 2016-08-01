package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/asukakenji/pokemon-go/cmd"

	language "github.com/asukakenji/pokemon-go/lang"
	"github.com/asukakenji/pokemon-go/lv"
	"github.com/asukakenji/pokemon-go/pokemon"
)

type spec struct {
	pokemon.IndividualValues
	Level lv.Level
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Select Language
	lang := cmd.ReadLanguage(reader)

	mode := "N" // N = New, P = Power Up, E = Evolve
	pkm := pokemon.None
	cp := 0
	hp := 0
	sd := 0
	cd := 0
	wild := true
	ivs := []pokemon.IndividualValues(nil)

	for {
		// Select Pokémon
		switch mode {
		case "N":
			pkm = cmd.ReadPokemon(lang, reader, pkm)
			fmt.Printf("#%03d (%s)\n", pkm.Id(), pkm.LocalName(lang))
		case "P":
			// Keep previous value
		case "E":
			// TODO: Write this!
			pkmIterable := pkm.EvolveTo()

			// This should be implemented as ToSlice()
			pkms := make([]pokemon.Pokemon, 0)
			pkmIterable.ForEach(func (p pokemon.Pokemon) {
				pkms = append(pkms, p)
			})

			for _, p := range pkms {
				fmt.Println("%d: %s\n", p.Id(), p.LocalName(lang))
			}
			fmt.Println()

			// TODO: Add valid list
			pkm = cmd.ReadPokemon(lang, reader, pkm)
		default:
			panic("Invalid mode")
		}

		// Combat Power
		// TODO: Add valid range
		cp = cmd.ReadInt(reader, "CP [%d]: ", cp, nil)

		// Hit Points
		// TODO: Add valid range
		hp = cmd.ReadInt(reader, "HP [%d]: ", hp, nil)

		// Stardust
		// TODO: Add valid list
		sd = cmd.ReadInt(reader, "Stardust [%d]: ", sd, nil)

		// Candy
		// TODO: Add valid list
		lvls_temp := lv.LevelsByStardust(sd)
		for _, lvl := range lvls_temp {
			cd = lvl.CandyToPowerUp()
			break
		}
		cd = cmd.ReadInt(reader, "Candy (0 = auto) [%d]: ", 0, nil)

		// Wild
		switch mode {
		case "N":
			wildString := cmd.ReadStringWithChoices(reader, "Wild? Yes (Y) / No (N) [%s]: ", "Y", "Y", "N")
			wild = strings.ToUpper(wildString) != "N"
		case "P":
			wild = false
		case "E":
			// Keep previous value
		default:
			panic("Invalid mode")
		}

		specs := make([]spec, 0)

		if mode == "N" {
			ivs = allIVs()
		} else {
			// Keep previous value
		}

		lvls := lv.LevelsByStardustAndCandy(sd, cd)
		ivs2 := make([]pokemon.IndividualValues, 0)
		for _, iv := range ivs {
			isMatched := false
			for _, lvl := range lvls {
				if wild && !lvl.IsInteger() {
					continue
				}
				if matchCpAndHp(pkm, iv, lvl, cp, hp) {
					specs = append(specs, spec{iv, lvl})
					isMatched = true
				}
			}
			if isMatched {
				ivs2 = append(ivs2, iv)
			}
		}
		ivs = ivs2
		printSpecs(lang, pkm, cp, hp, sd, cd, wild, specs)

		// Mode
		mode = cmd.ReadStringWithChoices(reader, "New (N) / Power Up (P) / Evolve (E) [%s]: ", "N", "N", "P", "E")
	}
}

func printSpecs(lang language.Language, pkm pokemon.Pokemon, cp, hp, sd, cd int, wild bool, specs []spec) {
	fmt.Println()
	fmt.Printf("%s\n", "Input")
	fmt.Printf("%s\n", "-----")
	fmt.Printf("Pokémon: #%03d (%s)\n", pkm.Id(), pkm.LocalName(lang))
	fmt.Printf("CP: %d\n", cp)
	fmt.Printf("HP: %d\n", hp)
	fmt.Printf("Stardust: %d\n", sd)
	fmt.Printf("Candy: %d\n", cd)
	wildString := ""
	if wild {
		wildString = "Yes"
	} else {
		wildString = "No"
	}
	fmt.Printf("Wild: %s\n", wildString)
	fmt.Printf("%s\n", "-----")
	fmt.Println()
	count := len(specs)
	perfectionSum := 0.0
	perfectionMin := 100.0
	perfectionMax := 0.0
	fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", "Level", "Attack", "Defense", "Stamina", "CP Min", "CP Max", "HP Min", "HP Max", "Perfection")
	fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", "-----", "------", "-------", "-------", "------", "------", "------", "------", "----------")
	for _, spec := range specs {
		perfection := float64(spec.Attack+spec.Defense+spec.Stamina) * 100.0 / 45.0
		perfectionSum += perfection
		if perfection < perfectionMin {
			perfectionMin = perfection
		}
		if perfection > perfectionMax {
			perfectionMax = perfection
		}
		cpMin, cpMax, hpMin, hpMax := pkm.RawCombatPowerAndHitPoints(spec.IndividualValues, spec.Level)
		fmt.Printf("%.1f\t%d\t%d\t%d\t%.2f\t%.2f\t%.2f\t%.2f\t%.2f%%\n", spec.Level, spec.Attack, spec.Defense, spec.Stamina, cpMin, cpMax, hpMin, hpMax, perfection)
	}
	fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", "-----", "------", "-------", "-------", "------", "------", "------", "------", "----------")
	fmt.Printf("Number of results: %d\n", count)
	if count == 0 {
		fmt.Printf("Perfection (Min, Avg, Max): %f%%, %f%%, %f%%\n", 0.0, 0.0, 0.0)
	} else {
		fmt.Printf("Perfection (Min, Avg, Max): %f%%, %f%%, %f%%\n", perfectionMin, perfectionSum/float64(count), perfectionMax)
	}
	fmt.Println()
}

func matchCpAndHp(pkm pokemon.Pokemon, iv pokemon.IndividualValues, lvl lv.Level, cp, hp int) bool {
	cpMin, cpMax, hpMin, hpMax := pkm.CombatPowerAndHitPoints(iv, lvl)
	if lvl.IsInteger() {
		return cpMin == cp && hpMin == hp
	}
	return cpMin < cp && cp < cpMax && hpMin < hp && hp < hpMax
}

func allLevels() []lv.Level {
	result := make([]lv.Level, (40-1)*2)
	for level := lv.Level(float32(1.0)); level <= 40.0; level += 0.5 {
		result = append(result, level)
	}
	return result
}

func allIVs() []pokemon.IndividualValues {
	result := make([]pokemon.IndividualValues, 0, 16*16*16)
	for stamina := 0; stamina <= 15; stamina++ {
		for attack := 0; attack <= 15; attack++ {
			for defense := 0; defense <= 15; defense++ {
				result = append(result, pokemon.IndividualValues{stamina, attack, defense})
			}
		}
	}
	return result
}
