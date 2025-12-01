package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")

		scanner.Scan()

		text := scanner.Text()

		cleanText := cleanInput(text)
		
		if len(cleanText) == 0 {
			continue
		}

		cliCommands := getCommands()

		command, exists := cliCommands[cleanText[0]]

		if exists {
			err := command.callback(cfg)

			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

func cleanInput(text string) []string {

	return strings.Fields(strings.ToLower(text))

}
