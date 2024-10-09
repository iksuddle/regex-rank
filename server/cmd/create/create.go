package main

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/iksuddle/regex-rank/config"
	"github.com/iksuddle/regex-rank/database"
	"github.com/iksuddle/regex-rank/types"
)

// read toml containing problem and store in db with unix date as key
func main() {
	// read path from args
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("usage: make create [path]")
	}

	path := args[0]

	// create new problem store
	db := database.NewDB(config.NewConfig())
	store := database.NewProblemStore(db)

	var statements struct {
		Match  []string
		Ignore []string
	}

	_, err := toml.DecodeFile(path, &statements)
	if err != nil {
		log.Fatal(err)
	}

	// create problem
	problem := types.NewProblem()
	problemId, err := store.CreateProblem(&problem)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range statements.Match {
		statement := types.NewStatement(problemId, "m", s)
		_, err := store.CreateStatement(&statement)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, s := range statements.Ignore {
		statement := types.NewStatement(problemId, "i", s)
		store.CreateStatement(&statement)
		if err != nil {
			log.Fatal(err)
		}
	}
}
