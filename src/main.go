package main

import "github.com/cubismod/fly-image-updater/src/config"

func main() {
	config := config.Load()

	// parse through the configured repositories
	for _, s := range config.Apps {

	}

}
