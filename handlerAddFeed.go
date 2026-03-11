package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jcuello/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return errors.New("addFeed expects two arguments <name> <url>")
	}
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
	if err != nil {
		return err
	}

	name, url := cmd.args[0], cmd.args[1]
	now := time.Now()
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", feed)

	return nil
}
