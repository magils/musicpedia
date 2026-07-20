package main

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Functions non-binding to the config struct

func getDefaults() map[string]string {
	var DEFAULT_VALUES = map[string]string{
		"DB_HOST": "localhost",
		"DB_PORT": "3306",
		"DB_NAME": "musicpedia",
	}

	return DEFAULT_VALUES
}

func getEnvValueOrDefault(key string) string {
	defaults := getDefaults()

	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaults[key]
}

func getEnvValueOrDefaultAsInt(key string) (int, error) {
	if value := os.Getenv(key); value != "" {
		intValue, err := strconv.Atoi(value)
		return intValue, err
	}

	defaults := getDefaults()
	return strconv.Atoi(defaults[key])
}

// Config Struct

type Config struct {
	Host       string
	Port       int
	Env        string
	DbHost     string
	DbPort     int
	DbName     string
	DbUser     string
	DbPassword string
}

func (c *Config) LoadDotEnvConfig() error {
	err := godotenv.Load()

	if err != nil {
		return err
	}

	c.DbUser = os.Getenv("DB_USER")
	c.DbPassword = os.Getenv("DB_PASSWORD")
	c.Env = os.Getenv("ENV")
	c.Host = getEnvValueOrDefault("HOST")
	c.DbHost = getEnvValueOrDefault("DB_HOST")
	c.DbName = getEnvValueOrDefault("DB_NAME")

	dbPortValue, dbPortErr := getEnvValueOrDefaultAsInt("DB_PORT")
	portValue, portErr := getEnvValueOrDefaultAsInt("PORT")

	if dbPortErr != nil {
		c.DbPort = dbPortValue
	}

	if portErr != nil {
		c.Port = portValue
	}

	return nil
}
