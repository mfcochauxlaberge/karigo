package util

func ReadConfig() (*Config, error) {
	return &Config{
		Port: 6280,
	}, nil
}
