package main

import (
	// Std library
	"bufio"
	"fmt"
	"os"
	"time"
	// Project
	"github.com/rasamadeu/pokedexcli/command"
	"github.com/rasamadeu/pokedexcli/internal/pokeapi"
	"github.com/rasamadeu/pokedexcli/internal/pokecache"
)

func replStart(){

	// Create scanner pointer that reads from stdin
	scanner := bufio.NewScanner(os.Stdin)

	// Configuration used by commands
	config := &command.Config{
		PokeapiClient: pokeapi.NewClient(5 * time.Second),
		Pokecache:     pokecache.NewCache(10 * time.Second),
	}

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

		if err := command.Callback(config); err != nil {
			fmt.Println(err)
		}
	}
}
