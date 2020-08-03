package cloudwatch

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
	"time"
)

type Scheduler struct {
	ruleName string
	cw       *cloudwatchevents.CloudWatchEvents
}

func Client(ruleName string) Scheduler {
	return Scheduler{
		ruleName: ruleName,
		cw:       cloudwatchevents.New(session.Must(session.NewSession())),
	}
}

// Schedule creates a scheduler tick at the specified time t.
// It takes the year, month, day, hour and minutes fields into account.
// Cloudwatch will execute the worker lambda function at that time.
func (client Scheduler) Schedule(t time.Time) error {
	schedule := fmt.Sprintf("cron(%s)", onlyAt(t))
	input := cloudwatchevents.PutRuleInput{
		Name:               aws.String(client.ruleName),
		ScheduleExpression: aws.String(schedule),
		State:              aws.String(cloudwatchevents.RuleStateEnabled),
	}
	_, err := client.cw.PutRule(&input)
	return err
}
