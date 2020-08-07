// SPDX-FileCopyrightText: 2020 inovex GmbH <https://www.inovex.de>
//
// SPDX-License-Identifier: MIT
package main

import (
	"github.com/go-chi/chi"
	"github.com/inovex/cloudwatch-scheduler/cloudwatch"
	"github.com/inovex/cloudwatch-scheduler/scheduling"
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
