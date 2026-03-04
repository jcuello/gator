package main

import (
	"fmt"

	c "github.com/jcuello/gator/internal/config"
)

type state struct {
	cfg *c.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	cmdFunc, ok := c.cmds[cmd.name]
	if !ok {
		return fmt.Errorf("unknown command: %v\n", cmd.name)
	}

	return cmdFunc(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}
