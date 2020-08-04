package main

import (
	"github.com/dtext/cloudwatch-scheduler/cloudwatch"
	"github.com/dtext/cloudwatch-scheduler/scheduling"
	"github.com/go-chi/chi"
)

type editor struct {
	service Service
}

func newEditor(schedulerRule string, taskTable string) editor {
	return editor{
		service: Service{
			repo:      scheduling.NewTaskRepository(taskTable),
			scheduler: cloudwatch.Client(schedulerRule),
		},
	}
}

func (s editor) configureRoutes(r *chi.Mux) {
	r.Get("/tasks", listTasks(s.service))
	r.Post("/tasks", createTask(s.service))
	r.Delete("/tasks/{id}", deleteTask(s.service))
}
