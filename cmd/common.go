package cmd

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	language "github.com/asukakenji/pokemon-go/lang"
	"github.com/asukakenji/pokemon-go/pokemon"
)

func ReadLanguage(scanner *bufio.Scanner) (lang language.Language, ok bool) {
	for {
		lang := language.Language(0)
		fmt.Println()
		language.All().ForEach(func(lang language.Language) {
			fmt.Printf("%d: %s\n", lang.Id(), lang.LocalName())
		})
		fmt.Println()
		fmt.Print("> ")
		if !scanner.Scan() {
			return language.Language(0), false
		}
		input := scanner.Text()
		if i, err := strconv.Atoi(input); err != nil {
			continue
		} else {
			lang = language.Language(i)
		}
		if !lang.IsValid() {
			continue
		}
		fmt.Println()
		return lang, true
	}
}

func ReadPokemon(lang language.Language, scanner *bufio.Scanner, defaultValue pokemon.Pokemon, printChoices func(), isValid func(pokemon.Pokemon) bool) (pkm pokemon.Pokemon, ok bool) {
	var msgMap = map[language.Language][]string{
		language.Japanese: {
			"%s (なまえまたはずかん番号): ",
			"%s (なまえまたはずかん番号) [%s]: ",
			"%sが見つからなかったよ。\n",
			"無効入力。もう一度入力してください。\n",
		},
		language.English: {
			"%s (name or number): ",
			"%s (name or number) [%s]: ",
			"No %s matched your search!\n",
			"Invalid input. Please enter again.\n",
		},
		language.French: {
			"%s (nom ou numéro): ",
			"%s (nom ou numéro) [%s]: ",
			"Aucun %s n'a été trouvé!\n",
			"Entrée invalide. S'il vous plaît entrez de nouveau.\n",
		},
		language.German: {
			"%s (name oder nummer): ",
			"%s (name oder nummer) [%s]: ",
			"Mit deinen Suchkriterien wurde kein %s gefunden!\n",
			"Ungültige Eingabe. Bitte nochmal eingeben.\n",
		},
		language.Italian: {
			"%s (nome o numero): ",
			"%s (nome o numero) [%s]: ",
			"Nessun %s corrisponde alla ricerca!\n",
			"Inserimento non valido. Si prega di inserire di nuovo.\n",
		},
		language.Korean: {
			"%s (이름 또는 도감 번호): ",
			"%s (이름 또는 도감 번호) [%s]: ",
			"어떤 %s은 검색 결과가 없습니다!\n",
			"잘못된 입력. 다시 입력하십시오.\n",
		},
		language.Spanish: {
			"%s (nombre o número): ",
			"%s (nombre o número) [%s]: ",
			"Sin %s coincide con su búsqueda!\n",
			"Entrada inválida. Por favor, introduzca de nuevo.\n",
		},
		language.ChineseSimplified: {
			"%s（名字 或者 图鉴编号）：",
			"%s（名字 或者 图鉴编号）[%s]：",
			"找不到这个%s\n",
			"输入无效。请重新输入。\n",
		},
		language.ChineseTraditional: {
			"%s（名字 或者 圖鑑編號）：",
			"%s（名字 或者 圖鑑編號）[%s]：",
			"找不到這個%s\n",
			"輸入錯誤。請重新輸入。\n",
		},
		language.ChineseChina: {
			"%s（名字 或者 图鉴编号）：",
			"%s（名字 或者 图鉴编号）[%s]：",
			"找不到这个%s\n",
			"输入无效。请重新输入。\n",
		},
		language.ChineseHongKong: {
			"%s（名 或者 圖鑑編號）：",
			"%s（名 或者 圖鑑編號）[%s]：",
			"搵唔到呢隻%s\n",
			"入錯嘢。再入過吖唔該。\n",
		},
		language.ChineseTaiwan: {
			"%s（名字 或者 圖鑑編號）：",
			"%s（名字 或者 圖鑑編號）[%s]：",
			"找不到這個%s\n",
			"輸入錯誤。請重新輸入。\n",
		},
	}
	msgs := msgMap[lang]
	for {
		if printChoices != nil {
			printChoices()
		}
		value := defaultValue
		if value == pokemon.None {
			fmt.Printf(msgs[0], pokemon.PackageName(lang))
		} else {
			fmt.Printf(msgs[1], pokemon.PackageName(lang), PokemonIdAndName(lang, value))
		}
		if !scanner.Scan() {
			return pokemon.None, false
		}
		input := scanner.Text()
		if input != "" {
			if i, err := strconv.Atoi(input); err != nil {
				// TODO: Rewrite the matching algorithm!
				value = pokemon.ByCodeName(input)
			} else {
				value = pokemon.Pokemon(i)
			}
		}
		if !value.IsValid() {
			fmt.Printf(msgs[2], pokemon.PackageName(lang))
			continue
		}
		if isValid != nil && !isValid(value) {
			fmt.Printf(msgs[3])
			continue
		}
		return value, true
	}
}

