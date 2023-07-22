package config

import (
	"log"
	"os"
)

var configPath = "./config.json"

func isFile(fp string) bool {
	f, e := os.Stat(fp)
	if e != nil {
		return false
	}

	return !f.IsDir()
}

func SetConfigPath(path string) {
	if !isFile(path) {
		log.Panic("config filepath does not exist")
	}

	configPath = path
}
