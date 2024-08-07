package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, name string) error {
	if name == "" {
		return errors.New("unknown command")
	}

	pokemons, err := cfg.pokeapiClient.ListPokemonsInArea(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", name)
	fmt.Println("Found Pokemon:")
	for _, p := range pokemons.PokemonEncounters {
		fmt.Printf(" - %s\n", p.Pokemon.Name)
	}
	return nil
}
