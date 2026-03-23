package main

import (
	"strings"

	// tc "github.com/gdamore/tcell/v2"
	"github.com/pgavlin/femto"
)


func NewKeyBindings() femto.KeyBindings {
	return femto.DefaultKeyBindings
}

func NewKeyBinding(bindings femto.KeyBindings, key string, actions []string) femto.KeyBindings {
	return bindings.BindKey(key, strings.Join(actions, ","))
}

// type Bindings struct {
// 	configuration *bind.Configuration
// }
//
// func NewBindings() *Bindings {
// 	return &Bindings {
// 		configuration: bind.NewConfiguration(),
// 	}
// }
//
// func (this *Bindings) Add(binding string, cb func(event *tc.EventKey) *tc.EventKey) error {
// 	return this.configuration.Set(binding, cb)
// }
//
// func (this *Bindings) Capture(cb func(*tc.EventKey)) func(*tc.EventKey) *tc.EventKey {
// 	return func(event *tc.EventKey) *tc.EventKey {
// 		this.configuration.Capture(event)
// 		cb(event)
// 		return event
// 	}
// }
