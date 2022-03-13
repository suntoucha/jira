package jira

type IssueSql struct {
	Key         string `db:"key"`
	ProjectKey  string `db:"project_key"`
	Description string
	Summary     string
	Raw         string `db:"raw"`
}

func IssueToSql(i Issue) IssueSql {
	var x IssueSql

	x.Key = i.Key
	x.ProjectKey = i.Fields.Project.Key
	x.Description = i.Fields.Description
	x.Summary = i.Fields.Summary
	x.Raw = i.Raw

	return x
}
