package jira

import (
	"errors"
	"regexp"
)

var (
	reOnlyLetters = regexp.MustCompile("[^a-zA-Z0-9\\-]")
)

func IsValidKey(key string) error {
	if len(key) == 0 {
		return errors.New("Key is empty")
	}

	if len(key) > 20 {
		return errors.New("Key is too long")
	}

	if reOnlyLetters.MatchString(key) {
		return errors.New("Key is not valid")
	}

	return nil
}
