// SPDX-FileCopyrightText: 2020 inovex GmbH <https://www.inovex.de>
//
// SPDX-License-Identifier: MIT
package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/urfave/cli/v2"
	"net/http"
	"os"
)

const (
	awsAccessKey    = "aws-access-key"
	awsAccessKeyEnv = "AWS_ACCESS_KEY_ID"
	awsSecretKey    = "aws-secret-key"
	awsSecretKeyEnv = "AWS_SECRET_ACCESS_KEY"
	awsRegion       = "aws-region"
	awsRegionEnv    = "AWS_REGION"
)

// buildCLI creates the cli app for this control service. The app parses program arguments and passes
// them to the action functions in the form of a context object.
func buildCLI() *cli.App {
	app := cli.NewApp()
	app.Name = "example-service.go"
	flagsForAllCommands := []cli.Flag{
		&cli.StringFlag{
			Name:     awsAccessKey,
			EnvVars:  []string{awsAccessKeyEnv},
			Required: true,
			Usage:    "the access key of an AWS user",
			FilePath: ".aws.access",
		},
		&cli.StringFlag{
			Name:     awsSecretKey,
			EnvVars:  []string{awsSecretKeyEnv},
			Required: true,
			Usage:    "the secret key belonging to the access key",
			FilePath: ".aws.secret",
		},
		&cli.StringFlag{
			Name:    awsRegion,
			EnvVars: []string{awsRegionEnv},
			Value:   "eu-central-1",
			Usage:   "the AWS region where all resources are located",
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:   "start-server",
			Usage:  "starts the server to accept tasks via the REST API",
			Action: startServer,
			Flags:  flagsForAllCommands,
		},
	}
	return app
}

// setAWSEnv sets some environment variables, just in case their values are obtained from a file.
// These variables are required for the AWS SDK.
func setAWSEnv(c *cli.Context) {
	const errmsg = "could not set %s environment variable: "
	err := os.Setenv(awsRegionEnv, c.String(awsRegion))
	if err != nil {
		fmt.Println(fmt.Sprintf(errmsg, awsRegionEnv), err.Error())
	}
	err = os.Setenv(awsAccessKeyEnv, c.String(awsAccessKey))
	if err != nil {
		fmt.Print(fmt.Sprintf(errmsg, awsAccessKeyEnv), err.Error())
	}
	err = os.Setenv(awsSecretKeyEnv, c.String(awsSecretKey))
	if err != nil {
		fmt.Print(fmt.Sprintf(errmsg, awsSecretKeyEnv), err.Error())
	}
}

func main() {
	app := buildCLI()
	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error running app: ", err)
	}
}

func startServer(ctx *cli.Context) error {
	setAWSEnv(ctx)
	r := chi.NewRouter()
	e := newEditor("cwscheduler-task-worker_schedule", "cwscheduler-jobs")
	e.configureRoutes(r)
	return http.ListenAndServe(":8080", r)
}
