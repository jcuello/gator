package main

import (
	"context"
	"fmt"
	"strings"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	currentUser := strings.ToLower(s.cfg.CurrentUsername)
	for _, user := range users {
		username := strings.ToLower(user.Name)

		if currentUser == username {
			username += " (current)"
		}
		fmt.Printf("* %v\n", username)
	}
	return nil
}
