package main

import (
	"encoding/json"
)

type User struct {
	GitHubId  int    `json:"github_id"`
	Username  string `json:"username"`
	AvatarUrl string `json:"avatar_url"`
}

func NewUserFromJson(data []byte) (*User, error) {
	var userData map[string]any

	err := json.Unmarshal(data, &userData)
	if err != nil {
		return nil, err
	}

	userGithubId := userData["id"].(float64)
	user := &User{
		GitHubId:  int(userGithubId),
		Username:  userData["login"].(string),
		AvatarUrl: userData["avatar_url"].(string),
	}

	return user, nil
}

