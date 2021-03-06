package jira

import (
	_ "embed"
	"github.com/google/go-cmp/cmp"
	"testing"
)

//https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-users/#api-rest-api-2-user-get
//go:embed test_json/user1.json
var userJson1 string

func TestJiraUser(t *testing.T) {
	testcase := []struct {
		Input  string
		Result JiraUser
	}{
		{
			Input: userJson1,
			Result: JiraUser{
				Self:         "https://your-domain.atlassian.net/rest/api/2/user?accountId=5b10a2844c20165700ede21g",
				AccountID:    "5b10a2844c20165700ede21g",
				EmailAddress: "mia@example.com",
				DisplayName:  "Mia Krystof",
			},
		},
	}

	for _, x := range testcase {
		tmp, err := JiraUserFromJson([]byte(x.Input))
		if err != nil {
			t.Errorf("UserFromJson error: %v", err)
			continue
		}

		if diff := cmp.Diff(x.Result, tmp); diff != "" {
			t.Errorf("UserFromJson diff: %v", diff)
			continue
		}

	}
}
