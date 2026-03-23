package main

import (
	"fmt"
	"strings"
)

type Command func(args []string) (string, error)

var Commands = make(map[string]Command)

func Register(name string, cmd Command, aliases ...string) {
	Commands[name] = cmd
	for _, alias := range aliases {
		Commands[alias] = cmd
	}
}

func Run(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", fmt.Errorf("empty command")
	}

	var cmd string
	var args []string

	if input[0] == ':' {
		parts := strings.SplitN(input[1:], " ", 2)
		cmd = parts[0]
		if len(parts) > 1 {
			args = strings.Fields(parts[1])
		}
	} else {
		parts := strings.SplitN(input, " ", 2)
		cmd = parts[0]
		if len(parts) > 1 {
			args = strings.Fields(parts[1])
		}
	}

	if fn, ok := Commands[cmd]; ok {
		return fn(args)
	}

	return "", fmt.Errorf("unknown command: %s", cmd)
}

func CommandsList() []string {
	var list []string
	for name := range Commands {
		list = append(list, name)
	}
	return list
}
