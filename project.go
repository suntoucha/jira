package main

import (
	"encoding/json"
)

type Project struct {
	Self        string `json:"self"`
	Expand      string `json:"expand"`
	ID          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProjectList []Project

func ProjectListFromJson(raw []byte) (ProjectList, error) {
	var (
		list ProjectList
	)

	if err := json.Unmarshal(raw, &list); err != nil {
		return nil, err
	}

	return list, nil
}

func (p *Project) Issue() IssueCursor {
	return IssueCursor{ProjectKey: p.Key, StartAt: 0, MaxResults: 10}
}
