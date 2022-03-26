package jira

import (
	"encoding/json"
)

type JiraProject struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func JiraProjectFromJson(raw []byte) (JiraProject, error) {
	var (
		prj JiraProject
	)

	if err := json.Unmarshal(raw, &prj); err != nil {
		return JiraProject{}, err
	}

	return prj, nil
}
