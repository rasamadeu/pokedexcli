package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	
	// Create scanner pointer that reads from stdin
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			return 
		}
		command := cleanInput(scanner.Text())[0]
		fmt.Println("Your command was:", command)
	}

	return
}
