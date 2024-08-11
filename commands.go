package main

import (
	"fmt"
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
	}
}
