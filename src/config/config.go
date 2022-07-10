package config

import (
	"io/ioutil"
	"log"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type App struct {
	GitRepo  string `yaml:"git,omitempty"`
	Path string `yaml:"path,omitempty"`
	Name string
}

// Go representation of our YAML configuration file
type Config struct {
	FlyToken    string `yaml:"fly_token"`
	GotifyToken string `yaml:"gotify_token"`
	GotifyUrl   string `yaml:"gotify_url"`
	Apps        []App  `yaml:"apps"`
	StaggerSecs int    `yaml:"stagger_secs"`
}

func Load() Config {
	config := Config{}

	file, err := ioutil.ReadFile("./conf/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	yamlErr := yaml.Unmarshal(file, &config)
	if yamlErr != nil {
		log.Fatal(err)
	}

	logrus.Info("loaded YAML configuration")

	return config
}
