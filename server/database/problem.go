package database

import (
	"log"

	"github.com/iksuddle/regex-rank/types"
	"github.com/jmoiron/sqlx"
)

type ProblemStore struct {
	db *sqlx.DB
}

func NewProblemStore(db *sqlx.DB) *ProblemStore {
	return &ProblemStore{
		db: db,
	}
}

func (s *ProblemStore) CreateProblem(problem *types.Problem) (int, error) {
	result, err := s.db.Exec("INSERT INTO problems (created_at) VALUES (?)", problem.CreatedAt)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	log.Printf("created problem %d.\n", lastId)
	return int(lastId), nil
}

func (s *ProblemStore) CreateStatement(statement *types.Statement) (int, error) {
	result, err := s.db.Exec("INSERT INTO statements (problem_id, `match`, literal) VALUES (?, ?, ?)", statement.ProblemId, statement.Match, statement.Literal)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	log.Printf("created statement %d for problem %d.\n", lastId, statement.ProblemId)
	return int(lastId), nil
}
