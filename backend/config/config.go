package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	ClientId     string
	ClientSecret string
	SessionKey   string
	JWTKey       string
	DBUser       string
	DBPassword   string
	DBAddress    string
	DBName       string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not read .env file")
	}

	return &Config{
		Port:         getEnv("PORT"),
		ClientId:     getEnv("CLIENT_ID"),
		ClientSecret: getEnv("CLIENT_SECRET"),
		SessionKey:   getEnv("SESSION_KEY"),
		JWTKey:       getEnv("JWT_KEY"),
		DBUser:       getEnv("DB_USER"),
		DBPassword:   getEnv("DB_PASSWORD"),
		DBAddress:    fmt.Sprintf("%s:%s", getEnv("DB_HOST"), getEnv("DB_PORT")),
		DBName:       getEnv("DB_NAME"),
	}
}

func getEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("key %s not set in .env\n", key)
	}

	return val
}
