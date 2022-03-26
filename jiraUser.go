package jira

import (
	"encoding/json"
)

type JiraUser struct {
	Self         string `json:"self"`
	Key          string `json:"key"`
	AccountID    string `json:"accountId"`
	Name         string `json:"name"`
	EmailAddress string `json:"emailAddress"`
	DisplayName  string `json:"displayName"`
}

func JiraUserFromJson(raw []byte) (JiraUser, error) {
	var (
		user JiraUser
	)

	if err := json.Unmarshal(raw, &user); err != nil {
		return JiraUser{}, err
	}

	return user, nil
}
