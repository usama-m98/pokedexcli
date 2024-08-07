package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, name string) error {

	dat, err := cfg.pokeapiClient.GetLocations(cfg.next)
	if err != nil {
		return err
	}

	cfg.next = &dat.Next
	cfg.previous = dat.Previous
	for _, l := range dat.Results {
		fmt.Printf("%s\n", l.Name)
	}
	return nil
}

func commandMapb(cfg *config, name string) error {
	if cfg.previous == nil {
		return errors.New("you're on the first page")
	}
	dat, err := cfg.pokeapiClient.GetLocations(cfg.previous)
	if err != nil {
		return err
	}

	cfg.next = &dat.Next
	cfg.previous = dat.Previous

	for _, l := range dat.Results {
		fmt.Printf("%s\n", l.Name)
	}
	return nil
}
