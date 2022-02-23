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


	startAt := 0
	maxResults := 10

	res, err := cli.IssueByProject(PRJKEY, startAt, maxResults)
	for err == nil && len(res.Issues) > 0 {
		list, err := res.IssueList()
		if err != nil {
			fmt.Println("Issue list error:", err)
			return
		}
		fmt.Println(list)
		startAt += maxResults
		res, err = cli.IssueByProject(PRJKEY, startAt, maxResults)
	}

	if err != nil {
		fmt.Println("Issue error:", err)
	}
}
