package main

import (
	"fmt"
	"github.com/PPrinsloo/pokedexcli/poke_api"
	"os"
)

type commandCallback func(commands map[string]cliCommand) error

type cliCommand struct {
	name        string
	description string
	callback    commandCallback
}

func commandHelp(commands map[string]cliCommand) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Available commands:\n")

	for _, command := range commands {
		fmt.Printf("%s: %s.\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandExit(commands map[string]cliCommand) error {
	fmt.Println(commands["exit"].description)
	os.Exit(0)
	return nil
}

func location(_ map[string]cliCommand) error {
	poke_api.Locations()
	return nil
}

func locations_back(_ map[string]cliCommand) error {
	poke_api.LocationsBack()
	return nil
}

func commands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exist the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "prints some locations names",
			callback:    location,
		},
		"mapb": {
			name:        "mapb",
			description: "prints previous locations names",
			callback:    locations_back,
		},
	}
}