func ReadPokemonWithChoices(lang language.Language, scanner *bufio.Scanner, defaultValue pokemon.Pokemon, validValues ...pokemon.Pokemon) (pkm pokemon.Pokemon, ok bool) {
	printChoices := func() {
		fmt.Println()
		for _, v := range validValues {
			fmt.Printf("%s\n", PokemonIdAndName(lang, v))
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
	return ReadPokemon(lang, scanner, defaultValue, printChoices, isValid)
}

func ReadString(scanner *bufio.Scanner, prompt string, defaultValue string, printChoices func(), isValid func(string) bool) (s string, ok bool) {
	for {
		if printChoices != nil {
			printChoices()
		}
		value := defaultValue
		fmt.Printf(prompt, defaultValue)
		if !scanner.Scan() {
			return "", false
		}
		input := scanner.Text()
		if input != "" {
			value = input
		}
		if isValid != nil && !isValid(value) {
			fmt.Println("Invalid input")
			continue
		}
		return value, true
	}
}

func ReadStringWithChoices(scanner *bufio.Scanner, prompt string, defaultValue string, validValues ...string) (s string, ok bool) {
	printChoices := func() {
		fmt.Println()
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
	return ReadString(scanner, prompt, defaultValue, printChoices, isValid)
}

func ReadStringWithChoices2(scanner *bufio.Scanner, prompt string, defaultValue string, choices []struct {
	Choice string
	Desc   string
}) (s string, ok bool) {
	printChoices := func() {
		fmt.Println()
		fmt.Println(prompt)
		fmt.Println()
		for _, v := range choices {
			fmt.Printf("%s: %s\n", v.Choice, v.Desc)
		}
		fmt.Println()
	}
	isValid := func(value string) bool {
		for _, v := range choices {
			if strings.ToUpper(value) == strings.ToUpper(v.Choice) {
				return true
			}
		}
		return false
	}
	return ReadString(scanner, "[%s]: ", defaultValue, printChoices, isValid)
}

func ReadInt(scanner *bufio.Scanner, prompt string, defaultValue int, printChoices func(), isValid func(int) bool) (i int, ok bool) {
	for {
		if printChoices != nil {
			printChoices()
		}
		value := defaultValue
		fmt.Printf(prompt, defaultValue)
		if !scanner.Scan() {
			return 0, false
		}
		input := scanner.Text()
		if input != "" {
			if i, err := strconv.Atoi(input); err != nil {
				continue
			} else {
				value = i
			}
		}
		if isValid != nil && !isValid(value) {
			fmt.Println("Invalid input")
			continue
		}
		return value, true
	}
}

func ReadIntWithChoices(scanner *bufio.Scanner, prompt string, defaultValue int, validValues ...int) (i int, ok bool) {
	printChoices := func() {
		fmt.Println()
		for _, v := range validValues {
			fmt.Printf("%d\n", v)
		}
		fmt.Println()
	}
	isValid := func(value int) bool {
		for _, v := range validValues {
			if value == v {
				return true
			}
		}
		return false
	}
	return ReadInt(scanner, prompt, defaultValue, printChoices, isValid)
}

func ReadIntInRange(scanner *bufio.Scanner, prompt string, defaultValue, minValue, maxValue int) (i int, ok bool) {
	isValid := func(value int) bool {
		return minValue <= value && value <= maxValue
	}
	return ReadInt(scanner, prompt, defaultValue, nil, isValid)
}

func ReadFloat32(scanner *bufio.Scanner, prompt string, defaultValue float32, printChoices func(), isValid func(float32) bool) (f float32, ok bool) {
	for {
		if printChoices != nil {
			printChoices()
		}
		value := defaultValue
		fmt.Printf(prompt, defaultValue)
		if !scanner.Scan() {
			return 0.0, false
		}
		input := scanner.Text()
		if input != "" {
			if f, err := strconv.ParseFloat(input, 32); err != nil {
				continue
			} else {
				value = float32(f)
			}
		}
		if isValid != nil && !isValid(value) {
			fmt.Println("Invalid input")
			continue
		}
		return value, true
	}
}

func PokemonIdAndName(lang language.Language, pkm pokemon.Pokemon) string {
	if lang == language.English {
		return fmt.Sprintf("%03d: %s", pkm.Id(), pkm.LocalName(language.English))
	}
	return fmt.Sprintf("%03d: %s (%s)", pkm.Id(), pkm.LocalName(language.English), pkm.LocalName(lang))
}
