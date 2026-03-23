package main

import (
	"fmt"
	"os"
	"path/filepath"

	tv "github.com/rivo/tview"
)

func Open(pages *tv.Pages, paths []string) error {
	switch len(paths) {
	case 1:
		path, err := os.Getwd()
		if err != nil {
			log.Fatal("os.Getwd()", err)
		}
		return openDirectory(pages, path)
		// err = openDirectory(app.UI, path)
		// if err != nil { log.Fatal(fmt.Sprintf("openDirectory(%s)", path), err) }
	default:
		// TODO for := range paths
		path, err := filepath.Abs(paths[1])
		if err != nil {
			log.Fatal(fmt.Sprintf("filepath.Abs(%s)", os.Args[1]), err)
		}
		info, err := os.Stat(path)
		if err != nil {
			log.Fatal(fmt.Sprintf("os.Stat(%s)", path), err)
		}
		if info.IsDir() {
			return openDirectory(pages, path)
		} else {
			return openFile(pages, path)
		}
		return nil
	}
}

// func open(pages *tv.Pages, path string) error {
// 	if len(path) == 0 {
// 		return openFile(pages, "/tmp/ed")
// 	}
// 	path, err := filepath.Abs(path)
// 	if err != nil {
// 		log.Fatal(fmt.Sprintf("filepath.Abs(%s)", os.Args[1]), err)
// 		return err
// 	}
// 	info, err := os.Stat(path)
// 	if err != nil {
// 		log.Fatal(fmt.Sprintf("os.Stat(%s)", path), err)
// 		return err
// 	}
// 	if info.IsDir() {
// 		return openDirectory(pages, path)
// 	} else {
// 		return openFile(pages, path)
// 	}
// 	return nil
// }

func openFile(pages *tv.Pages, path string) error {
	buffer, err := NewBufferFromFile(path)
	if err != nil {
		log.Fatal("NewBufferFromFile", err)
		return err
	}
	window := NewWindow(buffer)
	window.SetKeyBindings(NewKeyBindings())
	window.Update()
	file := filepath.Base(buffer.Path)
	pages.AddAndSwitchToPage(file, window.root, true)
	AddWindow(file, window)
	SetCurrentPage(file)
	return nil
}

func openDirectory(pages *tv.Pages, path string) error {
	explorer := NewFuzzyExplorer(path, func(selected string) {
		log.Debug(fmt.Sprintf("openDirectory %s", selected))
		if buffer, err := NewBufferFromFile(selected); err != nil {
			log.Fatal(fmt.Sprintf("NewBufferFromFile(%s)", selected), err)
		} else {
			window := NewWindow(buffer)
			window.SetKeyBindings(NewKeyBindings())
			file := filepath.Base(buffer.Path)
			pages.AddAndSwitchToPage(file, window.root, true)
			AddWindow(file, window)
			SetCurrentPage(file)
		}
	})
	pages.AddAndSwitchToPage("explorer", explorer.Root(), true)
	explorer.Focus()
	return nil
}
