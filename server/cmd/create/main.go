package main

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/iksuddle/regex-rank/config"
	"github.com/iksuddle/regex-rank/database"
	"github.com/iksuddle/regex-rank/types"
)

func main() {
	path := os.Args[len(os.Args)-1]

	config := config.NewConfig()
	db := database.NewDB(database.NewMySQLConfig(config))

	var statements struct {
		Match  []string
		Ignore []string
	}

	_, err := toml.DecodeFile(path, &statements)

	if err != nil {
		log.Fatal(err)
	}

	p := types.NewProblem()

	result, err := db.Exec("INSERT INTO problems (created_at) VALUES (?)", p.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("problem #%d created\n", id)

	for _, s := range statements.Match {
		_, err := db.Exec("INSERT INTO statements (problem_id, `match`, literal) VALUES (?, ?, ?)", id, true, s)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, s := range statements.Ignore {
		_, err := db.Exec("INSERT INTO statements (problem_id, `match`, literal) VALUES (?, ?, ?)", id, false, s)
		if err != nil {
			log.Fatal(err)
		}
	}
}
