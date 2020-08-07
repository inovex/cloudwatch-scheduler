// SPDX-FileCopyrightText: 2020 inovex GmbH <https://www.inovex.de>
//
// SPDX-License-Identifier: MIT
package main

import (
	"fmt"
	"github.com/inovex/cloudwatch-scheduler/cloudwatch"
	"github.com/inovex/cloudwatch-scheduler/scheduling"
	"github.com/mitchellh/mapstructure"
	"sort"
	"time"
)

type processor struct {
	upTo      time.Time
	items     ItemService
	tasks     scheduling.TaskRepository
	scheduler cloudwatch.Scheduler
}

func (p processor) processTasks() error {
	// retrieve all tasks
	tasks, err := p.tasks.GetTasks()
	if err != nil {
		return err
	}

	// sort tasks by due date
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Due.Unix() < tasks[j].Due.Unix()
	})

	// process tasks
	for _, task := range tasks {
		// see if task is due
		if task.Due.Unix() > p.upTo.Unix() {
			// if it isn't, schedule it and exit
			return p.scheduler.Schedule(task.Due)
		}
		// if it is, process it
		err = p.processTask(task)
		if err != nil {
			// log processing errors for single tasks
			fmt.Println("error processing task ", task.ID, ": ", err)
			continue
		}
		// delete task from task queue
		err = p.tasks.Done(task)
	}
	// if all queued tasks have been processed, disable scheduler
	return p.scheduler.Unschedule()
}

func (p processor) processTask(t scheduling.Task) error {
	switch t.Action {
	case "APPLY_SALE":
		var a SaleAction
		err := mapstructure.Decode(t.Payload, &a)
		if err != nil {
			return err
		}
		p.items.ApplySale(a)
		return nil
	default:
		return fmt.Errorf("I don't know the action type %s\n", t.Action)
	}
}
