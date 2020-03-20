package util

import (
	"github.com/mfcochauxlaberge/karigo"

	"github.com/spf13/viper"
)

func ReadConfig(path string) (karigo.Config, error) {
	config := karigo.Config{}

	c := viper.New()

	// If path is not specified, try
	// the current directory.
	if path == "" {
		path = "karigo.yml"
	}

	c.SetConfigFile(path)

	// Read the config
	_ = c.ReadInConfig()
	// Error ignored for now
	// because a configuration
	// file is not necessary.

	// Defaults
	err := c.Unmarshal(&config)
	if err != nil {
		return karigo.Config{}, err
	}

	return config, nil
}
