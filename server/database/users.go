package database

import (
	"database/sql"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
	"suddle.dev/regex-rank/types"
)

type UserStore struct {
	Db *sqlx.DB
}

func (s *UserStore) CreateUser(user types.User) error {
	_, err := s.Db.Exec("INSERT INTO users (id, username, avatar_url, created_at) VALUES (?, ?, ?, ?)",
		user.Id, user.Username, user.AvatarUrl, user.CreatedAt)
	if err != nil {
		return err
	}

	log.Printf("created new user `%s` with id %d\n", user.Username, user.Id)
	return nil
}

// returns the user if they exist in the db, otherwise returns err
func (s *UserStore) GetUserById(id int64) (types.User, error) {
	user := types.User{}
	err := s.Db.Get(&user, "SELECT * FROM users WHERE id=?", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("user %d not found\n", id)
		}
		return user, err
	}

	return user, nil
}
