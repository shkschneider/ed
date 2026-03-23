package main

import (
	"fmt"
	"os"
	// "io/ioutil"
	// "path/filepath"

	semver "github.com/Masterminds/semver/v3"
	tc "github.com/gdamore/tcell/v2"
	tv "github.com/rivo/tview"
)

var (
	ID         = "github.com/shkschneider/ed"
	NAME       = "ed"
	VERSION, _ = semver.NewVersion("0.1.1")
)

func main() {
	// logger
	f, err := os.Create(fmt.Sprintf("/tmp/%s.log", NAME))
	if err == nil {
		NewLog(f, LogLevelDebug, true)
		defer f.Close()
	} else {
		NewLog(os.Stderr, LogLevelDebug, true)
	}
	log.Info(fmt.Sprintf("%s %s %s", ID, NAME, VERSION.String()))
	// config
	conf, err := NewConfig(fmt.Sprintf(NAME, ".kdl"))
	if err != nil {
		log.Error("NewConfig()", err)
	}
	log.Debug(conf)
	// application
	app := tv.NewApplication()
	app.SetTitle(NAME)
	SetApp(app)
	// ui
	ui := tv.NewPages()
	SetUI(ui)
	// ...
	Open(ui, os.Args)

	// win.root.HandleEvent(tc.Event)
	// win.root.ExecuteAction(actions []func(*View) bool)
	// []string{"Delete", "Insert", "Backspace", "Cut", "Play", "Paste", "Move", "Add", "DuplicateLine", "Macro"}

	// keyboard
	// bindings := NewBindings()
	// bindings.Add("Ctrl+Q", func(event *tc.EventKey) *tc.EventKey {
	// 	return event
	// })
	// app.SetInputCapture(bindings.Capture(func (event *tc.EventKey) {
	// 	win.Update(buffer)
	// }))
	app.SetInputCapture(func(event *tc.EventKey) *tc.EventKey {
		if event.Key() == tc.KeyCtrlSpace {
			win := GetCurrentWindow()
			if win != nil {
				win.FocusMinibuffer()
				SetMinibufferActive(true)
			}
			return nil
		}
		if IsMinibufferActive() && event.Key() == tc.KeyEsc {
			win := GetCurrentWindow()
			if win != nil {
				win.ClearMinibuffer()
				app.SetFocus(win.Content)
			}
			SetMinibufferActive(false)
			return nil
		}
		log.Debug(fmt.Sprintf("Name: %s", event.Name()))
		return event
	})
	// ioutil.WriteFile(os.Args[1], []byte(buffer.String()), 0644)

	// mouse
	app.EnableMouse(true)
	app.EnablePaste(true)
	// mouseLeftClick := false
	// mouseLeftDoubleClick := false
	// mouseMiddleClick := false
	// mouseMiddleDoubleClick := false
	// mouseRightClick := false
	// mouseRightDoubleClick := false
	// mouseScrollUp := false
	// mouseScrollDown := false
	app.SetMouseCapture(func(event *tc.EventMouse, action tv.MouseAction) (*tc.EventMouse, tv.MouseAction) {
		x, y := event.Position()
		switch event.Buttons() {
		case tc.ButtonPrimary:
			log.Debug(fmt.Sprintf("Primary %d:%d", x, y))
		case tc.ButtonSecondary:
			log.Debug(fmt.Sprintf("Secondary %d:%d", x, y))
		}
		return event, action
	})
	// app.SetMouseCapture(func(event *tc.EventMouse, action tv.MouseAction) (*tc.EventMouse, tv.MouseAction) {
	// 	x, y := event.Position()
	// 	switch event.Buttons() {
	// 	case tc.ButtonPrimary:
	// 		log.Debug(fmt.Sprintf("Primary %d:%d", x, y))
	// 	case tc.ButtonSecondary:
	// 		log.Debug(fmt.Sprintf("Secondary %d:%d", x, y))
	// 	}
	// 	switch action {
	// 	case tv.MouseLeftDown:
	// 		buffer.Cursor.GotoLoc(femto.Loc { X: x, Y: y })
	// 		log.Debug("MouseLeftDown -> SetSelectionStart")
	// 		if buffer.Cursor.HasSelection() {
	// 			buffer.Cursor.ResetSelection()
	// 		}
	// 		buffer.Cursor.SetSelectionStart(femto.Loc { X: x, Y: y })
	// 	case tv.MouseLeftUp:
	// 		log.Debug("MouseLeftUp -> SetSelectionEnd,CopySelection")
	// 		buffer.Cursor.SetSelectionEnd(femto.Loc { X: x, Y: y })
	// 		buffer.Cursor.CopySelection("clipboard")
	// 	case tv.MouseLeftClick:
	// 		log.Debug("MouseLeftClick -> GotoLoc")
	// 		buffer.Cursor.GotoLoc(femto.Loc { X: x, Y: y })
	// 	}
	// 	return action, event
	// })

	// run
	app.SetRoot(ui, true)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
