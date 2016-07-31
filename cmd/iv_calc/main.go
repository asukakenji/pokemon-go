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
	Level lv.Level
}

var msg = map[lang.Language][]string{
	lang.Japanese: {
		"%s (name or number): ",
		"%s (name or number) [#%03d (%s)]: ",
		"This %s is not found\n",
	},
	lang.English: {
		"%s (name or number): ",
		"%s (name or number) [#%03d (%s)]: ",
		"This %s is not found\n",
	},
	lang.French: {
		"%s (name or number): ",
		"%s (name or number) [#%03d (%s)]: ",
		"This %s is not found\n",
	},
	lang.German: {
		"%s (name or number): ",
		"%s (name or number) [#%03d (%s)]: ",
		"This %s is not found\n",
	},
	lang.Italian: {
		"%s (name or number): ",
		"%s (name or number) [#%03d (%s)]: ",
		"This %s is not found\n",
	},
	lang.Korean: {
		"%s (name or number): ",
		"%s (name or number) [#%03d (%s)]: ",
		"This %s is not found\n",
	},
	lang.Spanish: {
		"%s (name or number): ",
		"%s (name or number) [#%03d (%s)]: ",
		"This %s is not found\n",
	},
	lang.ChineseSimplified: {
		"%s（编号 或者 名字）：",
		"%s（编号 或者 名字）[#%03d (%s)]：",
		"找不到这个%s\n",
	},
	lang.ChineseTraditional: {
		"%s（編號 或者 名字）：",
		"%s（編號 或者 名字）[#%03d (%s)]：",
		"找不到這個%s\n",
	},
	lang.ChineseChina: {
		"%s（编号 或者 名字）：",
		"%s（编号 或者 名字）[#%03d (%s)]：",
		"找不到这个%s\n",
	},
	lang.ChineseHongKong: {
		"%s（冧把 或者 名）：",
		"%s（冧把 或者 名）[#%03d (%s)]：",
		"搵唔到呢隻%s\n",
	},
	lang.ChineseTaiwan: {
		"%s（編號 或者 名字）：",
		"%s（編號 或者 名字）[#%03d (%s)]：",
		"找不到這個%s\n",
	},
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Select Language
	l := lang.Language(0)
	for {
		lang.All().ForEach(func(l lang.Language) {
			fmt.Printf("%d: %s\n", l.Id(), l.LocalName())
		})
		fmt.Print("> ")
		if n, err := fmt.Fscanf(reader, "%d\n", &l); n != 1 || err != nil {
			fmt.Println(err)
			continue
		}
		if !l.IsValid() {
			continue
		}
		break
	}

	pkm := pokemon.None
	cp := 0
	hp := 0
	sd := 0
	cd := 0
	wild := true

	for {
		// Select Pokémon
		for {
			p2 := pkm
			pokemonString := ""
			if pkm == pokemon.None {
				fmt.Printf(msg[l][0], pokemon.PackageName(l))
				if n, err := fmt.Fscanf(reader, "%s\n", &pokemonString); n != 1 || err != nil {
					fmt.Println(err)
					continue
				}
			} else {
				fmt.Printf(msg[l][1], pokemon.PackageName(l), pkm.Id(), pkm.LocalName(l))
				if line, err := reader.ReadString('\n'); err != nil {
					fmt.Println(err)
					continue
				} else {
					pokemonString = strings.TrimSuffix(strings.TrimSuffix(line, "\n"), "\r")
				}
			}
			if pokemonString != "" {
				if i, err := strconv.Atoi(pokemonString); err != nil {
					p2 = pokemon.ByCodeName(pokemonString)
				} else {
					p2 = pokemon.Pokemon(i)
				}
			}
			if !p2.IsValid() {
				fmt.Printf(msg[l][2], pokemon.PackageName(l))
				continue
			}
			pkm = p2
			fmt.Printf("#%03d (%s)\n", pkm.Id(), pkm.LocalName(l))
			break
		}

		// Combat Power
		for {
			cp = readInt(reader, "CP [%d]: ", cp)
			break
		}

		// Hit Points
		for {
			hp = readInt(reader, "HP [%d]: ", hp)
			break
		}

		// Stardust
		for {
			sd = readInt(reader, "Stardust [%d]: ", sd)
			break
		}

		// Candy
		for {
			cd = readInt(reader, "Candy (0 = auto) [%d]: ", 0)
			break
		}

		// Wild
		for {
			wild = readBool(reader, "Wild? (Y/N) [%s]: ", true)
			break
		}

		specs := make([]spec, 0)

		var lvls []lv.Level
		if cd == 0 {
			lvls = lv.LevelsByStardust(sd)
			for _, lvl := range lvls {
				cd = lvl.CandyToPowerUp()
				break
			}
		} else {
			lvls = lv.LevelsByStardustAndCandy(sd, cd)
		}

		ivs := allIVs()
		for _, iv := range ivs {
			for _, lvl := range lvls {
				if wild && !lvl.IsInteger() {
					continue
				}
				if matchCpAndHp(pkm, iv, lvl, cp, hp) {
					specs = append(specs, spec{iv, lvl})
				}
			}
		}
		printSpecs(l, pkm, cp, hp, sd, cd, wild, specs)
	}
}

func printSpecs(l lang.Language, pkm pokemon.Pokemon, cp, hp, sd, cd int, wild bool, specs []spec) {
	fmt.Println()
	fmt.Printf("%s\n", "Input")
	fmt.Printf("%s\n", "-----")
	fmt.Printf("Pokémon: #%03d (%s)\n", pkm.Id(), pkm.LocalName(l))
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

func readInt(reader *bufio.Reader, prompt string, defaultValue int) int {
	for {
		value := defaultValue
		input := ""
		fmt.Printf(prompt, defaultValue)
		if line, err := reader.ReadString('\n'); err != nil {
			fmt.Println(err)
			continue
		} else {
			input = strings.TrimSuffix(strings.TrimSuffix(line, "\n"), "\r")
		}
		if input != "" {
			if i, err := strconv.Atoi(input); err != nil {
				fmt.Println(err)
				continue
			} else {
				value = i
			}
		}
		return value
	}
}

func readBool(reader *bufio.Reader, prompt string, defaultValue bool) bool {
	for {
		value := defaultValue
		input := ""
		defaultValueString := ""
		if defaultValue {
			defaultValueString = "Y"
		} else {
			defaultValueString = "N"
		}
		fmt.Printf(prompt, defaultValueString)
		if line, err := reader.ReadString('\n'); err != nil {
			fmt.Println(err)
			continue
		} else {
			input = strings.TrimSuffix(strings.TrimSuffix(line, "\n"), "\r")
		}
		if input != "" {
			switch input[0] {
			case 'N', 'n':
				value = false
			case 'Y', 'y':
				value = true
			default:
				continue
			}
		}
		return value
	}
}
