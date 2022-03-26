package jira

import (
	"github.com/jmoiron/sqlx"
)

type IssueRaw struct {
	Key string `db:"key"`
	Raw string `db:"raw"`
}

type IssueRawTable struct {
	DB *sqlx.DB
}

func (t *IssueRawTable) Insert(x IssueRaw) error {
	if _, err := t.DB.NamedExec(`insert into issue_raw(key, raw) values(:key, :raw);`, x); err != nil {
		return err
	}

	return nil
}

func (t *IssueRawTable) CursorAll() (*sqlx.Rows, error) {
	rows, err := t.DB.Queryx("select key, raw from issue_raw;")

	return rows, err
}

func (t *IssueRawTable) Count() (int, error) {
	var cnt struct {
		Cnt int `db:"cnt"`
	}

	err := t.DB.Get(&cnt, "select count(1) cnt from issue_raw;")

	return cnt.Cnt, err
}
