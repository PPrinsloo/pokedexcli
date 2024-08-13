package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	commands := commands()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Pokedex>")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("Please enter a command.")
		}

		if command, ok := commands[input]; ok {
			if err := command.callback(commands); err != nil {
				fmt.Println("Error executing command:", err)
			}
		} else {
			fmt.Println("Unknown command. Type 'help' for a list of	commands.")
			continue
		}
	}
}
