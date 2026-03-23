package main

import (
	"fmt"
	"path/filepath"

	tc "github.com/gdamore/tcell/v2"
	"github.com/pgavlin/femto"
	"github.com/pgavlin/femto/runtime"
	tv "github.com/rivo/tview"
)

type Window struct {
	root			tv.Primitive
	Header			*tv.TextView
	Content			*femto.View
	buffer			*femto.Buffer
	FooterStatus	*tv.TextView
	FooterCommand	*tv.TextView
}

func newHeaderView() *tv.TextView {
	view := tv.NewTextView()
	view.SetText("header")
	view.SetTextColor(tc.GetColor("white"))
	view.SetBackgroundColor(tc.GetColor("black"))
	return view
}

func newContentView(buffer *femto.Buffer) *femto.View {
	view := femto.NewView(buffer)
	view.SetRuntimeFiles(runtime.Files)
	return view
}

func newFooterStatusView() *tv.TextView {
	view := tv.NewTextView()
	view.SetText("status")
	view.SetTextColor(tc.GetColor("black"))
	view.SetBackgroundColor(tc.GetColor("white"))
	return view
}

func newFooterCommandView() *tv.TextView {
	view := tv.NewTextView()
	view.SetText("command")
	view.SetTextColor(tc.GetColor("white"))
	view.SetBackgroundColor(tc.GetColor("black"))
	return view
	// view := tv.NewInputField()
	// view.SetLabel("command:")
	// view.SetPlaceholder("...")
	// view.SetFieldTextColor(tc.GetColor("white"))
	// view.SetFieldBackgroundColor(tc.GetColor("black"))
	// //view.SetFocus(false)
	// view.SetDisabled(false)
	// view.SetDoneFunc(func(key tc.Key) {
	// 	log.Debug(key)
	// })
	// return view
}

func NewWindow(buffer *Buffer) *Window {
	header := newHeaderView()
	content := newContentView(buffer)
	footer_status := newFooterStatusView()
	footer_command := newFooterCommandView()
	content.SetRuntimeFiles(runtime.Files)
	if cs := runtime.Files.FindFile(femto.RTColorscheme, "monokai"); cs != nil {
		if data, err := cs.Data(); err == nil {
			content.SetColorscheme(femto.ParseColorscheme(string(data)))
		}
	}
	return &Window {
		root: tv.NewGrid().SetRows(1, 0, 1, 1).SetBorders(false).
			//     (...,            row, column, rowSpan, colSpan, minGridHeight, minGridWidth int, focus bool)
			AddItem(header,         0,   0,      1,       1,       1,             0,                false).
			AddItem(content,        1,   0,      1,       1,       0,             0,                true).
			AddItem(footer_status,  2,   0,      1,       1,       1,             0,                false).
			AddItem(footer_command, 3,   0,      1,       1,       1,             0,                false),
		Header: header,
		Content: content,
		buffer: buffer,
		FooterStatus: footer_status,
		FooterCommand: footer_command,
	}
}

func (this *Window) SetKeyBindings(bindings femto.KeyBindings) {
	this.Content.SetKeybindings(bindings)
}

func (this *Window) Update() {
	this.Header.SetText(fmt.Sprintf("[%s]", filepath.Base(this.buffer.Path)))
	position := fmt.Sprintf("%d:%d/%d", this.buffer.Cursor.Loc.X, this.buffer.Cursor.Loc.Y, this.buffer.NumLines)
	this.FooterStatus.SetText(fmt.Sprintf("[%s] %s (%s)", this.buffer.FileType(), this.buffer.Path, position))
}
