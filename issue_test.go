package jira

import (
	_ "embed"
	"github.com/google/go-cmp/cmp"
	"testing"
)

//https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-projects/#api-rest-api-2-project-search-get
//go:embed issue1.json
var issueJson1 string

func TestIssue(t *testing.T) {
	testcase := []struct {
		Input  string
		Result Issue
	}{
		{
			Input: issueJson1,
			Result: Issue{
				Self: "https://your-domain.atlassian.net/rest/api/2/issue/10002",
				ID:   "10002",
				Key:  "ED-1",
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
