package main

import (
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
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	if len(os.Args) < 2 {
		log.Fatalf("Usage: %v <command> [args...]\n", cmdName)
	}

	err = cliCommands.run(appState, command{name: cmdName, args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
