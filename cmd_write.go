package main

import (
	"fmt"
	"io/ioutil"
)

func init() {
	Register("write", cmdWrite, "w")
}

func cmdWrite(args []string) (string, error) {
	win := GetCurrentWindow()
	if win == nil {
		return "", fmt.Errorf("no current buffer")
	}
	if win.buffer == nil {
		return "", fmt.Errorf("no buffer in current window")
	}
	path := win.buffer.Path
	if len(args) > 0 {
		path = args[0]
	}
	content := win.buffer.String()
	err := ioutil.WriteFile(path, []byte(content), 0644)
	if err != nil {
		return "", err
	}
	win.buffer.Path = path
	win.Update()
	return fmt.Sprintf("Saved: %s", path), nil
}
