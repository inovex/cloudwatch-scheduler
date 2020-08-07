// SPDX-FileCopyrightText: 2020 inovex GmbH <https://www.inovex.de>
//
// SPDX-License-Identifier: MIT
package main

import (
	"fmt"
	"github.com/dtext/cloudwatch-scheduler/cloudwatch"
	"github.com/dtext/cloudwatch-scheduler/scheduling"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

const fakeNow = "fake-now"

func main() {
	app := cli.NewApp()
	app.Name = "worker"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name: fakeNow,
			Usage: "Fakes the current date and time for easier testing." +
				"The default value 'now' leads to the actual current system time being used.",
			Value: "now",
		},
	}
	app.Action = run

	if err := app.Run(os.Args); err != nil {
		fmt.Println("error before or after processing tasks: ", err)
	}
}

func run(ctx *cli.Context) error {
	t := parseTimeArg(ctx.String(fakeNow))
	p := processor{
		upTo:      t,
		items:     ItemService{},
		scheduler: cloudwatch.Client("cwscheduler-task-worker_schedule"),
		tasks:     scheduling.NewTaskRepository("cwscheduler-jobs"),
	}
	return p.processTasks()
}

func parseTimeArg(arg string) time.Time {
	t, err := time.Parse(time.RFC3339, arg)
	if err == nil {
		return t
	} else if arg != "now" {
		fmt.Println("error parsing provided time argument, using current time instead")
	}
	return time.Now()

}
