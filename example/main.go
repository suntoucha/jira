package main

import (
	"fmt"
	"github.com/suntoucha/jira"
)

var (
	HOST, USER, TOKEN string
)

func main() {
	fmt.Println("Hello World")

	cli := jira.Client{Host: HOST, User: USER, Token: TOKEN}

	list, err := cli.ProjectList()
	if err != nil {
		fmt.Println("Project list error:", err)
		return
	}

	for _, x := range list {
		fmt.Printf("%#v\n\n", x)

		cur := x.Issue()
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
}
