package types

import "time"

type User struct {
	Id        int    `json:"id" db:"id"`
	Username  string `json:"username" db:"username"`
	AvatarUrl string `json:"avatar_url" db:"avatar_url"`
	CreatedAt int64  `json:"created_at" db:"created_at"`
}

func NewUserFromData(userData map[string]any) (*User, error) {
	userGithubId := userData["id"].(float64)
	user := &User{
		Id:        int(userGithubId),
		Username:  userData["login"].(string),
		AvatarUrl: userData["avatar_url"].(string),
		CreatedAt: time.Now().Unix(),
	}

	return user, nil
}
