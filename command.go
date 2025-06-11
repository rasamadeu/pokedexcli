package main

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// This function return a map containing all possible CLI commands
func createCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
	}
	commands["help"] = cliCommand {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp(commands),
		}
	return commands
}
