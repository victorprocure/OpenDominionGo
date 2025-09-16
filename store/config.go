package store

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	SecureFlag bool
	DBUser string
	DBPassword string
	DBAddress string
	DBName string
}

var Envs = initConfig()

func initConfig() Config {
	err := godotenv.Load(".env.local", ".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	return Config{
		Port: getEnv("APP_PORT", "8080"),
		SecureFlag: getEnv("SECURE_FLAG", "true") == "true",
		DBUser: getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBAddress: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "5432")),
		DBName: getEnv("DB_NAME", "postgres"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}