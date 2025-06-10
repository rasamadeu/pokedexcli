package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: " hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: " HELLO  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: " charmander  PIKACHU",
			expected: []string{"charmander", "pikachu"},
		},
		{
			input: " hello  world  CHARMANDER SHINY",
			expected: []string{"hello", "world", "charmander", "shiny"},
		},
	}
	for _, c := range cases {
		output := cleanInput(c.input)
		if len(output) != len(c.expected) {
			t.Errorf(
				"Failed test:\ninput: %s\nexpected: %v\noutput: %v\n",
				c.input, c.expected, output) 
		}
		for i := range output {
			if output[i] != c.expected[i] {
				t.Errorf("Failed test:\ninput: %s\nexpected: %v\noutput: %v\n",
					c.input, c.expected, output) 
			}
		}
	}
}
