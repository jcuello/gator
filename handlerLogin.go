package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("login expects one argument")
	}
	username := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), username)

	if err == sql.ErrNoRows {
		fmt.Printf("user '%v' not found in database\n", username)
		os.Exit(1)
	}

	err = s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("The user has been set to %v\n", s.cfg.CurrentUsername)
	return nil
}
