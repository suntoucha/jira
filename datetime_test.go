package jira

import (
	"encoding/json"
	"testing"
)

func TestDatetime(t *testing.T) {
	str := `
		{
			"dt":"2018-07-16T14:19:48.041+0300"	
		}`

	var tmp struct {
		Dt Datetime `json:"dt"`
	}

	if err := json.Unmarshal([]byte(str), &tmp); err != nil {
		t.Errorf("Datetime unmarshal error: %v", err)
		return
	}
}
