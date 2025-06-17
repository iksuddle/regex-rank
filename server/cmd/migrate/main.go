package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"suddle.dev/regex-rank/config"
	"suddle.dev/regex-rank/database"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: migrate [up|down]")
		os.Exit(1)
	}

	godotenv.Load()

	c := config.NewConfig()
	db := database.NewDB(c)

	driver, err := mysql.WithInstance(db.DB, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "mysql", driver)

	if os.Args[1] == "up" {
		err = m.Up()
		if err != nil {
			fmt.Println(err)
		}
	} else if os.Args[1] == "down" {
		err = m.Down()
		if err != nil {
			fmt.Println(err)
		}
	}
}
