package main

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Env      string `yaml:"env"`
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`

	API struct {
		Key string `yaml:"key"`
	} `yaml:"api"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
