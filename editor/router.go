package main

import (
	"github.com/go-chi/chi"
	"task-editor/cloudwatch"
	"task-editor/scheduling"
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
