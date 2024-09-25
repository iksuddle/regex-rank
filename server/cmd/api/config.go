package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type config struct {
	port         string
	clientId     string
	clientSecret string
	sessionKey   string
	jwtKey       string
	dbUser       string
	dbPassword   string
	dbAddress    string
	dbName       string
}

func newConfig() *config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not read .env file")
	}

	return &config{
		port:         getEnv("PORT"),
		clientId:     getEnv("CLIENT_ID"),
		clientSecret: getEnv("CLIENT_SECRET"),
		sessionKey:   getEnv("SESSION_KEY"),
		jwtKey:       getEnv("JWT_KEY"),
		dbUser:       getEnv("DB_USER"),
		dbPassword:   getEnv("DB_PASSWORD"),
		dbAddress:    fmt.Sprintf("%s:%s", getEnv("DB_HOST"), getEnv("DB_PORT")),
		dbName:       getEnv("DB_NAME"),
	}
}

func newMySQLConfig(config *config) *mysql.Config {
	mySQLConfig := mysql.NewConfig()

	mySQLConfig.User = config.dbUser
	mySQLConfig.Passwd = config.dbPassword
	mySQLConfig.Net = "tcp"
	mySQLConfig.Addr = config.dbAddress
	mySQLConfig.DBName = config.dbName
	mySQLConfig.ParseTime = true

	return mySQLConfig
}

func getEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("key %s not set in .env\n", key)
	}

	return val
}
