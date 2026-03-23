package main

import (
	"fmt"
)

func init() {
	Register("quit", cmdQuit, "q")
}

func cmdQuit(args []string) (string, error) {
	win := GetCurrentWindow()
	if win == nil {
		return "No windows to close", nil
	}
	current := GetCurrentPage()
	if current == "" {
		return "No current page", nil
	}
	ui.RemovePage(current)
	RemoveWindow(current)
	if len(windows) == 0 {
		app.Stop()
	}
	return fmt.Sprintf("Closed: %s", current), nil
}
