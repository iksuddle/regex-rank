package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	"suddle.dev/regex-rank/types"
)

type UserStore struct {
	Db *sqlx.DB
}

func (s *UserStore) CreateUser(user *types.User) error {
	_, err := s.Db.Exec("INSERT INTO users (id, username, avatar_url, created_at) VALUES (?, ?, ?, ?)", user.Id, user.Username, user.AvatarUrl, user.CreatedAt)
	if err != nil {
		return err
	}

	log.Printf("created new user `%s` with id %d\n", user.Username, user.Id)
	return nil
}
