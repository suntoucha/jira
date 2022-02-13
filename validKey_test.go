package jira

import (
	"testing"
)

func TestIsValidKey(t *testing.T) {
	testcase := []struct {
		Input  string
		Result bool
	}{
		{Input: "APPS", Result: true},
		{Input: "APPS-22", Result: true},
		{Input: "APPS;22", Result: false},
		{Input: "APPS-{22}", Result: false},
		{Input: "", Result: false},
		{Input: "12345678901234567890", Result: true},
		{Input: "123456789012345678901", Result: false},
	}

	for _, x := range testcase {
		err := IsValidKey(x.Input)
		if x.Result && err != nil {
			t.Errorf("Key [%v] must be valid but it's not. Error: %v", x.Input, err)
		} else if !x.Result && err == nil {
			t.Errorf("Key [%v] must be NOT valid, but no error", x.Input)
		}
	}
}
