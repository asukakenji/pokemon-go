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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Select Language
	lang, ok := cmd.ReadLanguage(scanner)
	if !ok {
		return
	}

	mode := "N" // N = New, P = Power Up, E = Evolve
	filter := NewPokemonFilter()
	pkm := pokemon.None
	cp := 0
	hp := 0
	sd := 0
	cd := 0
	isWild := true

	for {
		// Select Pokémon
		switch mode {
		case "N":
			// Reset pkm, specs
			pkm, ok = cmd.ReadPokemon(lang, scanner, pkm, nil, nil)
			if !ok {
				return
			}
			fmt.Printf("%s\n", cmd.PokemonIdAndName(lang, pkm))
			filter.Reset()
			filter.SetPokemon(pkm)
		case "P":
			// Keep pkm; Increment specs.lvl
			filter.IncrementLvl()
		case "E":
			// Keep specs; Reset pkm
			pkmIterable := pkm.EvolveTo()

			// This should be implemented as ToSlice()
			pkms := make([]pokemon.Pokemon, 0)
			pkmIterable.ForEach(func(p pokemon.Pokemon) {
				pkms = append(pkms, p)
			})

			if len(pkms) == 0 {
				// If the Pokémon is not evolve-able, add itself as the only choice;
				pkms = []pokemon.Pokemon{pkm}
			} else {
				// Otherwise, select the first choice as default.
				pkm = pkms[0]
			}

			pkm, ok = cmd.ReadPokemonWithChoices(lang, scanner, pkm, pkms...)
			if !ok {
				return
			}
			filter.SetPokemon(pkm)
		default:
			panic("Invalid mode")
		}

		// Combat Power
		minCp, maxCp, _, _ := filter.MinMaxCpHp()
		if !match(cp, minCp, maxCp) {
			cp = minCp
		}
		cpMsg := fmt.Sprintf("CP (%d ~ %d) [%%d]: ", minCp, maxCp)
		cp, ok = cmd.ReadIntInRange(scanner, cpMsg, cp, minCp, maxCp)
		if !ok {
			return
		}
		filter.SetCp(cp)

		// Hit Points
		_, _, minHp, maxHp := filter.MinMaxCpHp()
		if !match(hp, minHp, maxHp) {
			hp = minHp
		}
		hpMsg := fmt.Sprintf("HP (%d ~ %d) [%%d]: ", minHp, maxHp)
		hp, ok = cmd.ReadIntInRange(scanner, hpMsg, hp, minHp, maxHp)
		if !ok {
			return
		}
		filter.SetHp(hp)

		// Stardust
		minSd, maxSd, _, _, sds, _ := filter.MinMaxSdCd()
		if !match(sd, minSd, maxSd) {
			sd = minSd
		}
		sdMsg := fmt.Sprintf("Stardust (%d ~ %d) [%%d]: ", minSd, maxSd)
		sd, ok = cmd.ReadIntWithChoices(scanner, sdMsg, sd, sds...)
		if !ok {
			return
		}
		filter.SetSd(sd)

		// Candy
		_, _, minCd, maxCd, _, cds := filter.MinMaxSdCd()
		if !match(cd, minCd, maxCd) {
			cd = minCd
		}
		cdMsg := fmt.Sprintf("Candy (%d ~ %d) [%%d]: ", minCd, maxCd)
		cd, ok = cmd.ReadIntWithChoices(scanner, cdMsg, cd, cds...)
		if !ok {
			return
		}
		filter.SetCd(cd)

		// Wild
		switch mode {
		case "N":
			isWildString, ok := cmd.ReadStringWithChoices(scanner, "Wild? Yes (Y) / No (N) [%s]: ", "Y", "Y", "N")
			if !ok {
				return
			}
			isWild = (strings.ToUpper(isWildString) != "N")
		case "P":
			isWild = false
		case "E":
			// Keep previous value
		default:
			panic("Invalid mode")
		}
		filter.SetIsWild(isWild)
		fmt.Println()

		printInput(lang, pkm, cp, hp, sd, cd, isWild)
		printSpecs(lang, filter)

		// Mode
		if mode == "E" {
			mode = "P"
		}
		// TODO: Detect un-evolve-able Pokémons and change manual
		mode, ok = cmd.ReadStringWithChoices2(
			scanner,
			"What do you want to do next?",
			mode,
			[]struct {
				Choice string
				Desc   string
			}{
				{"N", "Calculate the next Pokémon"},
				{"P", "Power Up this Pokémon"},
				{"E", "Evolve this Pokémon"},
			},
		)
		if !ok {
			return
		}
		mode = strings.ToUpper(mode)
		fmt.Println()
	}
}

func printInput(lang language.Language, pkm pokemon.Pokemon, cp, hp, sd, cd int, isWild bool) {
	fmt.Printf("%s\n", "Input")
	fmt.Printf("%s\n", "-----")
	fmt.Printf("Pokémon: #%03d (%s)\n", pkm.Id(), pkm.LocalName(lang))
	fmt.Printf("CP: %d\n", cp)
	fmt.Printf("HP: %d\n", hp)
	fmt.Printf("Stardust: %d\n", sd)
	fmt.Printf("Candy: %d\n", cd)
	isWildString := ""
	if isWild {
		isWildString = "Yes"
	} else {
		isWildString = "No"
	}
	fmt.Printf("Wild: %s\n", isWildString)
	fmt.Printf("%s\n", "-----")
	fmt.Println()
}

func printSpecs(lang language.Language, filter *PokemonFilter) {
	count := 0
	perfectionSum := 0.0
	perfectionMin := 100.0
	perfectionMax := 0.0
	fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", "Level", "Attack", "Defense", "Stamina", "CP Min", "CP Max", "HP Min", "HP Max", "Perfection")
	fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", "-----", "------", "-------", "-------", "------", "------", "------", "------", "----------")
	filter.ForEach(func(iv pokemon.IndividualValues, lvl lv.Level) {
		perfection := float64(iv.Attack+iv.Defense+iv.Stamina) * 100.0 / 45.0
		perfectionSum += perfection
		if perfection < perfectionMin {
			perfectionMin = perfection
		}
		if perfection > perfectionMax {
			perfectionMax = perfection
		}
		cpMin, cpMax, hpMin, hpMax := filter.Pokemon().RawCombatPowerAndHitPoints(iv, lvl)
		fmt.Printf("%.1f\t%d\t%d\t%d\t%.2f\t%.2f\t%.2f\t%.2f\t%.2f%%\n", lvl, iv.Attack, iv.Defense, iv.Stamina, cpMin, cpMax, hpMin, hpMax, perfection)
		count++
	})
	fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", "-----", "------", "-------", "-------", "------", "------", "------", "------", "----------")
	fmt.Printf("Number of results: %d\n", count)
	if count == 0 {
		fmt.Printf("Perfection (Min, Avg, Max): %f%%, %f%%, %f%%\n", 0.0, 0.0, 0.0)
	} else {
		fmt.Printf("Perfection (Min, Avg, Max): %f%%, %f%%, %f%%\n", perfectionMin, perfectionSum/float64(count), perfectionMax)
	}
	fmt.Println()
}
