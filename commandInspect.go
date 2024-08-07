package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, name string) error {
	if name == "" {
		return errors.New("unknown command")
	}

	pokemon, ok := cfg.pokedex[name]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, p := range pokemon.Stats {
		fmt.Printf(" -%s: %v\n", p.Stat.Name, p.BaseStat)
	}

	fmt.Println("Types:")
	for _, p := range pokemon.Types {
		fmt.Printf(" - %s\n", p.Type.Name)
	}

	return nil
}
