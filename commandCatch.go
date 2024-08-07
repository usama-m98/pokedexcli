package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, name string) error {
	if name == "" {
		return errors.New("unknown command")
	}
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	randomNumber := rand.IntN(pokemon.BaseExperience)
	chance := int(float64(randomNumber) / float64(pokemon.BaseExperience) * 100)

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	if chance < 50 {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}

	fmt.Printf("%s was caught\n", name)
	fmt.Println("You may now inspect it with the inspect command.")
	cfg.pokedex[name] = pokemon
	return nil
}
