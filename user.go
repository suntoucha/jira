package jira

import (
	"encoding/json"
)

type User struct {
	Self         string `json:"self"`
	Key          string `json:"key"`
	AccountID    string `json:"accountId"`
	Name         string `json:"name"`
	EmailAddress string `json:"emailAddress"`
	DisplayName  string `json:"displayName"`
}

func UserFromJson(raw []byte) (User, error) {
	var (
		user User
	)

	if err := json.Unmarshal(raw, &user); err != nil {
		return User{}, err
	}

	return user, nil
}
