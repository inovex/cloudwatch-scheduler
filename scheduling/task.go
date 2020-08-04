package scheduling

import "time"

type Task struct {
	ID      string                 `dynamo:"ID" json:"id"`
	Due     time.Time              `dynamo:"Due" json:"due"`
	Action  string                 `dynamo:"Action" json:"action"`
	Payload map[string]interface{} `dynamo:"Payload" json:"payload"`
}

// Valid checks the given task for validity.
// It will only accept tasks with a due date that is at least 1 Minute in the future.
func (t Task) Valid() bool {
	deltaNow := time.Until(t.Due)
	return deltaNow > time.Minute
}
