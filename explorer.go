package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	tc "github.com/gdamore/tcell/v2"
	tv "github.com/rivo/tview"
	"golang.org/x/term"
)

type FuzzyExplorer struct {
	root     tv.Primitive
	input    *tv.InputField
	list     *tv.List
	status   *tv.TextView
	fullList []string
	basePath string
	callback func(string)
}

func NewFuzzyExplorer(path string, callback func(string)) *FuzzyExplorer {
	fe := &FuzzyExplorer{
		fullList: make([]string, 0),
		basePath: path,
		callback: callback,
	}

	fe.input = tv.NewInputField()
	fe.input.SetPlaceholder("Filter files...")
	fe.input.SetLabel("")
	fe.input.SetFieldTextColor(tc.GetColor("white"))
	fe.input.SetPlaceholderTextColor(tc.GetColor("gray"))
	fe.input.SetFieldBackgroundColor(tc.ColorBlack)

	fe.list = tv.NewList()
	fe.list.SetHighlightFullLine(false)
	fe.list.ShowSecondaryText(false)
	fe.list.SetWrapAround(true)

	fe.status = tv.NewTextView()
	fe.status.SetText("No file selected")
	fe.status.SetTextColor(tc.GetColor("black"))
	fe.status.SetBackgroundColor(tc.GetColor("white"))

	fe.root = tv.NewGrid().SetRows(0, 1, 1).SetBorders(false).
		AddItem(fe.list, 0, 0, 1, 1, 0, 0, false).
		AddItem(fe.status, 1, 0, 1, 1, 1, 0, false).
		AddItem(fe.input, 2, 0, 1, 1, 1, 0, true)

	fe.loadFiles(path)

	fe.input.SetChangedFunc(func(text string) {
		fe.filter(text)
	})

	fe.input.SetDoneFunc(func(key tc.Key) {
		if key == tc.KeyEnter {
			index := fe.list.GetCurrentItem()
			if index >= 0 && index < fe.list.GetItemCount() {
				mainText, _ := fe.list.GetItemText(index)
				fe.callback(mainText)
			}
		}
	})

	fe.input.SetInputCapture(func(event *tc.EventKey) *tc.EventKey {
		if event.Key() == tc.KeyUp {
			current := fe.list.GetCurrentItem()
			fe.list.SetCurrentItem(current - 1)
			fe.updateStatus()
			return nil
		}
		if event.Key() == tc.KeyDown {
			current := fe.list.GetCurrentItem()
			fe.list.SetCurrentItem(current + 1)
			fe.updateStatus()
			return nil
		}
		return event
	})

	return fe
}

func (fe *FuzzyExplorer) loadFiles(path string) {
	filepath.WalkDir(path, func(path string, dir fs.DirEntry, err error) error {
		name := dir.Name()
		if !strings.HasPrefix(name, ".") {
			if !dir.IsDir() && dir.Type().IsRegular() {
				log.Debug(fmt.Sprintf("[%s] %s\n", name, path))
				fe.fullList = append(fe.fullList, path)
			}
		} else {
			if dir.IsDir() {
				return filepath.SkipDir
			}
		}
		return err
	})
	fe.filter("")
}

func (fe *FuzzyExplorer) getScreenWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil || width < 10 {
		return 80
	}
	return width - 1
}

func (fe *FuzzyExplorer) truncatePath(path string, maxWidth int) string {
	if len(path) <= maxWidth {
		return path
	}
	filename := filepath.Base(path)
	if len(filename) >= maxWidth {
		if len(filename) > maxWidth {
			filename = filename[len(filename)-maxWidth:]
		}
		return "..." + filename
	}
	prefixLen := maxWidth - len(filename) - 4
	if prefixLen > 0 {
		return ".../" + path[len(path)-prefixLen:] + "/" + filename
	}
	return "..." + filename
}

func (fe *FuzzyExplorer) filter(query string) {
	fe.list.Clear()
	query = strings.ToLower(query)
	screenWidth := fe.getScreenWidth()
	for _, path := range fe.fullList {
		name := strings.ToLower(filepath.Base(path))
		if fe.fuzzyMatch(name, query) {
			display := strings.TrimPrefix(path, fe.basePath+string(os.PathSeparator))
			display = fe.truncatePath(display, screenWidth)
			fe.list.AddItem(display, path, 0, nil)
		}
	}
	if fe.list.GetItemCount() > 0 {
		fe.list.SetCurrentItem(0)
	}
	fe.updateStatus()
}

func (fe *FuzzyExplorer) fuzzyMatch(text, query string) bool {
	if query == "" {
		return true
	}
	for _, r := range query {
		idx := strings.Index(text, string(r))
		if idx == -1 {
			return false
		}
		text = text[idx+1:]
	}
	return true
}

func (fe *FuzzyExplorer) updateStatus() {
	index := fe.list.GetCurrentItem()
	if index >= 0 && index < fe.list.GetItemCount() {
		_, path := fe.list.GetItemText(index)
		info, err := os.Stat(path)
		if err == nil {
			filename := filepath.Base(path)
			ext := filepath.Ext(filename)
			sizeStr := formatSize(info.Size())
			fe.status.SetText(fmt.Sprintf("[%s] %s (%s)", ext, filename, sizeStr))
		} else {
			fe.status.SetText(filepath.Base(path))
		}
	} else {
		fe.status.SetText("No file selected")
	}
}

func formatSize(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func (fe *FuzzyExplorer) Root() tv.Primitive {
	return fe.root
}

func (fe *FuzzyExplorer) Focus() {
	app.SetFocus(fe.input)
}
