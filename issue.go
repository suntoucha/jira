package jira

import (
	"encoding/json"
)

type Issue struct {
	ID  string `json:"id"`
	Key string `json:"key"`
}

type IssueList []Issue

func IssueListFromJson(raw []byte) (IssueList, error) {
	var (
		res IssueResult
	)

	if err := json.Unmarshal(raw, &res); err != nil {
		return nil, err
	}

	return res.Issues, nil
}
