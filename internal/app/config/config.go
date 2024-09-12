package config

import "os"

type Config struct {
	HTTPAddr       string
	DSN            string
	MigrationsPath string
}

func Read() Config {
	var config Config
	httpAddr, exists := os.LookupEnv("SERVER_ADDRESS")
	if exists {
		config.HTTPAddr = httpAddr
	}

	dsn, exists := os.LookupEnv("POSTGRES_CONN")
	if exists {
		config.DSN = dsn
	}

	return config
}
