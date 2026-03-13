package main

import (
	"errors"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("agg requires a duration string e.g. '1s, 1m, 1h'")
	}

	duration, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}

	ticker := time.NewTicker(duration)
	fmt.Printf("Collecting feeds every %v\n", duration)
	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			fmt.Println(err)
		}
	}
}
