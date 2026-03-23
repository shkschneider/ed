package main

import (
	"fmt"
	"strings"
)

func init() {
	Register("help", cmdHelp, "h", "?")
}

func cmdHelp(args []string) (string, error) {
	if len(args) > 0 {
		if _, ok := Commands[args[0]]; ok {
			return fmt.Sprintf("Command: %s", args[0]), nil
		}
		return "", fmt.Errorf("unknown command: %s", args[0])
	}
	list := CommandsList()
	return fmt.Sprintf("Available commands: %s", strings.Join(list, ", ")), nil
}
