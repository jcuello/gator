package main

import (
	"context"
	"fmt"

	"github.com/jcuello/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	feedsFollowed, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	fmt.Printf("Feeds followed by %v:\n", user.Name)
	for _, feed := range feedsFollowed {
		fmt.Println(feed.FeedName)
	}
	return nil
}
