package jira

type Exporter interface {
	Export(list IssueList, startAt int, maxResult int, total int) error
}
