package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/iksuddle/regex-rank/types"
	"github.com/jmoiron/sqlx"
)

type UserStore struct {
	db *sqlx.DB
}

func NewUserStore(db *sqlx.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

// returns the user if they exist in the db, otherwise returns nil with err
func (s *UserStore) GetUserById(id int) (*types.User, error) {
	user := &types.User{}
	err := s.db.Get(user, "SELECT * FROM users WHERE id=?", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("user %d not found\n", id)
		}
		return nil, err
	}

	return user, nil
}

// todo: prepare statement
func (s *UserStore) CreateUser(user *types.User) error {
	_, err := s.db.Exec("INSERT INTO users (id, username, avatar_url) VALUES (?, ?, ?)", user.Id, user.Username, user.AvatarUrl)
	if err != nil {
		return err
	}

    log.Printf("created new user `%s` with id %d\n", user.Username, user.Id)
	return nil
}
