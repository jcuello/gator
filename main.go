package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jcuello/gator/internal/config"
)

func main() {
	cfg, err := config.Read()

	if err != nil {
		log.Fatal(err)
	}

	appState := &state{cfg: &cfg}
	cliCommands := commands{cmds: map[string]func(*state, command) error{}}
	cliCommands.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %v <command> [args...]\n", os.Args[0])
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cliCommands.run(appState, command{name: cmdName, args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
