package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/asukakenji/pokemon-go/lang"
	"github.com/asukakenji/pokemon-go/pokemon"
)

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

		fmt.Print("Enter Level (1.0 ~ 40.0): ")
		level := float32(0.0)
		if n, err := fmt.Fscanf(reader, "%f\n", &level); n != 1 || err != nil {
			panic(err)
		}
		if level < 1.0 || level > 40.0 {
			panic("Invalid Level")
		}

		fmt.Print("Enter Attack (0 ~ 15): ")
		attack := 0
		if n, err := fmt.Fscanf(reader, "%d\n", &attack); n != 1 || err != nil {
			panic(err)
		}
		if attack < 0 || attack > 15 {
			panic("Invalid Attack")
		}

		fmt.Print("Enter Defense (0 ~ 15): ")
		defense := 0
		if n, err := fmt.Fscanf(reader, "%d\n", &defense); n != 1 || err != nil {
			panic(err)
		}
		if defense < 0 || defense > 15 {
			panic("Invalid Defense")
		}

		fmt.Print("Enter Stamina (0 ~ 15): ")
		stamina := 0
		if n, err := fmt.Fscanf(reader, "%d\n", &stamina); n != 1 || err != nil {
			panic(err)
		}
		if stamina < 0 || stamina > 15 {
			panic("Invalid Stamina")
		}

		iv := pokemon.IndividualValues{stamina, attack, defense}
		fmt.Printf("CP: %d\n", p.CombatPower(iv, level))
		fmt.Printf("HP: %d\n", p.HitPoints(iv, level))
	}
}
