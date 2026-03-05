package main

import (
	"context"
	"fmt"
	"os"
)

func handlerReset(s *state, _ command) error {
	err := s.db.DeleteAllUsers(context.Background())

	if err != nil {
		fmt.Printf("Unable to delete all users: %v\n", err.Error())
		os.Exit(1)
	}

	fmt.Println("All users were deleted.")
	return nil
}
