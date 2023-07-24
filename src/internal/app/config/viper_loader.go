package config

import (
	"log"
	"sync"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// viperLoader is a struct that implements Loader interface
type viperLoader struct{}

// ViperLoader is a function that returns a new instance of viperLoader
type ViperLoader interface {
	Unmarshal(i interface{})
}

var viperLoadOnce sync.Once

// NewViperLoader returns a new instance of ViperLoader. It is safe to call this more than once in a single go routine.
//
// Returns:
//
//	A new instance of ViperLoader. This will be an uninitialized instance of ViperLoader and will panic on error
func NewViperLoader() ViperLoader {
	viperLoadOnce.Do(func() {
		viper.SetConfigFile(configPath)
		// ReadInConfig reads in the config file and logs the error if it fails.
		if err := viper.ReadInConfig(); err != nil {
			log.Panic(errors.Wrap(err, "failed to read config"))
		}
	})

	return &viperLoader{}
}

// Unmarshal unmarshals viper into an interface. Panics on error. Use this to unmarshal config from a file or config.
//
// Args:
//
//	v: the loader to unmarshal into
//	i: the interface to unmarshal into. Must be a pointer
func (v viperLoader) Unmarshal(i interface{}) {
	// Unmarshal config to a config object
	if err := viper.Unmarshal(&i); err != nil {
		log.Panic(errors.Wrap(err, "failed to marshal config"))
	}
}
