package main

import (
	"fmt"
	"task-editor/cloudwatch"
	"task-editor/scheduling"
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
