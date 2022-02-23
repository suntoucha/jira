package jira

import (
	"encoding/json"
)

type Issue struct {
	ID     string      `json:"id"`
	Self   string      `json:"self"`
	Key    string      `json:"key"`
	Fields IssueFields `json:"fields"`
}

type IssueFields struct {
	Description string      `json:"description"`
	Summary     string      `json:"summary"`
	Project     Project     `json:"project"`
	Assignee    User        `json:"assignee"`
	Reporter    User        `json:"reporter"`
	Creator     User        `json:"creator"`
	Status      IssueStatus `json::status:`
	IssueType   IssueType   `json:"issuetype"`
	Created     Datetime    `json:"created"`
	//Updated     Datetime    `json:"updated"`
	ResolutionDate Datetime `json:"resolutiondate"`
}

type IssueType struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Subtask     bool   `json:"subtask"`
}

type IssueStatus struct {
	Self           string              `json:"self"`
	ID             string              `json:"ID"`
	Name           string              `json:"name"`
	Description    string              `json:"description"`
	StatusCategory IssueStatusCategory `json:"statusCategory"`
}

type IssueStatusCategory struct {
	Self      string `json:"self"`
	ID        int    `json:"ID"`
	Key       string `json:"key"`
	ColorName string `json:"colorName"`
	Name      string `json:"name"`
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
