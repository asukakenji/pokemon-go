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
	lvl := lv.Level(1.0)
	atk := 0
	def := 0
	sta := 0
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

		// Enter Level
		for {
			lvl2 := lv.Level(readFloat32(reader, "Level (1.0 ~ 40.0) [%.1f]: ", float32(lvl)))
			if !lvl2.IsValid() {
				fmt.Println("Invalid Level")
				continue
			}
			lvl = lvl2
			fmt.Printf("%s\n", lvl)
			break
		}

		// Attack
		for {
			atk2 := readInt(reader, "Attack (0 ~ 15) [%d]: ", atk)
			if atk2 < 0 || atk2 > 15 {
				fmt.Println("Invalid Attack")
				continue
			}
			atk = atk2
			fmt.Printf("%d\n", atk)
			break
		}

		// Defense
		for {
			def2 := readInt(reader, "Defense (0 ~ 15) [%d]: ", def)
			if def2 < 0 || def2 > 15 {
				fmt.Println("Invalid Defense")
				continue
			}
			def = def2
			fmt.Printf("%d\n", def)
			break
		}

		// Stamina
		for {
			sta2 := readInt(reader, "Stamina (0 ~ 15) [%d]: ", sta)
			if sta2 < 0 || sta2 > 15 {
				fmt.Println("Invalid Stamina")
				continue
			}
			sta = sta2
			fmt.Printf("%d\n", sta)
			break
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

func readFloat32(reader *bufio.Reader, prompt string, defaultValue float32) float32 {
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
			if f, err := strconv.ParseFloat(input, 32); err != nil {
				fmt.Println(err)
				continue
			} else {
				value = float32(f)
			}
		}
		return value
	}
}
