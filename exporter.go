package jira

type Exporter interface {
	Export(list []Issue, startAt int, maxResult int, total int) error
}
