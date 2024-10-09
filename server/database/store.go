package database

import (
	"github.com/jmoiron/sqlx"
)

type Store struct {
	Users    *UserStore
	Problems *ProblemStore
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		Users: NewUserStore(db),
        Problems: NewProblemStore(db),
	}
}
