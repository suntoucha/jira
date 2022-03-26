package pgsql

import (
	"github.com/jmoiron/sqlx"
	"github.com/suntoucha/jira"
	"time"
)

type IssueFlat struct {
	Key         string    `db:"key"`
	Project     string    `db:"project"`
	Description string    `db:"description"`
	Summary     string    `db:"summary"`
	Type        string    `db:"type"`
	TypeName    string    `db:"type_name"`
	IsSubtask   bool      `db:"is_subtask"`
	Status      string    `db:"status"`
	StatusName  string    `db:"status_name"`
	Assignee    string    `db:"assignee"`
	Reporter    string    `db:"reporter"`
	Created     time.Time `db:"dt_created"`
	Updated     time.Time `db:"dt_updated"`
	Resolution  time.Time `db:"dt_resolution"`
}

func IssueToFlat(i jira.Issue) IssueFlat {
	var x IssueFlat

	x.Key = i.Key
	x.Project = i.Fields.Project.Key
	x.Description = i.Fields.Description
	x.Summary = i.Fields.Summary
	x.Type = i.Fields.IssueType.ID
	x.TypeName = i.Fields.IssueType.Name
	x.IsSubtask = i.Fields.IssueType.Subtask
	x.Status = i.Fields.Status.ID
	x.StatusName = i.Fields.Status.Name
	x.Assignee = i.Fields.Assignee.EmailAddress
	x.Reporter = i.Fields.Reporter.EmailAddress
	x.Created = i.Fields.Created.Time
	x.Updated = i.Fields.Updated.Time
	x.Resolution = i.Fields.ResolutionDate.Time

	return x
}

type IssueFlatTable struct {
	DB *sqlx.DB
}

func (t *IssueFlatTable) Insert(x IssueFlat) error {
	ins := `insert into issue_flat(key, project, description, summary, type, type_name, is_subtask, status, status_name, assignee, reporter, dt_created, dt_updated, dt_resolution) 
		values(:key, :project, :description, :summary, :type, :type_name, :is_subtask, :status, :status_name, :assignee, :reporter, :dt_created, :dt_updated, :dt_resolution);`

	if _, err := t.DB.NamedExec(ins, x); err != nil {
		return err
	}

	return nil
}
