package main

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/suntoucha/jira"
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

	me := MyExporter{DB: db}
	for _, x := range PRJLIST {
		cli.ExportIssueByProject(x, 100, me)
	}
}

type MyExporter struct {
	DB *sqlx.DB
}

type IssueExport struct {
	Key string          `db:"key"`
	Raw json.RawMessage `db:"raw"`
}

func (me MyExporter) Export(raw json.RawMessage, index int, total int) error {
	ins := `insert into issue(key, raw) values(:key, :raw);`

	tmp, err := jira.IssueFromJson(raw)
	if err != nil {
		fmt.Println("IssueFromJson error:", err)
		return err
	}

	exp := IssueExport{Key: tmp.Key, Raw: raw}
	_, err = me.DB.NamedExec(ins, exp)
	if err != nil {
		fmt.Println("Insert error:", err)
		return err
	}

	ii := index / 100
	if ii*100 == index {
		fmt.Printf("[%v] %v index %v, total %v\n", time.Now().Format("2006-01-02 15:04:05"), tmp.Fields.Project.Key, index, total)
	}
	return nil
}
