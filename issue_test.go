package jira

import (
	_ "embed"
	"github.com/google/go-cmp/cmp"
	"testing"
)

//https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issues/#api-rest-api-2-issue-issueidorkey-get
//go:embed test_json/issue1.json
var issueJson1 string

//Cassby sample issue
//go:embed test_json/issue2.json
var issueJson2 string

func TestIssue(t *testing.T) {
	testcase := []struct {
		Input  string
		Result Issue
	}{
		{
			Input: issueJson1,
			Result: Issue{
				ID:   "10002",
				Self: "https://your-domain.atlassian.net/rest/api/2/issue/10002",
				Key:  "ED-1",
				Fields: IssueFields{
					Description: "Main order flow broken",
					Project: Project{
						Self: "https://your-domain.atlassian.net/rest/api/2/project/EX",
						ID:   "10000",
						Key:  "EX",
						Name: "Example",
					},
				},
			},
		},
		{
			Input: issueJson2,
			Result: Issue{
				ID:   "20281",
				Self: "https://xxx.atlassian.net/rest/api/2/issue/20281",
				Key:  "APPS-22",
				Fields: IssueFields{
					Description: "Some description here",
					Summary:     "Summary here",
					Project: Project{
						Self: "https://xxx.atlassian.net/rest/api/2/project/10401",
						ID:   "10401",
						Key:  "APPS",
						Name: "Apps",
					},
					Assignee: User{
						Self:         "https://xxx.atlassian.net/rest/api/2/user?accountId=1",
						Key:          "",
						AccountID:    "accountId1",
						Name:         "",
						EmailAddress: "",
						DisplayName:  "Alexey Ka",
					},
					Reporter: User{
						Self:         "https://xxx.atlassian.net/rest/api/2/user?accountId=2",
						Key:          "",
						AccountID:    "accountId2",
						Name:         "",
						EmailAddress: "xxx@gmail.com",
						DisplayName:  "Serge Ku",
					},
					Creator: User{
						Self:         "https://xxx.atlassian.net/rest/api/2/user?accountId=3",
						Key:          "",
						AccountID:    "accountId3",
						Name:         "",
						EmailAddress: "aaa@gmail.com",
						DisplayName:  "S K",
					},
					Status: IssueStatus{
						Self:        "https://xxx.atlassian.net/rest/api/2/status/10112",
						ID:          "10112",
						Name:        "Готово",
						Description: "",
						StatusCategory: IssueStatusCategory{
							Self:      "https://xxx.atlassian.net/rest/api/2/statuscategory/3",
							ID:        3,
							Key:       "done",
							ColorName: "green",
							Name:      "Done",
						},
					},
					IssueType: IssueType{
						Self:        "https://xxx.atlassian.net/rest/api/2/issuetype/10101",
						ID:          "10101",
						Name:        "Баг",
						Description: "Bugs track problems or errors.",
						Subtask:     false,
					},
					Created:        NewDatetimeMustCompile("2018-05-02T20:41:01.285+0300"),
					Updated:        NewDatetimeIgnore("2018-07-16T14:19:48.041+0300"),
					ResolutionDate: NewDatetimeMustCompile("2018-07-16T14:19:48.037+0300"),
				},
			},
		},
	}

	for _, x := range testcase {
		tmp, err := IssueFromJson([]byte(x.Input))
		if err != nil {
			t.Errorf("IssueFromJson error: %v", err)
			continue
		}

		if diff := cmp.Diff(x.Result, tmp); diff != "" {
			t.Errorf("IssueFromJson diff: %v", diff)
			continue
		}

	}
}
