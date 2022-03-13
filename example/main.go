package main

import (
	"fmt"
	"github.com/suntoucha/jira"
)

var (
	HOST, USER, TOKEN string
	PRJKEY, ISSUEKEY  string
)

func main() {
	fmt.Println("Hello World")

	cli := jira.Client{Host: HOST, User: USER, Token: TOKEN}

	prj, err := cli.Project(PRJKEY)
	if err != nil {
		fmt.Println("Project error:", err)
		return
	}
	fmt.Printf("PROJECT: %#v\n", prj)

	iss, err := cli.Issue(ISSUEKEY)
	if err != nil {
		fmt.Println("Issue error:", err)
		return
	}
	fmt.Printf("ISSUE: %#v\n", iss)

	var (
		me MyExporter
	)
	cli.ExportIssueByProject(PRJKEY, 10, me)
}

type MyExporter struct {
}

func (x MyExporter) Export(list jira.IssueList, startAt int, maxResult int, total int) error {
	fmt.Printf("startAt %v, maxResult %v, total %v, issue-len %v\n", startAt, maxResult, total, len(list))
	return nil
}
