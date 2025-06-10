package main

import (
	"strings"
)

func cleanInput(text string) []string {
	sliceWords := strings.Fields(text)
	for i := range sliceWords{
		sliceWords[i] = strings.ToLower(sliceWords[i])
	}
	return sliceWords
}
