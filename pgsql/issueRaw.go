package pgsql

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
