package command

import (
	"fmt"
	"os"
)

// Exit command callback function
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
