package jira

import (
	"strings"
	"time"
)

const (
	DATETIME_LAYOUT = "2006-01-02T15:04:05.999999999-0700"
)

type Datetime struct {
	time.Time
}

func (dt *Datetime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		return nil
	}

	var err error
	if dt.Time, err = time.Parse(DATETIME_LAYOUT, s); err != nil {
		return err
	}

	return nil
}

func NewDatetime(s string) (Datetime, error) {
	var (
		tmp Datetime
		err error
	)

	if tmp.Time, err = time.Parse(DATETIME_LAYOUT, s); err != nil {
		return Datetime{}, err
	}

	return tmp, nil
}

func NewDatetimeMustCompile(s string) Datetime {
	tmp, _ := NewDatetime(s)

	return tmp
}

type DatetimeIgnore struct {
	time.Time
}

func (dt *DatetimeIgnore) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		return nil
	}

	dt.Time, _ = time.Parse(DATETIME_LAYOUT, s)
	return nil
}

func NewDatetimeIgnore(s string) DatetimeIgnore {
	var (
		tmp DatetimeIgnore
	)

	tmp.Time, _ = time.Parse(DATETIME_LAYOUT, s)
	return tmp
}
