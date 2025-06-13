package command

type config struct {
	next     string
	previous string
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

// This function return a map containing all possible Cli commands
func CreateCommands() map[string]CliCommand {

	// Pointer that will be shared by map and bmap commands
	configsMap := &config{
		next:     "",
		previous: "",
	}

	commands := map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the pokedex",
			Callback:    commandExit,
		},
	}

	commands["help"] = CliCommand{
		Name:        "help",
		Description: "Displays a help message",
		Callback:    commandHelp(commands),
	}

	commands["map"] = CliCommand{
		Name:        "map",
		Description: "Displays the next 20 locations",
		Callback:    commandMap(commands, configsMap),
	}

	commands["bmap"] = CliCommand{
		Name:        "bmap",
		Description: "Displays the previous 20 locations",
		Callback:    commandBmap(commands, configsMap),
	}

	return commands
}
