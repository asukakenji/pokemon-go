package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"golang.org/x/net/html"
)

type Pokemon struct {
	Id   int
	Name string
}

func atoi(s string) int {
	if i, err := strconv.Atoi(s); err != nil {
		panic(err)
	} else {
		return i
	}
}

func getUrl(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return body
}

func parsePokemonsForPokemonCom(langUrl string, body []byte) []*Pokemon {
	pokemons := make([]*Pokemon, 0, 1000)
	reString := fmt.Sprintf(`<li><a href="/%s/pokedex/.+">([0-9]+) - (.*)</a></li>`, langUrl)
	re := regexp.MustCompile(reString)
	submatches := re.FindAllSubmatch(body, -1)
	for _, submatch := range submatches {
		id := atoi(string(submatch[1]))
		name := html.UnescapeString(string(submatch[2]))
		pokemons = append(pokemons, &Pokemon{id, name})
	}
	return pokemons
}

func generateSource(w io.Writer, lang string, pokemons []*Pokemon) {
	maxLength := 0
	for _, pokemon := range pokemons {
		length := utf8.RuneCountInString(pokemon.Name)
		if length > maxLength {
			maxLength = length
		}
	}

	fmt.Fprintf(w, "package %s\n", lang)
	fmt.Fprintln(w)
	fmt.Fprintln(w, "var Names = [...]string{")
	fmt.Fprintln(w, `	"",`)
	for _, pokemon := range pokemons {
		length := utf8.RuneCountInString(pokemon.Name)
		fmt.Fprintf(
			w,
			"	%#v,%s // %d\n",
			pokemon.Name,
			strings.Repeat(" ", maxLength-length),
			pokemon.Id,
		)
	}
	fmt.Fprintln(w, "}")
}

var languageMap = map[string]string{
	"de": "de",
	"en": "us",
	"es": "es",
	"fr": "fr",
	"it": "it",
}

func main() {
	lang := ""
	langUrl := ""
	flag.StringVar(&lang, "lang", "", "Language, one of: de, en, es, fr, it")
	flag.Parse()
	if lang == "" {
		fmt.Fprintln(os.Stderr, "No language is defined")
		os.Exit(1)
	} else {
		ok := false
		if langUrl, ok = languageMap[lang]; !ok {
			fmt.Fprintf(os.Stderr, "Invalid language: %s\n", lang)
			os.Exit(1)
		}
	}

	url := fmt.Sprintf("http://www.pokemon.com/%s/pokedex/", langUrl)
	body := getUrl(url)

	pokemons := parsePokemonsForPokemonCom(langUrl, body)
	generateSource(os.Stdout, lang, pokemons)
}
