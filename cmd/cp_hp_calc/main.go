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
	reader := bufio.NewReader(os.Stdin)

	// Select Language
	lang := cmd.ReadLanguage(reader)

	pkm := pokemon.None
	lvl := lv.Level(1.0)
	atk := 0
	def := 0
	sta := 0
	for {
		// Select Pokémon
		pkm = cmd.ReadPokemon(lang, reader, pkm, nil)
		fmt.Printf("#%03d (%s)\n", pkm.Id(), pkm.LocalName(lang))

		// Enter Level
		lvl := lv.Level(cmd.ReadFloat32(reader, "Level (1.0 ~ 40.0) [%.1f]: ", float32(lvl), checkLvl))

		// Attack
		atk := cmd.ReadInt(reader, "Attack (0 ~ 15) [%d]: ", atk, checkAtkDefSta)

		// Defense
		def := cmd.ReadInt(reader, "Defense (0 ~ 15) [%d]: ", def, checkAtkDefSta)

		// Stamina
		sta := cmd.ReadInt(reader, "Stamina (0 ~ 15) [%d]: ", sta, checkAtkDefSta)

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
