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
			continue
		}
		if !l.IsValid() {
			continue
		}
		fmt.Println()
		return l
	}
}

func ReadPokemon(l lang.Language, reader *bufio.Reader, defaultValue pokemon.Pokemon, printChoices func(), isValid func(pokemon.Pokemon) bool) pokemon.Pokemon {
	var msgMap = map[lang.Language][]string{
		lang.Japanese: {
			"%s (なまえまたはずかん番号): ",
			"%s (なまえまたはずかん番号) [%s]: ",
			"%sが見つからなかったよ。\n",
			"無効入力。もう一度入力してください。\n",
		},
		lang.English: {
			"%s (name or number): ",
			"%s (name or number) [%s]: ",
			"No %s matched your search!\n",
			"Invalid input. Please enter again.\n",
		},
		lang.French: {
			"%s (nom ou numéro): ",
			"%s (nom ou numéro) [%s]: ",
			"Aucun %s n'a été trouvé!\n",
			"Entrée invalide. S'il vous plaît entrez de nouveau.\n",
		},
		lang.German: {
			"%s (name oder nummer): ",
			"%s (name oder nummer) [%s]: ",
			"Mit deinen Suchkriterien wurde kein %s gefunden!\n",
			"Ungültige Eingabe. Bitte nochmal eingeben.\n",
		},
		lang.Italian: {
			"%s (nome o numero): ",
			"%s (nome o numero) [%s]: ",
			"Nessun %s corrisponde alla ricerca!\n",
			"Inserimento non valido. Si prega di inserire di nuovo.\n",
		},
		lang.Korean: {
			"%s (이름 또는 도감 번호): ",
			"%s (이름 또는 도감 번호) [%s]: ",
			"어떤 %s은 검색 결과가 없습니다!\n",
			"잘못된 입력. 다시 입력하십시오.\n",
		},
		lang.Spanish: {
			"%s (nombre o número): ",
			"%s (nombre o número) [%s]: ",
			"Sin %s coincide con su búsqueda!\n",
			"Entrada inválida. Por favor, introduzca de nuevo.\n",
		},
		lang.ChineseSimplified: {
			"%s（名字 或者 图鉴编号）：",
			"%s（名字 或者 图鉴编号）[%s]：",
			"找不到这个%s\n",
			"输入无效。请重新输入。\n",
		},
		lang.ChineseTraditional: {
			"%s（名字 或者 圖鑑編號）：",
			"%s（名字 或者 圖鑑編號）[%s]：",
			"找不到這個%s\n",
			"輸入錯誤。請重新輸入。\n",
		},
		lang.ChineseChina: {
			"%s（名字 或者 图鉴编号）：",
			"%s（名字 或者 图鉴编号）[%s]：",
			"找不到这个%s\n",
			"输入无效。请重新输入。\n",
		},
		lang.ChineseHongKong: {
			"%s（名 或者 圖鑑編號）：",
			"%s（名 或者 圖鑑編號）[%s]：",
			"搵唔到呢隻%s\n",
			"入錯嘢。再入過吖唔該。\n",
		},
		lang.ChineseTaiwan: {
			"%s（名字 或者 圖鑑編號）：",
			"%s（名字 或者 圖鑑編號）[%s]：",
			"找不到這個%s\n",
			"輸入錯誤。請重新輸入。\n",
		},
	}
	msgs := msgMap[l]
	for {
		if printChoices != nil {
			printChoices()
		}
		value := defaultValue
		input := ""
		if value == pokemon.None {
			fmt.Printf(msgs[0], pokemon.PackageName(l))
			if n, err := fmt.Fscanf(reader, "%s\n", &input); n != 1 || err != nil {
				continue
			}
		} else {
			fmt.Printf(msgs[1], pokemon.PackageName(l), PokemonIdAndName(l, value))
			if line, err := reader.ReadString('\n'); err != nil {
				continue
			} else {
				input = strings.TrimSuffix(strings.TrimSuffix(line, "\n"), "\r")
			}
		}
		if input != "" {
			if i, err := strconv.Atoi(input); err != nil {
				// TODO: Rewrite the matching algorithm!
				value = pokemon.ByCodeName(input)
			} else {
				value = pokemon.Pokemon(i)
			}
		}
		if !value.IsValid() {
			fmt.Printf(msgs[2], pokemon.PackageName(l))
			continue
		}
		if isValid != nil && !isValid(value) {
			fmt.Printf(msgs[3])
			continue
		}
		return value
	}
}

func ReadPokemonWithChoices(l lang.Language, reader *bufio.Reader, defaultValue pokemon.Pokemon, validValues ...pokemon.Pokemon) pokemon.Pokemon {
	printChoices := func() {
		for _, v := range validValues {
			fmt.Printf("%s\n", PokemonIdAndName(l, v))
		}
		fmt.Println()
	}
	isValid := func(value pokemon.Pokemon) bool {
		for _, v := range validValues {
			if value == v {
				return true
			}
		}
		return false
	}
	return ReadPokemon(l, reader, defaultValue, printChoices, isValid)
}

func ReadString(reader *bufio.Reader, prompt string, defaultValue string, printChoices func(), isValid func(string) bool) string {
	for {
		if printChoices != nil {
			printChoices()
		}
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
	printChoices := func() {
		for _, v := range validValues {
			fmt.Printf("%s\n", v)
		}
		fmt.Println()
	}
	isValid := func(value string) bool {
		for _, v := range validValues {
			if strings.ToUpper(value) == strings.ToUpper(v) {
				return true
			}
		}
		return false
	}
	return ReadString(reader, prompt, defaultValue, printChoices, isValid)
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
		for _, v := range validValues {
			if value == v {
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

func PokemonIdAndName(l lang.Language, pkm pokemon.Pokemon) string {
	if l == lang.English {
		return fmt.Sprintf("%03d: %s", pkm.Id(), pkm.LocalName(lang.English))
	}
	return fmt.Sprintf("%03d: %s (%s)", pkm.Id(), pkm.LocalName(lang.English), pkm.LocalName(l))
}
