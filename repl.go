package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/rasamadeu/pokedexcli/internal/command"
)

func replStart(){

	// Create scanner pointer that reads from stdin
	scanner := bufio.NewScanner(os.Stdin)

	commands := command.CreateCommands()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			return 
		}

		words, err := cleanInput(scanner.Text())
		if err != nil {
			continue
		}
		input := words[0]

		command, ok := commands[input]
		if ok == false {
			fmt.Println("Unknown command: insert \"help\" for guidance")
			continue
		}

		if err := command.Callback(); err != nil {
			fmt.Println(err)
		}
	}
}
