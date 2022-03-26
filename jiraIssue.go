package jira

import (
	"encoding/json"
)

type JiraIssue struct {
	ID     string          `json:"id"`
	Self   string          `json:"self"`
	Key    string          `json:"key"`
	Fields JiraIssueFields `json:"fields"`
}

type JiraIssueFields struct {
	Description    string          `json:"description"`
	Summary        string          `json:"summary"`
	Project        JiraProject     `json:"project"`
	Assignee       JiraUser        `json:"assignee"`
	Reporter       JiraUser        `json:"reporter"`
	Creator        JiraUser        `json:"creator"`
	Status         JiraIssueStatus `json::status:`
	IssueType      JiraIssueType   `json:"issuetype"`
	Created        Datetime        `json:"created"`
	Updated        DatetimeIgnore  `json:"updated"`
	ResolutionDate Datetime        `json:"resolutiondate"`
}

type JiraIssueType struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Subtask     bool   `json:"subtask"`
}

type JiraIssueStatus struct {
	Self           string                  `json:"self"`
	ID             string                  `json:"ID"`
	Name           string                  `json:"name"`
	Description    string                  `json:"description"`
	StatusCategory JiraIssueStatusCategory `json:"statusCategory"`
}

type JiraIssueStatusCategory struct {
	Self      string `json:"self"`
	ID        int    `json:"ID"`
	Key       string `json:"key"`
	ColorName string `json:"colorName"`
	Name      string `json:"name"`
}

type JiraIssueList []JiraIssue

func JiraIssueFromJson(raw []byte) (JiraIssue, error) {
	var (
		res JiraIssue
	)

	if err := json.Unmarshal(raw, &res); err != nil {
		return JiraIssue{}, err
	}

	return res, nil
}
