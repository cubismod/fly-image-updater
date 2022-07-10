package main

import (
	"github.com/cubismod/fly-image-updater/src/config"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	config := config.Load()

	// parse through the configured repositories
	for _, s := range config.Apps {
		logrus.Info(s)
	}

}
