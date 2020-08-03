package models

import "time"

type Task struct {
	ID      string                 `dynamo:"ID" json:"id"`
	Due     time.Time              `dynamo:"Due" json:"due"`
	Action  string                 `dynamo:"Action" json:"action"`
	Payload map[string]interface{} `dynamo:"Payload" json:"payload"`
}
