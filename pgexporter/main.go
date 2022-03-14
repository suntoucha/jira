package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/suntoucha/jira"
	"time"
)

var (
	HOST, USER, TOKEN string
	PRJKEY, ISSUEKEY  string
	PGSQL             string
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

	me := MyExporter{DB: db}
	cli.ExportIssueByProject(PRJKEY, 100, me)
}

type MyExporter struct {
	DB *sqlx.DB
}

func (me MyExporter) Export(list jira.IssueList, startAt int, maxResult int, total int) error {
	ins := `insert into issue(key, project_key, description, summary, type_id, type_name, is_subtask, status_id, status_name, assignee_email, reporter_email, dt_created, dt_updated, dt_resolution, raw) 
		values(:key, :project_key, :description, :summary, :type_id, :type_name, :is_subtask, :status_id, :status_name, :assignee_email, :reporter_email, :dt_created, :dt_updated, :dt_resolution, :raw);`

	for _, x := range list {
		tmp := jira.IssueToSql(x)
		_, err := me.DB.NamedExec(ins, tmp)
		if err != nil {
			fmt.Println("Insert error:", err)
			return err
		}
	}

	fmt.Printf("[%v] startAt %v, maxResult %v, total %v, issue-len %v\n", time.Now().Format("2006-01-02 15:04:05"), startAt, maxResult, total, len(list))
	return nil
}
