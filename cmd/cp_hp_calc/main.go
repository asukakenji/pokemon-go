package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/asukakenji/pokemon-go/cmd"

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

	pkm := pokemon.None
	lvl := lv.Level(1.0)
	atk := 0
	def := 0
	sta := 0
	for {
		// Select Pokémon
		pkm, ok = cmd.ReadPokemon(lang, scanner, pkm, nil, nil)
		if !ok {
			return
		}
		fmt.Printf("%s\n", cmd.PokemonIdAndName(lang, pkm))

		// Enter Level
		f, ok := cmd.ReadFloat32(scanner, "Level (1.0 ~ 40.0) [%.1f]: ", float32(lvl), nil, checkLvl)
		if !ok {
			return
		}
		lvl = lv.Level(f)

		// Attack
		atk, ok = cmd.ReadInt(scanner, "Attack (0 ~ 15) [%d]: ", atk, nil, checkAtkDefSta)
		if !ok {
			return
		}

		// Defense
		def, ok = cmd.ReadInt(scanner, "Defense (0 ~ 15) [%d]: ", def, nil, checkAtkDefSta)
		if !ok {
			return
		}

		// Stamina
		sta, ok = cmd.ReadInt(scanner, "Stamina (0 ~ 15) [%d]: ", sta, nil, checkAtkDefSta)
		if !ok {
			return
		}

		iv := pokemon.IndividualValues{sta, atk, def}
		cpMin, cpMax, hpMin, hpMax := pkm.CombatPowerAndHitPoints(iv, lvl)
		if lvl.IsInteger() {
			fmt.Printf("CP: %d\n", cpMin)
			fmt.Printf("HP: %d\n", hpMin)
		} else {
			fmt.Printf("CP: %d < cp < %d\n", cpMin, cpMax)
			fmt.Printf("HP: %d < hp < %d\n", hpMin, hpMax)
		}
	}
}

func checkLvl(value float32) bool {
	return lv.Level(value).IsValid()
}

func checkAtkDefSta(value int) bool {
	return 0 <= value && value <= 15
}
