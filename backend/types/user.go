package types

type User struct {
	Id        int    `json:"id" db:"id"`
	Username  string `json:"username" db:"username"`
	AvatarUrl string `json:"avatar_url" db:"avatar_url"`
}

func NewUserFromData(userData map[string]any) (*User, error) {
	userGithubId := userData["id"].(float64)
	user := &User{
		Id:        int(userGithubId),
		Username:  userData["login"].(string),
		AvatarUrl: userData["avatar_url"].(string),
	}

	return user, nil
}
