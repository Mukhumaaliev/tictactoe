package datasource

import (
	"fmt"
	"os"
)

type Config struct {
	dbname   string
	user     string
	password string
	host     string
	port     string
	sslmode  string
}

func NewConfig() *Config {
	return &Config{
		dbname:   getEnv("DB_NAME", "tictactoe"),
		user:     getEnv("DB_USER", "postgres"),
		password: getEnv("DB_PASSWORD", "password"),
		host:     getEnv("DB_HOST", "localhost"),
		port:     getEnv("DB_PORT", "5432"),
		sslmode:  getEnv("DB_SSL_MODE", "disable"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", c.user, c.password, c.host, c.port, c.dbname, c.sslmode)
}
