package main

import (
	"fmt"
	"errors"
	"io/ioutil"
	"path/filepath"

	"github.com/adrg/xdg" // TODO os.UserConfigDir() (string, error)

	kdl "github.com/sblinch/kdl-go"
)

type Config struct {
	Id			string `kdl:"id"`
	Name		string `kdl:"name"`
	Version		string `kdl:"version"`
	Bindings	struct {
		string	string
	} `kdl:"bindings"`
}

func NewConfig(path string) (*Config, error) {
	config := &Config {
		Id: ID,
		Name: NAME,
		Version: VERSION.String(),
	}
	switch filepath.Ext(path)[1:] {
	case "kdl":
		return newKdlConfig(path)
	default:
		return config, errors.New("No configuration")
	}
	return config, nil
}

func newKdlConfig(path string) (*Config, error) {
	if path == "" {
		path, _ = xdg.SearchConfigFile(fmt.Sprintf("%s/config.kdl", NAME))
	}
	data, err :=  ioutil.ReadFile(path) ; if err != nil {
		log.Error(err)
		return nil, err
	}
	config := &Config {}
	if err := kdl.Unmarshal(data, config) ; err != nil {
		log.Error(err)
		return nil, err
	} else {
		config.Id = ID
		config.Name = NAME
		config.Version = VERSION.String()
	}
    return config, nil
}
