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
		path = "./karigo.yml"
	}

	c.SetConfigFile(path)

	// Read the config
	err := c.ReadInConfig()
	if err != nil {
		return karigo.Config{}, err
	}

	// Defaults
	err = c.Unmarshal(&config)
	if err != nil {
		return karigo.Config{}, err
	}

	return config, nil
}
