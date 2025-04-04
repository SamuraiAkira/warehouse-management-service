package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	HTTP     HTTPConfig
	Postgres PostgresConfig
}

type HTTPConfig struct {
	Host            string        `yaml:"host"`
	Port            string        `yaml:"port"`
	ReadTimeout     time.Duration `yaml:"read_timeout"`
	WriteTimeout    time.Duration `yaml:"write_timeout"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
}

type PostgresConfig struct {
	Host           string `yaml:"host"`
	Port           string `yaml:"port"`
	Username       string `yaml:"username"`
	Password       string `yaml:"password"`
	DBName         string `yaml:"db_name"`
	SSLMode        string `yaml:"ssl_mode"`
	MaxConnections int    `yaml:"max_connections"`
}

func Load() (*Config, error) {
	cfg := &Config{
		HTTP: HTTPConfig{
			Host:            getEnv("HTTP_HOST", "0.0.0.0"),
			Port:            getEnv("HTTP_PORT", "8080"),
			ReadTimeout:     10 * time.Second,
			WriteTimeout:    10 * time.Second,
			ShutdownTimeout: 5 * time.Second,
		},
		Postgres: PostgresConfig{
			Host:           getEnv("DB_HOST", "postgres"),
			Port:           getEnv("DB_PORT", "5432"),
			Username:       getEnv("DB_USER", "postgres"),
			Password:       getEnv("DB_PASSWORD", "postgres"),
			DBName:         getEnv("DB_NAME", "warehouse"),
			SSLMode:        getEnv("DB_SSLMODE", "disable"),
			MaxConnections: getEnvAsInt("DB_MAX_CONN", 10),
		},
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
