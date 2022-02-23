package jira

import (
	"encoding/json"
)

type Project struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ProjectFromJson(raw []byte) (Project, error) {
	var (
		prj Project
	)

	if err := json.Unmarshal(raw, &prj); err != nil {
		return Project{}, err
	}

	return prj, nil
}
