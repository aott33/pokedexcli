package main

import (
	"strings"
	"bufio"
	"os"
	"fmt"
)


func startRepl() {
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
			err := command.callback()

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
