package config

import (
	"log"
	"os"
)

type Config struct {
	ClientId     string
	ClientSecret string
	Port         string
	DBUser       string
	DBPassword   string
	DBName       string
	SessionKey   string
	ClientUrl    string
}

func NewConfig() Config {
	return Config{
		ClientId:     getEnv("CLIENT_ID"),
		ClientSecret: getEnv("CLIENT_SECRET"),
		Port:         getEnv("PORT"),
		DBUser:       getEnv("DB_USER"),
		DBPassword:   getEnv("DB_PASSWORD"),
		DBName:       getEnv("DB_NAME"),
		SessionKey:   getEnv("SESSION_KEY"),
		ClientUrl:    getEnv("CLIENT_URL"),
	}
}

func getEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatal("variable", key, "not set")
	}
	return v
}
