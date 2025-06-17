package database

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"suddle.dev/regex-rank/config"
)

func NewDB(config config.Config) *sqlx.DB {
	mySQLConfig := mysql.NewConfig()

	mySQLConfig.User = config.DBUser
	mySQLConfig.Passwd = config.DBPassword
	mySQLConfig.DBName = config.DBName
	mySQLConfig.ParseTime = true

	db, err := sqlx.Connect("mysql", mySQLConfig.FormatDSN())
	if err != nil {
		log.Fatal("error connecting to db:", err)
	}

	return db
}
