package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type App struct {
	Git  string `yaml:"git,omitempty"`
	Name string
}

type Config struct {
	Token string
	Apps  []App
}

func load() Config {
	config := Config{}

	file, err := ioutil.ReadFile("../conf/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	yamlErr := yaml.Unmarshal(file, &config)
	if yamlErr != nil {
		log.Fatal(err)
	}

	return config
}
