package scheduling

import (
	"github.com/go-chi/chi"
	"task-editor/cloudwatch"
)

type SchedulerEditor struct {
	service Service
}

func NewEditor(schedulerRule string, taskTable string) SchedulerEditor {
	return SchedulerEditor{
		service: Service{
			repo:      newDynamoRepo(taskTable),
			scheduler: cloudwatch.Client(schedulerRule),
		},
	}
}

func (s SchedulerEditor) ConfigureRoutes(r *chi.Mux) {
	r.Get("/tasks", listTasks(s.service))
	r.Post("/tasks", createTask(s.service))
	r.Delete("/tasks/{id}", deleteTask(s.service))
}
