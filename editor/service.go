// SPDX-FileCopyrightText: 2020 inovex GmbH <https://www.inovex.de>
//
// SPDX-License-Identifier: MIT
package main

import (
	"github.com/inovex/cloudwatch-scheduler/scheduling"
	"time"
)

type Repository interface {
	GetTasks() ([]scheduling.Task, error)
	AddTask(scheduling.Task) error
	DeleteTask(id string) error
}

type WorkerScheduler interface {
	Schedule(time.Time) error
}

type Service struct {
	repo      Repository
	scheduler WorkerScheduler
}

func (s Service) ListTasks() ([]scheduling.Task, error) {
	return s.repo.GetTasks()
}

func isFirstTask(task scheduling.Task, allTasks []scheduling.Task) bool {
	for _, existing := range allTasks {
		if existing.Due.Unix() < task.Due.Unix() {
			return false
		}
	}
	return true
}

func (s Service) AddTask(task scheduling.Task) error {
	// get all tasks or return error
	tasks, err := s.repo.GetTasks()
	if err != nil {
		return err
	}
	// schedule task if new task is the first one to be executed
	if isFirstTask(task, tasks) {
		err = s.scheduler.Schedule(task.Due)
		// return error if scheduling fails
		if err != nil {
			return err
		}
	}
	// add new task to queue
	return s.repo.AddTask(task)
}

func (s Service) DeleteTask(id string) error {
	// We can ignore the schedule here because the worker will update it.
	// If the given task was the first one, the worker will just do nothing.
	return s.repo.DeleteTask(id)
}
