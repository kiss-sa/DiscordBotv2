package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Token    string `json:"token"`
	FilePath string `json:"filepath"`
}

var (
	Config *Configuration
)

func ReadConfig() (*Configuration, error) {
	configFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		return nil, err
	}
	var config Configuration
	if err := yaml.Unmarshal(configFile, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
