package jira

import (
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
