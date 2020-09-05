package config

import "os"

// Config ...
type Config struct {
	Host string
}

// New ...
func New() *Config {
	return &Config{
		Host: mustEnv("HOST"),
	}
}

func env(e string) string {
	return os.Getenv(e)
}

func mustEnv(e string) string {
	v := os.Getenv(e)
	if v == "" {
		panic("missing env " + e)
	}
	return v
}
