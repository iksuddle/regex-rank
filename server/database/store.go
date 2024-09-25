package database

import (
	"github.com/jmoiron/sqlx"
)

type Store struct {
	Users *UserStore
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		Users: NewUserStore(db),
	}
}
