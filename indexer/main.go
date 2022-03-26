package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/suntoucha/jira"
	"github.com/cheggaaa/pb/v3"

)

var (
	PGSQL             string
)

func main() {
	fmt.Println("JIRA indexer")

	db, err := sqlx.Open("postgres", PGSQL)
	defer db.Close()
	if err != nil {
		fmt.Println("PGSQL error:", err)
		return
	}

	table := jira.IssueTable{DB: db}
	tableRaw := jira.IssueRawTable{DB: db}
	total, err := tableRaw.Count()
	if err != nil {
		fmt.Println("Count error:" ,err)
		return
	}

	bar := pb.StartNew(total)
	defer bar.Finish()
	bar.SetTemplate(pb.Full)
	bar.Set("prefix", "[Indexing issues] ")

	rows, err := tableRaw.CursorAll()
	if err != nil {
		fmt.Println("Cursor error:", err)
		return
	}
	for rows.Next() {
		var (
			tmp jira.JiraIssue
			raw jira.IssueRaw
			iss jira.Issue
		)

		if err = rows.StructScan(&raw); err != nil {
			fmt.Println("Scan error", err)
			return
		}

		if tmp, err = jira.JiraIssueFromJson([]byte(raw.Raw)); err != nil {
			fmt.Println("From json error:", err)
			return
		}

		iss = jira.JiraToIssue(tmp)
		if err = table.Insert(iss); err != nil {
			fmt.Println("Insert error:", err)
			return
		}

		bar.Increment()
	}
}