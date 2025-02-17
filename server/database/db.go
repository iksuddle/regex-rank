package database

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"suddle.dev/regex-rank/config"
)

func NewDB(config config.Config) *sqlx.DB {
	mySqlConfig := newMySQLConfig(config)

	db, err := sqlx.Connect("mysql", mySqlConfig.FormatDSN())
	if err != nil {
		log.Fatal("error connecting to db:", err)
	}

	log.Println("successfully connected to db")
	return db
}

func newMySQLConfig(config config.Config) *mysql.Config {
	mySQLConfig := mysql.NewConfig()

	mySQLConfig.User = config.DBUser
	mySQLConfig.Passwd = config.DBPassword
	mySQLConfig.DBName = config.DBName
	mySQLConfig.ParseTime = true

	return mySQLConfig
}
