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
		input := cleanInput(scanner.Text())[0]
		command, ok := commands[input]
		if ok == false {
			fmt.Println("Unknown command: insert \"help\" for guidance")
			continue
		}
		err := command.Callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}
