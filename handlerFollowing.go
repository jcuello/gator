package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	username := s.cfg.CurrentUsername
	feedsFollowed, err := s.db.GetFeedFollowsForUser(context.Background(), username)
	if err != nil {
		return err
	}

	fmt.Printf("Feeds followed by %v:\n", username)
	for _, feed := range feedsFollowed {
		fmt.Println(feed.FeedName)
	}
	return nil
}
