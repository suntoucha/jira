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
				Key:           "APPS-22",
				ProjectKey:    "APPS",
				Description:   "Some description here",
				Summary:       "Summary here",
				TypeID:        "10101",
				TypeName:      "Баг",
				IsSubtask:     false,
				StatusID:      "10112",
				StatusName:    "Готово",
				ReporterEmail: "xxx@gmail.com",
				Created:       NewDatetimeMustCompile("2018-05-02T20:41:01.285+0300").Time,
				Updated:       NewDatetimeIgnore("2018-07-16T14:19:48.041+0300").Time,
				Resolution:    NewDatetimeMustCompile("2018-07-16T14:19:48.037+0300").Time,
				Raw:           issueSqlJson2,
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
