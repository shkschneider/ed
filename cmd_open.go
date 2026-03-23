package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func init() {
	Register("open", cmdOpen, "o")
}

func cmdOpen(args []string) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("usage: :open <path>")
	}
	path := args[0]
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	info, err := os.Stat(absPath)
	if err != nil {
		return "", err
	}
	if info.IsDir() {
		openDirectory(ui, absPath)
		return fmt.Sprintf("Opened directory: %s", absPath), nil
	}
	openFile(ui, absPath)
	return fmt.Sprintf("Opened file: %s", absPath), nil
}
