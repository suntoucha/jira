package jira

import (
	"encoding/json"
)

type JiraIssueResult struct {
	Expand    string            `json:"expand"`
	StartAt   int               `json:"startAt"`
	MaxResult int               `json:"maxResults"`
	Total     int               `json:"total"`
	Issues    []json.RawMessage `json:"issues"`
}

func JiraIssueResultFromJson(raw []byte) (JiraIssueResult, error) {
	var (
		res JiraIssueResult
	)

	if err := json.Unmarshal(raw, &res); err != nil {
		return JiraIssueResult{}, err
	}

	return res, nil
}

func (res *JiraIssueResult) IssueList() (JiraIssueList, error) {
	var (
		list JiraIssueList
	)

	for _, x := range res.Issues {
		tmp, err := JiraIssueFromJson(x)
		if err != nil {
			return nil, err
		}
		list = append(list, tmp)
	}

	return list, nil
}
