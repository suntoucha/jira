package main

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/suntoucha/jira"
	"github.com/suntoucha/jira/pgsql"
	"time"
)

var (
	HOST, USER, TOKEN string
	PGSQL             string
	PRJLIST           []string
)

func main() {
	fmt.Println("Hello World")

	cli := jira.Client{Host: HOST, User: USER, Token: TOKEN}

	db, err := sqlx.Open("postgres", PGSQL)
	defer db.Close()
	if err != nil {
		fmt.Println("PGSQL error:", err)
		return
	}

	tableRaw := pgsql.IssueRawTable{DB: db}
	tableFlat := pgsql.IssueFlatTable{DB: db}
	me := MyExporter{Table: &tableRaw}
	for _, x := range PRJLIST {
		cli.ExportIssueByProject(x, 100, me)
	}

	rows, err := tableRaw.CursorAll()
	if err != nil {
		fmt.Println("Cursor error:", err)
		return
	}
	for rows.Next() {
		var (
			tmp  jira.JiraIssue
			raw  pgsql.IssueRaw
			flat pgsql.IssueFlat
		)

		if err = rows.StructScan(&raw); err != nil {
			fmt.Println("Scan error", err)
			return
		}
		fmt.Println(raw.Key)

		if tmp, err = jira.JiraIssueFromJson([]byte(raw.Raw)); err != nil {
			fmt.Println("From json error:", err)
			return
		}

		flat = pgsql.IssueToFlat(tmp)
		if err = tableFlat.Insert(flat); err != nil {
			fmt.Println("Insert error:", err)
			return
		}
	}
}

type MyExporter struct {
	Table *pgsql.IssueRawTable
}

func (me MyExporter) Export(raw json.RawMessage, index int, total int) error {
	tmp, err := jira.JiraIssueFromJson(raw)
	if err != nil {
		fmt.Println("IssueFromJson error:", err)
		return err
	}

	x := pgsql.IssueRaw{Key: tmp.Key, Raw: string(raw)}
	if err := me.Table.Insert(x); err != nil {
		fmt.Println("Insert error:", err)
		return err
	}

	ii := index / 100
	if ii*100 == index {
		fmt.Printf("[%v] %v index %v, total %v\n", time.Now().Format("2006-01-02 15:04:05"), tmp.Fields.Project.Key, index, total)
	}
	return nil
}
