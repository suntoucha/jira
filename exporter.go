package jira

import (
	"encoding/json"
)

type Exporter interface {
	Export(raw json.RawMessage, index int, total int) error
}
