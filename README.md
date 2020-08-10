# cloudwatch-scheduler
Example implementation for a serverless scheduling solution described in the [blog article](https://www.inovex.de/blog/schedule-aws-lambda) "Schedule AWS Lambda Invocations: How to Build Slow Schedulers".

## What it does
This is a *very slow* scheduler implementation based on a classical serverless AWS stack:
EventBridge, Cloudwatch, Lambda and Dynamo.
Its intended purpose is to schedule infrequent events far into the future and have them processed by a serverless function.
A good practical use case for this is the publishing and withdrawal of certain media in a content-drive platform.
Sometimes, the publishing and withdrawal dates of certain items are known months (if not years) in advance due to license agreements.
A scheduler implementation like this would allow users to automate tasks like these while efficiently working around certain limitations imposed by EventBridge.

## How to build it and use it
This is a monorepo containing two build targets.
The `editor/` package contains the sources for a task editor that is used to submit tasks to the scheduler.
The code for the worker lambda function that is triggered by Cloudwatch is located in `worker/`.
Shared code is within the other packages that can be found in the repository root.
There are Makefiles for both targets in their respective directories.

I included some basic terraform code for the worker lambda function and for the Cloudwatch/EventBridge rule in the `terraform/` directory.
You can easily build everything and deploy it to your AWS account as-is to play around with it.
However, if you end up using (some of) this code, you should probably read the [article](https://www.inovex.de/blog/schedule-aws-lambda) and get a more conceptual understanding of how everything comes together.

### Editor
To create a binary:
```bash
cd editor
make build
```
This will create an `editor` binary which you can run:
```bash
./editor start-server
```
You can put `.aws.access` and `.aws.secret` files in the working directory or have your access credentials set up in `~/.aws`.
Make sure the AWS user has write access to DynamoDB and EventBridge.

#### API
The command will start a web server on port 8080 that offers a simple REST API to manipulate our task queue.

List Tasks:
```bash
curl --request GET \
  --url http://localhost:8080/tasks
```

Create new or overwrite existing task:
```bash
curl --request POST \
  --url http://localhost:8080/tasks \
  --header 'content-type: application/json' \
  --data '{
    "id": "some task id",
    "due": "2020-12-31T00:00:00Z",
    "action": "APPLY_SALE",
    "payload": {
      "itemID": "some item",
      "newPrice": 29.99,
      "salePercent": 50
    }
  }'
```

### Worker
You can create a binary of the worker just as easily:
```bash
cd worker
make build
```
The resulting `main` binary just processes all due tasks.
You can pass an additional `fakenow` argument to mock a call at a specific time.
The time needs to be formatted the same way as before when submitting a new task.
```bash
./main --fakenow 2020-12-31T00:00:00Z
```
You can also `make zip` to create a deployable `function.zip` in the `terraform/` directory.

## How not to use it (POC Warning)
This implementation is a proof of concept rather than an actual finished library you could use in your application.
The code contains around a handful of issues that I am aware of, and probably some more that I'm not.
I would advise against using this code without thorough review, and I am in no way responsible for the damage it does to your production system.
Use caution and carefully test everything.
