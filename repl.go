package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/usama-m98/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, name string) error
}

type config struct {
	pokeapiClient pokeapi.Client
	pokedex       map[string]pokeapi.Pokemon
	next          *string
	previous      *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		commandName := words[0]

		command, exists := getCommands()[commandName]

		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		arg := ""
		if len(words) > 1 {
			arg = words[1]
		}

		err := command.callback(cfg, arg)
		if err != nil {
			fmt.Println(err.Error())
		}

		continue
	}
}

func cleanInput(word string) []string {
	output := strings.ToLower(word)
	words := strings.Fields(output)

	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore <area_name>",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch <pokemon>",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect <pokemon>",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List of caught Pokemon",
			callback:    commandPokedex,
		},
	}
}
