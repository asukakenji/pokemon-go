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
	reader := bufio.NewReader(os.Stdin)

	// Select Language
	lang := cmd.ReadLanguage(reader)

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
			pkm = cmd.ReadPokemon(lang, reader, pkm, nil)
			fmt.Printf("#%03d (%s)\n", pkm.Id(), pkm.LocalName(lang))
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

			// TODO: Display choices on every iteration
			for _, p := range pkms {
				fmt.Printf("%d: %s\n", p.Id(), p.LocalName(lang))
			}
			fmt.Println()

			for _, p := range pkms {
				pkm = p
				break
			}

			pkm = cmd.ReadPokemonWithChoices(lang, reader, pkm, pkms...)
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
		cp = cmd.ReadIntInRange(reader, cpMsg, cp, minCp, maxCp)
		filter.SetCp(cp)

		// Hit Points
		_, _, minHp, maxHp := filter.MinMaxCpHp()
		if !match(hp, minHp, maxHp) {
			hp = minHp
		}
		hpMsg := fmt.Sprintf("HP (%d ~ %d) [%%d]: ", minHp, maxHp)
		hp = cmd.ReadIntInRange(reader, hpMsg, hp, minHp, maxHp)
		filter.SetHp(hp)

		// Stardust
		minSd, maxSd, _, _, sds, _ := filter.MinMaxSdCd()
		if !match(sd, minSd, maxSd) {
			sd = minSd
		}
		sdMsg := fmt.Sprintf("Stardust (%d ~ %d) [%%d]: ", minSd, maxSd)
		sd = cmd.ReadIntWithChoices(reader, sdMsg, sd, sds...)
		filter.SetSd(sd)

		// Candy
		_, _, minCd, maxCd, _, cds := filter.MinMaxSdCd()
		if !match(cd, minCd, maxCd) {
			cd = minCd
		}
		cdMsg := fmt.Sprintf("Candy (%d ~ %d) [%%d]: ", minCd, maxCd)
		cd = cmd.ReadIntWithChoices(reader, cdMsg, cd, cds...)
		filter.SetCd(cd)

		// Wild
		switch mode {
		case "N":
			isWildString := strings.ToUpper(cmd.ReadStringWithChoices(reader, "Wild? Yes (Y) / No (N) [%s]: ", "Y", "Y", "N"))
			isWild = (isWildString != "N")
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
		mode = strings.ToUpper(cmd.ReadStringWithChoices(reader, "New (N) / Power Up (P) / Evolve (E) [%s]: ", mode, "N", "P", "E"))
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
