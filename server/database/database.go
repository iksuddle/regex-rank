package database

import (
	"log"

	"github.com/go-sql-driver/mysql"
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
