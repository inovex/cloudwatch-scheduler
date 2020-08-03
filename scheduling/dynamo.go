package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/guregu/dynamo"
	"task-editor/models"
)

type dynamoRepo struct {
	table dynamo.Table
}

func newDynamoRepo(taskTable string) dynamoRepo {
	db := dynamodb.New(session.Must(session.NewSessionWithOptions(session.Options{})))
	t := dynamo.NewFromIface(db).Table(taskTable)
	return dynamoRepo{table: t}
}

func (r dynamoRepo) GetTasks() ([]models.Task, error) {
	var out []models.Task
	err := r.table.
		Scan().
		All(&out)
	return out, err
}

func (r dynamoRepo) AddTask(task models.Task) error {
	return r.table.
		Put(task).
		Run()
}

func (r dynamoRepo) DeleteTask(id string) error {
	return r.table.
		Delete("ID", id).
		Run()
}

