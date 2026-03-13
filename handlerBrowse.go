package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/jcuello/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	maxPosts := 2
	if len(cmd.args) > 0 {
		parsedLimit, err := strconv.Atoi(cmd.args[0])
		if err == nil {
			maxPosts = parsedLimit
		}
	}

	posts, err := s.db.GetPosts(context.Background(), database.GetPostsParams{
		UserID: user.ID,
		Limit:  int32(maxPosts),
	})

	if err != nil {
		return err
	}

	for _, post := range posts {
		jsonData, _ := json.MarshalIndent(post, "", "  ")
		fmt.Println(string(jsonData))
	}
	return nil
}
