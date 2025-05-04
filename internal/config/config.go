package config

import (
	"os"
)

type MongoConfig struct {
	URL      string
	PASS     string
	USER     string
	DATABASE string
}

type Config struct {
	Env       string // e.g. "dev" or "prod"
	Port      string
	DNSPrefix string
	Mongo     MongoConfig
	LogFile   string
	SUPERKEY  string
}

var AppConfig Config

func Load() {
	AppConfig = Config{
		Env:       getEnv("APP_ENV", "dev"),
		Port:      getEnv("PORT", "8080"),
		LogFile:   getEnv("LOG_FILE", "logs/email-api.log"),
		DNSPrefix: getEnv("DNS_PREFIX", "email-verification"),
		SUPERKEY:  getEnv("SUPER_KEY", "Admin123?"),
		Mongo:     loadMongoConfig(),
	}
}

func loadMongoConfig() MongoConfig {
	return MongoConfig{
		URL:      getEnv("MONGO_URL", "localhost:27017"),
		USER:     getEnv("MONGO_USER", ""),
		PASS:     getEnv("MONGO_PASS", ""),
		DATABASE: getEnv("MONGO_DB", "email_service"),
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
