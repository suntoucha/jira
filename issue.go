package main

import (
	"encoding/json"
	"strconv"
)

type IssueCursor struct {
	ProjectKey          string
	StartAt, MaxResults int
}

type IssueResult struct {
	Expand    string    `json:"expand"`
	StartAt   int       `json:"startAt"`
	MaxResult int       `json:"maxResults"`
	Total     int       `json:"total"`
	Issues    IssueList `json:"issues"`
}

func (cur *IssueCursor) Resourse() string {
	return "/rest/api/2/search?jql=project=" + cur.ProjectKey + "+order+by+key&startAt=" + strconv.Itoa(cur.StartAt) + "&maxResults=" + strconv.Itoa(cur.MaxResults)
}

func (cur *IssueCursor) Next() {
	cur.StartAt += cur.MaxResults
}

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
