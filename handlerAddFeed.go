package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jcuello/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return errors.New("addFeed expects two arguments <name> <url>")
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

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    feed.UserID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", feed)

	return nil
}
