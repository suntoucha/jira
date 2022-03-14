package jira

import (
	"time"
)

type IssueSql struct {
	Key           string    `db:"key"`
	ProjectKey    string    `db:"project_key"`
	Description   string    `db:"description"`
	Summary       string    `db:"summary"`
	TypeID        string    `db:"type_id"`
	TypeName      string    `db:"type_name"`
	IsSubtask     bool      `db:"is_subtask"`
	StatusID      string    `db:"status_id"`
	StatusName    string    `db:"status_name"`
	AssigneeEmail string    `db:"assignee_email"`
	ReporterEmail string    `db:"reporter_email"`
	Created       time.Time `db:"dt_created"`
	Updated       time.Time `db:"dt_updated"`
	Resolution    time.Time `db:"dt_resolution"`
	Raw           string    `db:"raw"`
}

func IssueToSql(i Issue) IssueSql {
	var x IssueSql

	x.Key = i.Key
	x.ProjectKey = i.Fields.Project.Key
	x.Description = i.Fields.Description
	x.Summary = i.Fields.Summary
	x.TypeID = i.Fields.IssueType.ID
	x.TypeName = i.Fields.IssueType.Name
	x.IsSubtask = i.Fields.IssueType.Subtask
	x.StatusID = i.Fields.Status.ID
	x.StatusName = i.Fields.Status.Name
	x.AssigneeEmail = i.Fields.Assignee.EmailAddress
	x.ReporterEmail = i.Fields.Reporter.EmailAddress
	x.Created = i.Fields.Created.Time
	x.Updated = i.Fields.Updated.Time
	x.Resolution = i.Fields.ResolutionDate.Time

	x.Raw = i.Raw

	return x
}
