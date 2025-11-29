package main

import (
	"strings"
)

func cleanInput(text string) []string {
	split := []string{}

	if len(text) == 0 {
		return split
	}

	split = strings.Fields(text)

	return split
	

}
