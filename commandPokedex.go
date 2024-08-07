package main

import "fmt"

func commandPokedex(cfg *config, name string) error {
	pokemon := cfg.pokedex

	if len(pokemon) == 0 {
		fmt.Println("No pokemon in Pokedex")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, p := range pokemon {
		fmt.Printf(" - %s\n", p.Name)
	}

	return nil
}
