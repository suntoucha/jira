package jira

import (
	_ "embed"
	"github.com/google/go-cmp/cmp"
	"testing"
)

//https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-projects/#api-rest-api-2-project-search-get
//go:embed test_json/project1.json
var projectJson1 string

func TestJiraProject(t *testing.T) {
	testcase := []struct {
		Input  string
		Result JiraProject
	}{
		{
			Input: projectJson1,
			Result: JiraProject{
				Self:        "https://your-domain.atlassian.net/rest/api/2/project/EX",
				ID:          "10000",
				Key:         "EX",
				Name:        "Example",
				Description: "This project was created as an example for REST.",
			},
		},
	}

	for _, x := range testcase {
		tmp, err := JiraProjectFromJson([]byte(x.Input))
		if err != nil {
			t.Errorf("ProjectFromJson error: %v", err)
			continue
		}

		if diff := cmp.Diff(x.Result, tmp); diff != "" {
			t.Errorf("ProjectFromJson diff: %v", diff)
			continue
		}

	}
}
