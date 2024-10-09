package database

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/iksuddle/regex-rank/config"
	"github.com/jmoiron/sqlx"
)

func NewDB(config *config.Config) *sqlx.DB {
    cfg := newMySQLConfig(config)

	db, err := sqlx.Connect("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal("error when connecting to database: ", err)
	}

	log.Println("successfully connected to database.")
	return db
}

func newMySQLConfig(config *config.Config) *mysql.Config {
	mySQLConfig := mysql.NewConfig()

	mySQLConfig.User = config.DbUser
	mySQLConfig.Passwd = config.DbPassword
	mySQLConfig.Net = "tcp"
	mySQLConfig.Addr = config.DbAddress
	mySQLConfig.DBName = config.DbName
	mySQLConfig.ParseTime = true

	return mySQLConfig
}

