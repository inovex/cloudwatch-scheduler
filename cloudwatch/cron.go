package cloudwatch

import (
	"strconv"
	"strings"
	"time"
)

type cronExpression struct {
	Minute     string
	Hour       string
	DayOfMonth string
	Month      string
	DayOfWeek  string
	Year       string
}

// onlyAt returns the time expression of a cronjob that runs once a year, exactly at the specified date and time.
// The year given within t is ignored.
func onlyAt(t time.Time) cronExpression {
	year, month, dayOfMonth := t.Date()
	return cronExpression{
		Minute:     strconv.Itoa(t.Minute()),
		Hour:       strconv.Itoa(t.Hour()),
		DayOfMonth: strconv.Itoa(dayOfMonth),
		Month:      strconv.Itoa(int(month)),
		DayOfWeek:  "?",
		Year:       strconv.Itoa(year),
	}
}

// String returns the string representation of the cron expression
func (e cronExpression) String() string {
	return strings.Join([]string{
		e.Minute,
		e.Hour,
		e.DayOfMonth,
		e.Month,
		e.DayOfWeek,
		e.Year,
	}, " ")
}
