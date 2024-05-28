package main

type User struct {
	Id        int    `json:"id"`
	GitHubId  int    `json:"github_id"`
	Username  string `json:"username"`
	AvatarUrl string `json:"avatar_url"`
}

func NewUserFromData(userData map[string]any) (*User, error) {
	userGithubId := userData["id"].(float64)
	user := &User{
		GitHubId:  int(userGithubId),
		Username:  userData["login"].(string),
		AvatarUrl: userData["avatar_url"].(string),
	}

	return user, nil
}
