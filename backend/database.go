package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB(cfg *mysql.Config) {
	db, err := sqlx.Connect("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal("error when connecting to database: ", err)
	}

	log.Println("successfully connected to database.")
	DB = db
}
