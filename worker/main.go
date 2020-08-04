package main

import (
	"fmt"
	"github.com/dtext/cloudwatch-scheduler/cloudwatch"
	"github.com/dtext/cloudwatch-scheduler/scheduling"
)

func main() {
	p := processor{
		items:     ItemService{},
		scheduler: cloudwatch.Client("cwscheduler-task-worker_schedule"),
		tasks:     scheduling.NewTaskRepository("cwscheduler-jobs"),
	}

	if err := p.processTasks(); err != nil {
		fmt.Println("error before or after processing tasks: ", err)
	}
}
