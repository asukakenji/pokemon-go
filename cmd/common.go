package cmd

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/asukakenji/pokemon-go/lang"
	"github.com/asukakenji/pokemon-go/pokemon"
)

func ReadLanguage(reader *bufio.Reader) lang.Language {
	for {
		l := lang.Language(0)
		fmt.Println()
		lang.All().ForEach(func(l lang.Language) {
			fmt.Printf("%d: %s\n", l.Id(), l.LocalName())
		})
		fmt.Println()
		fmt.Print("> ")
		if n, err := fmt.Fscanf(reader, "%d\n", &l); n != 1 || err != nil {
			fmt.Println(err)
			continue
		}
		if !l.IsValid() {
			continue
		}
		fmt.Println()
		return l
	}
}

func ReadPokemon(l lang.Language, reader *bufio.Reader, defaultValue pokemon.Pokemon, isValid func(pokemon.Pokemon) bool) pokemon.Pokemon {
	var msg = map[lang.Language][]string{
		lang.Japanese: {
			"%s (なまえまたはずかん番号): ",
			"%s (なまえまたはずかん番号) [#%03d (%s)]: ",
			"%sが見つからなかったよ。\n",
		},
		lang.English: {
			"%s (name or number): ",
			"%s (name or number) [#%03d (%s)]: ",
			"No %s matched your search!\n",
		},
		lang.French: {
			"%s (nom ou numéro): ",
			"%s (nom ou numéro) [#%03d (%s)]: ",
			"Aucun %s n'a été trouvé!\n",
		},
		lang.German: {
			"%s (name oder nummer): ",
			"%s (name oder nummer) [#%03d (%s)]: ",
			"Mit deinen Suchkriterien wurde kein %s gefunden!\n",
		},
		lang.Italian: {
			"%s (nome o numero): ",
			"%s (nome o numero) [#%03d (%s)]: ",
			"Nessun %s corrisponde alla ricerca!\n",
		},
		lang.Korean: {
			"%s (이름 또는 도감 번호): ",
			"%s (이름 또는 도감 번호) [#%03d (%s)]: ",
			"어떤 %s은 검색 결과가 없습니다!\n",
		},
		lang.Spanish: {
			"%s (nombre o número): ",
			"%s (nombre o número) [#%03d (%s)]: ",
			"Sin %s coincide con su búsqueda!\n",
		},
		lang.ChineseSimplified: {
			"%s（名字 或者 图鉴编号）：",
			"%s（名字 或者 图鉴编号）[#%03d (%s)]：",
			"找不到这个%s\n",
		},
		lang.ChineseTraditional: {
			"%s（名字 或者 圖鑑編號）：",
			"%s（名字 或者 圖鑑編號）[#%03d (%s)]：",
			"找不到這個%s\n",
		},
		lang.ChineseChina: {
			"%s（名字 或者 图鉴编号）：",
			"%s（名字 或者 图鉴编号）[#%03d (%s)]：",
			"找不到这个%s\n",
		},
		lang.ChineseHongKong: {
			"%s（名 或者 圖鑑編號）：",
			"%s（名 或者 圖鑑編號）[#%03d (%s)]：",
			"搵唔到呢隻%s\n",
		},
		lang.ChineseTaiwan: {
			"%s（名字 或者 圖鑑編號）：",
			"%s（名字 或者 圖鑑編號）[#%03d (%s)]：",
			"找不到這個%s\n",
		},
	}
	for {
		value := defaultValue
		input := ""
		if defaultValue == pokemon.None {
			fmt.Printf(msg[l][0], pokemon.PackageName(l))
			if n, err := fmt.Fscanf(reader, "%s\n", &input); n != 1 || err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			fmt.Printf(msg[l][1], pokemon.PackageName(l), defaultValue.Id(), defaultValue.LocalName(l))
			if line, err := reader.ReadString('\n'); err != nil {
				fmt.Println(err)
				continue
			} else {
				input = strings.TrimSuffix(strings.TrimSuffix(line, "\n"), "\r")
			}
		}
		if input != "" {
			if i, err := strconv.Atoi(input); err != nil {
				value = pokemon.ByCodeName(input)
			} else {
				value = pokemon.Pokemon(i)
			}
		}
		if !value.IsValid() {
			fmt.Printf(msg[l][2], pokemon.PackageName(l))
			continue
		}
		if isValid != nil && !isValid(value) {
			fmt.Println("Invalid input")
			continue
		}
		return value
	}
}

func ReadPokemonWithChoices(l lang.Language, reader *bufio.Reader, defaultValue pokemon.Pokemon, validValues ...pokemon.Pokemon) pokemon.Pokemon {
	return ReadPokemon(l, reader, defaultValue, func(value pokemon.Pokemon) bool {
		for _, validValue := range validValues {
			if value == validValue {
				return true
			}
		}
		return false
	})
}

func ReadString(reader *bufio.Reader, prompt string, defaultValue string, isValid func(string) bool) string {
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
			value = input
		}
		if isValid != nil && !isValid(value) {
			fmt.Println("Invalid input")
			continue
		}
		return value
	}
}

func ReadStringWithChoices(reader *bufio.Reader, prompt string, defaultValue string, validValues ...string) string {
	return ReadString(reader, prompt, defaultValue, func(value string) bool {
		for _, validValue := range validValues {
			if strings.ToUpper(value) == strings.ToUpper(validValue) {
				return true
			}
		}
		return false
	})
}

func ReadInt(reader *bufio.Reader, prompt string, defaultValue int, isValid func(int) bool) int {
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
		if isValid != nil && !isValid(value) {
			fmt.Println("Invalid input")
			continue
		}
		return value
	}
}

func ReadIntWithChoices(reader *bufio.Reader, prompt string, defaultValue int, validValues ...int) int {
	return ReadInt(reader, prompt, defaultValue, func(value int) bool {
		for _, validValue := range validValues {
			if value == validValue {
				return true
			}
		}
		return false
	})
}

func ReadIntInRange(reader *bufio.Reader, prompt string, defaultValue, minValue, maxValue int) int {
	return ReadInt(reader, prompt, defaultValue, func(value int) bool {
		return minValue <= value && value <= maxValue
	})
}

func ReadFloat32(reader *bufio.Reader, prompt string, defaultValue float32, isValid func(float32) bool) float32 {
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
		if isValid != nil && !isValid(value) {
			fmt.Println("Invalid input")
			continue
		}
		return value
	}
}
