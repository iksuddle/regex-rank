package database

import (
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

	return nil
}

// returns the user if they exist in the db, otherwise returns err
func (s *UserStore) GetUserById(id int64) (types.User, error) {
	user := types.User{}
	err := s.Db.Get(&user, "SELECT * FROM users WHERE id=?", id)
	if err != nil {
		return user, err
	}

	return user, nil
}
