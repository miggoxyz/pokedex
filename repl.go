package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/miggoxyz/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}


func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		cmd, exists := getCommands()[commandName]
		if exists {
			err := cmd.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unkown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name		string
	description	string
	callback	func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"help": {
			name:		 "help",
			description: "Displays a help message",
			callback:	 commandHelp,
		},
		"exit": {
			name: 		 "exit",
			description: "Exit pokedex",
			callback: 	 commandExit,
		},
		"clear": {
			name: 		 "clear",
			description: "Clears terminal",
			callback: 	 commandClear,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore":{
			name: "explore <location_name>",
			description: "Explore a location",
			callback: commandExplore,
		},
	}
}