package main

import (
	"io/ioutil"

	"github.com/pgavlin/femto"
)

type Buffer = femto.Buffer

func NewBufferFromString(text string, path string) (*Buffer, error) {
	return femto.NewBufferFromString(text, path), nil
}

func NewBufferFromScratch() (*Buffer, error) {
	return NewBufferFromString("", "*scratch*")
}

func NewBufferFromFile(path string) (*Buffer, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return NewBufferFromString(string(content), path)
}
