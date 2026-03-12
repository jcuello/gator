package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jcuello/gator/internal/config"
	"github.com/jcuello/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()

	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)
	appState := &state{cfg: &cfg, db: dbQueries}
	cliCommands := commands{cmds: map[string]func(*state, command) error{}}
	cliCommands.register("login", handlerLogin)
	cliCommands.register("register", handlerRegister)
	cliCommands.register("reset", handlerReset)
	cliCommands.register("users", handlerUsers)
	cliCommands.register("agg", handlerAgg)
	cliCommands.register("feeds", handlerFeeds)
	cliCommands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cliCommands.register("follow", middlewareLoggedIn(handlerFollow))
	cliCommands.register("following", middlewareLoggedIn(handlerFollowing))

	if len(os.Args) < 2 {
		fmt.Println("Usage: gator <command> [args...]")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cliCommands.run(appState, command{name: cmdName, args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
