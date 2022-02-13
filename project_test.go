package jira

import (
	_ "embed"
	"github.com/google/go-cmp/cmp"
	"testing"
)

//go:embed project1.json
var projectJson1 string

func TestProject(t *testing.T) {
	testcase := []struct {
		Input  string
		Result Project
	}{
		{
			Input: projectJson1,
			Result: Project{
				Self:        "https://your-domain.atlassian.net/rest/api/2/project/EX",
				ID:          "10000",
				Key:         "EX",
				Name:        "Example",
				Description: "This project was created as an example for REST.",
			},
		},
	}

	for _, x := range testcase {
		tmp, err := ProjectFromJson([]byte(x.Input))
		if err != nil {
			t.Errorf("ProjectListFromJson error: %v", err)
			continue
		}

		if diff := cmp.Diff(x.Result, tmp); diff != "" {
			t.Errorf("ProjectFromJson diff: %v", diff)
			continue
		}

	}
}
