package jira

import (
	"encoding/json"
)

type IssueResult struct {
	Expand    string            `json:"expand"`
	StartAt   int               `json:"startAt"`
	MaxResult int               `json:"maxResults"`
	Total     int               `json:"total"`
	Issues    []json.RawMessage `json:"issues"`
}

func IssueResultFromJson(raw []byte) (IssueResult, error) {
	var (
		res IssueResult
	)

	if err := json.Unmarshal(raw, &res); err != nil {
		return IssueResult{}, err
	}

	return res, nil
}

func (res *IssueResult) IssueList() (IssueList, error) {
	var (
		list IssueList
	)

	for _, x := range res.Issues {
		tmp, err := IssueFromJson(x)
		if err != nil {
			return nil, err
		}
		list = append(list, tmp)
	}

	return list, nil
}
