package jira

import (
	"encoding/json"
)

type Issue struct {
	Self string `json:"self"`
	ID   string `json:"id"`
	Key  string `json:"key"`
}

type IssueList []Issue

func IssueFromJson(raw []byte) (Issue, error) {
	var (
		res Issue
	)

	if err := json.Unmarshal(raw, &res); err != nil {
		return Issue{}, err
	}

	return res, nil
}

func IssueListFromJson(raw []byte) (IssueList, error) {
	var (
		res IssueResult
	)

	if err := json.Unmarshal(raw, &res); err != nil {
		return nil, err
	}

	return res.Issues, nil
}
