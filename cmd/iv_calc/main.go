package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/asukakenji/pokemon-go/lang"
	"github.com/asukakenji/pokemon-go/lv"
	"github.com/asukakenji/pokemon-go/pokemon"
)

type spec struct {
	pokemon.IndividualValues
	level lv.Level
}

func main() {
	p := pokemon.None
	reader := bufio.NewReader(os.Stdin)
	for {
		pokemonString := ""
		if p == pokemon.None {
			fmt.Print("Enter Pokémon (name or number): ")
			if n, err := fmt.Fscanf(reader, "%s\n", &pokemonString); n != 1 || err != nil {
				panic(err)
			}
		} else {
			fmt.Printf("Enter Pokémon (name or number) [%d]: ", p.Id())
			if line, err := reader.ReadString('\n'); err != nil {
				panic(err)
			} else {
				pokemonString = strings.TrimSuffix(strings.TrimSuffix(line, "\n"), "\r")
			}
		}
		if pokemonString != "" {
			if i, err := strconv.Atoi(pokemonString); err != nil {
				p = pokemon.ByCodeName(pokemonString)
			} else {
				p = pokemon.Pokemon(i)
			}
		}
		if p <= pokemon.None || p > pokemon.Mew {
			panic("No such Pokémon")
		} else {
			fmt.Printf("#%03d: %s\n", p.Id(), p.LocalName(lang.English))
		}

		fmt.Print("Enter CP: ")
		combatPower := 0
		if n, err := fmt.Fscanf(reader, "%d\n", &combatPower); n != 1 || err != nil {
			panic(err)
		}

		fmt.Print("Enter HP: ")
		hitPoints := 0
		if n, err := fmt.Fscanf(reader, "%d\n", &hitPoints); n != 1 || err != nil {
			panic(err)
		}

		fmt.Print("Enter Stardust: ")
		stardust := 0
		if n, err := fmt.Fscanf(reader, "%d\n", &stardust); n != 1 || err != nil {
			panic(err)
		}

		fmt.Print("Enter Candy: ")
		candy := 0
		if n, err := fmt.Fscanf(reader, "%d\n", &candy); n != 1 || err != nil {
			panic(err)
		}

		specs := make([]spec, 0)
		ivs := allIVs()
		levels := lv.LevelsByStardustAndCandy(stardust, candy)
		for _, iv := range ivs {
			for _, level := range levels {
				if p.CombatPower(iv, level) == combatPower && p.HitPoints(iv, level) == hitPoints {
					specs = append(specs, spec{iv, level})
				}
			}
		}
		printSpecs(specs)
	}
}

func printSpecs(specs []spec) {
	count := len(specs)
	perfectionSum := 0.0
	perfectionMin := 100.0
	perfectionMax := 0.0
	fmt.Printf("%s\t%s\t%s\t%s\t%s\n", "Level", "Attack", "Defense", "Stamina", "Perfection")
	fmt.Printf("%s\t%s\t%s\t%s\t%s\n", "-----", "------", "-------", "-------", "----------")
	for _, spec := range specs {
		perfection := float64(spec.Attack+spec.Defense+spec.Stamina) * 100.0 / 45.0
		perfectionSum += perfection
		if perfection < perfectionMin {
			perfectionMin = perfection
		}
		if perfection > perfectionMax {
			perfectionMax = perfection
		}
		fmt.Printf("%.1f\t%d\t%d\t%d\t%f%%\n", spec.level, spec.Attack, spec.Defense, spec.Stamina, perfection)
	}
	fmt.Printf("%s\t%s\t%s\t%s\t%s\n", "-----", "------", "-------", "-------", "----------")
	fmt.Printf("Number of results: %d\n", count)
	fmt.Printf("Perfection (Min, Avg, Max): %f%%, %f%%, %f%%\n", perfectionMin, perfectionSum/float64(count), perfectionMax)
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
