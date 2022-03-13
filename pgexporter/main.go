package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/suntoucha/jira"
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
	cli.ExportIssueByProject(PRJKEY, 10, me)
}

type MyExporter struct {
	DB *sqlx.DB
}

func (me MyExporter) Export(list jira.IssueList, startAt int, maxResult int, total int) error {
	ins := "insert into issue(key, project_key, raw) values(:key, :project_key, :raw);"

	for _, x := range list {
		tmp := jira.IssueToSql(x)
		_, err := me.DB.NamedExec(ins, tmp)
		if err != nil {
			fmt.Println("Insert error:", err)
			return err
		}
		fmt.Println("OK")
	}

	fmt.Printf("startAt %v, maxResult %v, total %v, issue-len %v\n", startAt, maxResult, total, len(list))
	return nil
}
