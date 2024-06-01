package database

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/iksuddle/regex-rank/config"
	"github.com/jmoiron/sqlx"
)

func NewDB(cfg *mysql.Config) *sqlx.DB {
	db, err := sqlx.Connect("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal("error when connecting to database: ", err)
	}

	log.Println("successfully connected to database.")
	return db
}

func NewMySQLConfig(config *config.Config) *mysql.Config {
	mysqlConfig := mysql.NewConfig()

	mysqlConfig.User = config.DBUser
	mysqlConfig.Passwd = config.DBPassword
	mysqlConfig.Net = "tcp"
	mysqlConfig.Addr = config.DBAddress
	mysqlConfig.DBName = config.DBName

	return mysqlConfig
}
