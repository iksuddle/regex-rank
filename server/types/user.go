package types

import "time"

type UserData struct {
	Id        int64  `json:"id"`
	Username  string `json:"login"`
	AvatarUrl string `json:"avatar_url"`
}

type User struct {
	Id        int64     `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	AvatarUrl string    `json:"avatar_url" db:"avatar_url"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func NewUser(data UserData) User {
	return User{
		Id:        data.Id,
		Username:  data.Username,
		AvatarUrl: data.AvatarUrl,
		CreatedAt: time.Now().UTC(),
	}
}
