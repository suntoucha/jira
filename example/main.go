package main

import (
	"fmt"
	"github.com/suntoucha/jira"
)

var (
	HOST, USER, TOKEN string
	PRJKEY            string
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



		cur := prj.Issue()
		list, err := cli.IssueCursor(&cur)
		for len(list) > 0 && err == nil {
			for _, iss := range list {
				fmt.Printf("\t%#v\n", iss)
			}
			list, err = cli.IssueCursor(&cur)
		}
		if err != nil {
			fmt.Println("Issue error:", err)
		}
}
