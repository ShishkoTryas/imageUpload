package config

import (
	"os"
	"strconv"
)

type config struct {
	DB     Postgres
	Phrase Phrase
}

type Postgres struct {
	Host     string
	Port     int
	Username string
	DBName   string
	SSLMode  string
	Password string
}

type Phrase struct {
	Salt   string
	Secret string
}

func New() *config {
	return &config{
		DB: Postgres{
			Host:     getEnv("DB_HOST", ""),
			Port:     getEnvAsInt("DB_PORT", 0),
			Username: getEnv("DB_USERNAME", ""),
			DBName:   getEnv("DB_NAME", ""),
			SSLMode:  getEnv("DB_SSLMODE", ""),
			Password: getEnv("DB_PASSWORD", ""),
		},
		Phrase: Phrase{
			Salt:   getEnv("SALT", ""),
			Secret: getEnv("SECRET", ""),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}
