package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/suntoucha/jira"
	"github.com/cheggaaa/pb/v3"
	"encoding/json"
)

var (
	HOST, USER, TOKEN string
	PGSQL             string
	PRJLIST           []string
)

func main() {
	fmt.Println("JIRA exporter. Project list:", PRJLIST)

	cli := jira.Client{Host: HOST, User: USER, Token: TOKEN}

	db, err := sqlx.Open("postgres", PGSQL)
	defer db.Close()
	if err != nil {
		fmt.Println("PGSQL error:", err)
		return
	}

	tableRaw := jira.IssueRawTable{DB: db}
	for _, prj := range PRJLIST {
		maxResults := 100
		res, err := cli.IssueByProject(prj, 0, maxResults)	

		bar := pb.StartNew(res.Total)
		bar.SetTemplate(pb.Full)
		bar.Set("prefix", "["+prj+"] ")

		for err == nil && len(res.Issues) > 0 {
			for _, raw := range res.Issues {
				if err := export(raw, &tableRaw); err != nil {
					fmt.Println("Export error:", err)
					return
				}
				bar.Increment()
			}

			res, err = cli.IssueByProject(prj, res.StartAt+len(res.Issues), maxResults)
		}

		bar.Finish()
		if err != nil {
			fmt.Println("Exporter error:", err)
			return
		}
	}
}

func export(raw json.RawMessage, table *jira.IssueRawTable) error {
	tmp, err := jira.JiraIssueFromJson(raw)
	if err != nil {
		return err
	}

	x := jira.IssueRaw{Key: tmp.Key, Raw: string(raw)}
	if err := table.Insert(x); err != nil {
		return err
	}

	return nil	
}