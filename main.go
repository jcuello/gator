package main

import (
	"fmt"
	"os"

	"github.com/jcuello/gator/internal/config"
)

func main() {
	cfg, err := config.Read()

	if err != nil {
		printErrorAndExit(err)
	}

	err = cfg.SetUser("jose")
	if err != nil {
		printErrorAndExit(err)
	}

	cfg, err = config.Read()
	if err != nil {
		printErrorAndExit(err)
	}

	fmt.Printf("%+v\n", cfg)
}

func printErrorAndExit(err error) {
	fmt.Printf("%v\n", err)
	os.Exit(1)
}
