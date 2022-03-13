package jira

import (
	_ "embed"
	"github.com/google/go-cmp/cmp"
	"testing"
)

//https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issues/#api-rest-api-2-issue-issueidorkey-get
//go:embed test_json/issue1.json
var issueSqlJson1 string

//Cassby sample issue
//go:embed test_json/issue2.json
var issueSqlJson2 string

func TestIssueSql(t *testing.T) {
	testcase := []struct {
		Input  string
		Result IssueSql
	}{
		{
			Input: issueSqlJson1,
			Result: IssueSql{
				Key:         "ED-1",
				ProjectKey:  "EX",
				Description: "Main order flow broken",
				Summary:     "",
				Raw:         issueSqlJson1,
			},
		},
		{
			Input: issueSqlJson2,
			Result: IssueSql{
				Key:         "APPS-22",
				ProjectKey:  "APPS",
				Description: "Some description here",
				Summary:     "Summary here",
				Raw:         issueSqlJson2,
			},
		},
	}

	for _, x := range testcase {
		tmp, err := IssueFromJson([]byte(x.Input))
		if err != nil {
			t.Errorf("IssueFromJson error: %v", err)
			continue
		}
		s := IssueToSql(tmp)

		if diff := cmp.Diff(x.Result, s); diff != "" {
			t.Errorf("IssueToSql diff: %v", diff)
			continue
		}

	}
}
