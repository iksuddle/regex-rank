package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

var Envs = initConfig()

func NewMySQLConfig() *mysql.Config {
	mysqlConfig := mysql.NewConfig()

	mysqlConfig.User = Envs.DBUser
	mysqlConfig.Passwd = Envs.DBPassword
	mysqlConfig.Net = "tcp"
	mysqlConfig.Addr = Envs.DBAddress
	mysqlConfig.DBName = Envs.DBName

	return mysqlConfig
}

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not read .env file")
	}

	return Config{
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "mypassword"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:     getEnv("DB_NAME", "rgx"),
	}
}

func getEnv(key string, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return val
}
