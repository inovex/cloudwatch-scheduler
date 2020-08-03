package scheduling

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/guregu/dynamo"
)

type TaskRepository struct {
	table dynamo.Table
}

func NewTaskRepository(taskTable string) TaskRepository {
	db := dynamodb.New(session.Must(session.NewSessionWithOptions(session.Options{})))
	t := dynamo.NewFromIface(db).Table(taskTable)
	return TaskRepository{table: t}
}

func (r TaskRepository) GetTasks() ([]Task, error) {
	var out []Task
	err := r.table.
		Scan().
		All(&out)
	return out, err
}

func (r TaskRepository) AddTask(task Task) error {
	return r.table.
		Put(task).
		Run()
}

func (r TaskRepository) DeleteTask(id string) error {
	return r.table.
		Delete("ID", id).
		Run()
}
