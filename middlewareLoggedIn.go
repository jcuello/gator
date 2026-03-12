package main

import (
	"context"

	"github.com/jcuello/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(innerState *state, innerCmd command) error {
		user, err := innerState.db.GetUser(context.Background(), innerState.cfg.CurrentUsername)
		if err != nil {
			return err
		}
		return handler(innerState, innerCmd, user)
	}
}
