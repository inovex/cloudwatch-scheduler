package scheduling

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
	"task-editor/cron"
	"time"
)

type cwClient struct {
	ruleName string
	cw *cloudwatchevents.CloudWatchEvents
}

func newCloudwatchClient(ruleName string) cwClient {
	return cwClient{
		ruleName: ruleName,
		cw: cloudwatchevents.New(session.Must(session.NewSession())),
	}
}

func (client cwClient) Schedule(t time.Time) error {
	schedule := fmt.Sprintf("cron(%s)", cron.OnlyAt(t))
	input := cloudwatchevents.PutRuleInput{
		Name:               aws.String(client.ruleName),
		ScheduleExpression: aws.String(schedule),
		State:              aws.String(cloudwatchevents.RuleStateEnabled),
	}
	_, err := client.cw.PutRule(&input)
	return err
}
