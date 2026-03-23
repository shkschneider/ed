package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	tv "github.com/rivo/tview"
)

func Explorer(path string, cb func(int, string, string, rune)) (*tv.List, error) {
	list := tv.NewList()
	list.SetHighlightFullLine(false)
	list.ShowSecondaryText(false)
	list.SetWrapAround(true)
	if files, err := find(path) ; err != nil {
		return nil, err
	} else {
		for i, file := range files[1:] {
			list.AddItem(strings.Replace(file, path + string(os.PathSeparator), "", 1), file, rune('a'+i), nil)
		}
	}
	list.SetSelectedFunc(cb)
	return list, nil
}

func find(path string) ([]string, error) {
	var files []string = make([]string, 1)
	err := filepath.WalkDir(path, func(path string, dir fs.DirEntry, err error) error {
		if !dir.IsDir() && dir.Type().IsRegular() {
		    log.Debug(fmt.Sprintf("[%s] %s\n", dir.Name(), path))
			files = append(files, path)
		}
	    return err
	})
	return files, err
}

func Open(pages *tv.Pages, paths []string) error {
	switch len(paths) {
	case 1:
		path, err := os.Getwd()
		if err != nil { log.Fatal("os.Getwd()", err) }
		return openDirectory(pages, path)
		// err = openDirectory(app.UI, path)
		// if err != nil { log.Fatal(fmt.Sprintf("openDirectory(%s)", path), err) }
	default:
		// TODO for := range paths
		path, err := filepath.Abs(paths[1])
		if err != nil { log.Fatal(fmt.Sprintf("filepath.Abs(%s)", os.Args[1]), err) }
		info, err := os.Stat(path)
		if err != nil { log.Fatal(fmt.Sprintf("os.Stat(%s)", path), err) }
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
	file := filepath.Base(buffer.Path)
	pages.AddAndSwitchToPage(file, window.root, true)
	return nil
}

func openDirectory(pages *tv.Pages, path string) error {
	list, err := Explorer(path, func(index int, primary string, secondary string, shortcut rune) {
		log.Debug(fmt.Sprintf("openDirectory %d %s %s", index, primary, secondary))
		if buffer, err := NewBufferFromFile(secondary) ; err != nil {
			log.Fatal(fmt.Sprintf("NewBufferFromFile(%s)", secondary), err)
		} else {
			window := NewWindow(buffer)
			window.SetKeyBindings(NewKeyBindings())
			file := filepath.Base(buffer.Path)
			pages.AddAndSwitchToPage(file, window.root, true)
		}
	})
	if err != nil { return err }
	pages.AddAndSwitchToPage("explorer", list, true)
	return nil
}
