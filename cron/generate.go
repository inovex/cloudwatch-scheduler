package cron

import (
	"strconv"
	"strings"
	"time"
)

type Expression struct {
	Minute     string
	Hour       string
	DayOfMonth string
	Month      string
	DayOfWeek  string
	Year       string
}

// OnlyAt returns the time expression of a cronjob that runs once a year, exactly at the specified date and time.
// The year given within t is ignored.
func OnlyAt(t time.Time) Expression {
	year, month, dayOfMonth := t.Date()
	return Expression{
		Minute:     strconv.Itoa(t.Minute()),
		Hour:       strconv.Itoa(t.Hour()),
		DayOfMonth: strconv.Itoa(dayOfMonth),
		Month:      strconv.Itoa(int(month)),
		DayOfWeek:  "?",
		Year:       strconv.Itoa(year),
	}
}

// String returns the string representation of the cron expression
func (e Expression) String() string {
	return strings.Join([]string{
		e.Minute,
		e.Hour,
		e.DayOfMonth,
		e.Month,
		e.DayOfWeek,
		e.Year,
	}, " ")
}
