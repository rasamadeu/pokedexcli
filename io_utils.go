package main

import (
	"errors"
	"strings"
)

func cleanInput(text string) ([]string, error) {
	if len(text) == 0 {
		return make([]string, 0), errors.New("Provided empty string as input")
	}
	sliceWords := strings.Fields(text)
	for i := range sliceWords{
		sliceWords[i] = strings.ToLower(sliceWords[i])
	}
	return sliceWords, nil
}
