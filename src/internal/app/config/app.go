package config

import (
	"log"
	"os"
)

var configPath = "./config.json"

// isFile checks if the file exists and is a file. It doesn t check if the file is a directory.
//
// Args:
//
//	fp: path to the file to check. This can be a file or a directory.
//
// Returns:
//
//	true if the file exists and is a file false otherwise. Note that this is different from os. IsNotExist
func isFile(fp string) bool {
	f, e := os.Stat(fp)
	// Return a boolean if a field has been set.
	if e != nil {
		return false
	}

	return !f.IsDir()
}

// SetConfigPath sets the path to the config file. It panics if the path does not exist. This is useful for debugging and to ensure that the config file is loaded before any tests are run
//
// Args:
//
//	path: the path to the
func SetConfigPath(path string) {
	// Check if the config file exists
	if !isFile(path) {
		log.Panic("config filepath does not exist")
	}

	configPath = path
}
