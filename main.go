package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/jcuello/gator/internal/config"
)

func main() {
	cfg, err := config.Read()

	if err != nil {
		printErrorAndExit(err)
	}

	appState := state{cfg: &cfg}
	cliCommands := commands{cmds: map[string]func(*state, command) error{}}
	cliCommands.register("login", handlerLogin)

	cliArgs := os.Args
	argCount := len(cliArgs)

	if argCount < 2 {
		printErrorAndExit(errors.New("error: not enough arguments provided"))
	}

	if argCount < 3 {
		printErrorAndExit(errors.New("error: login command requires a username"))
	}

	err = cliCommands.run(&appState, command{name: "login", args: cliArgs[2:]})
	if err != nil {
		printErrorAndExit(err)
	}

}

func printErrorAndExit(err error) {
	fmt.Printf("%v\n", err)
	os.Exit(1)
}
