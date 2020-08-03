# cloudwatch-scheduler
Reference implementation for a serverless scheduling solution described in my [blog article](https://www.inovex.de/blog/) "Building Slow Schedulers on AWS".
// TODO link

## What it does
This is a *very slow* scheduler implementation based on a classical serverless AWS stack:
EventBridge, Cloudwatch, Lambda and Dynamo.
Its intended purpose is to schedule infrequent events far into the future and have them processed by a serverless function.
A good practical use case for this is the publishing and withdrawal of certain media in a content-drive platform.
Sometimes, the publishing and withdrawal dates of certain items are known months (if not years) in advance due to license agreements.
A scheduler implementation like this would allow users to automate tasks like these while efficiently working around certain limitations imposed by EventBridge.

## How to build it
This is a monorepo containing two build targets.
The `editor/` package contains the sources for a task editor that is used to submit tasks to the scheduler.
The code for the worker lambda function that is triggered by Cloudwatch is located in `worker/`.
Shared code is within the other packages that can be found in the repository root.

I included some basic terraform code for the worker lambda function and for the Cloudwatch/EventBridge rule in the `terraform/` directory.
You can easily build everything and deploy it to your AWS account as-is to play around with it.
However, if you end up using (some of) this code (maybe even in production),
you should probably read my [blog article](https://www.inovex.de/blog/) and get a more conceptual understanding of how everything comes together.
// TODO link

### Editor
TODO

### Worker
TODO

## How to use it
TODO

## How not to use it (POC Warning)
This implementation is designed as a proof of concept rather than an actual finished library you could use in your application.
The code contains around a handful of issues that I am aware of, and probably some more that I'm not.
I would advise against using this code without thorough review, and I am in no way responsible for the damage it does to your production system.
Use caution and carefully test everything.
