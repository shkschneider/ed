package main

import (
	tv "github.com/rivo/tview"
)

var (
	app              *tv.Application
	ui               *tv.Pages
	windows          = make(map[string]*Window)
	currentPage      string
	minibufferActive bool
)

func SetMinibufferActive(active bool) {
	minibufferActive = active
}

func IsMinibufferActive() bool {
	return minibufferActive
}

func SetApp(a *tv.Application) {
	app = a
}

func GetApp() *tv.Application {
	return app
}

func SetUI(u *tv.Pages) {
	ui = u
}

func GetUI() *tv.Pages {
	return ui
}

func SetCurrentPage(name string) {
	currentPage = name
}

func GetCurrentPage() string {
	return currentPage
}

func AddWindow(name string, win *Window) {
	windows[name] = win
}

func GetCurrentWindow() *Window {
	return windows[currentPage]
}

func GetWindow(name string) *Window {
	return windows[name]
}

func RemoveWindow(name string) {
	delete(windows, name)
}

func WindowCount() int {
	return len(windows)
}
