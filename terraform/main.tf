provider "aws" {
  region  = "eu-central-1"
  version = "~> 2.29"
}

terraform {
  backend "s3" {
    region         = "eu-central-1"
    bucket         = "cwscheduler-tf-state"
    dynamodb_table = "cwscheduler-tf-locks"
    key            = "task-worker.tfstate"
  }
}

locals {
  project = "cwscheduler"
  service = "task-worker"
  base_name = "${local.project}-${local.service}"
  lambda_zip = "./function.zip"
}

# Lambda
resource "aws_lambda_function" "lambda" {
  function_name = "${local.base_name}_lambda"

  filename = local.lambda_zip

  handler = "main"
  runtime = "go1.x"

  role = aws_iam_role.lambda_exec.arn

  source_code_hash = filebase64sha256(local.lambda_zip)

  timeout = 20

  depends_on = [
    aws_iam_role_policy_attachment.lambda_logs,
    aws_iam_role_policy_attachment.dynamodb
  ]
}

resource "aws_iam_role" "lambda_exec" {
  name = "${local.base_name}_lambda_role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

# Logging
resource "aws_iam_role_policy_attachment" "lambda_logs" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = aws_iam_policy.lambda_logging.arn
}

resource "aws_iam_policy" "lambda_logging" {
  name        = "${local.base_name}_lambda_logging"
  path        = "/"
  description = "IAM policy for logging from a lambda"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": "arn:aws:logs:*:*:*",
      "Effect": "Allow"
    }
  ]
}
EOF
}

# Dynamo DB access
resource "aws_iam_role_policy_attachment" "dynamodb" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = aws_iam_policy.dynamodb.arn
}

resource "aws_iam_policy" "dynamodb" {
  name        = "${local.base_name}_dynamodb"
  path        = "/"
  description = "IAM policy for dynamo DB access"

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": [
                "dynamodb:*"
            ],
            "Effect": "Allow",
            "Resource": "*"
        }
    ]
}
EOF
}

# Dynamo table holding the jobs
resource "aws_dynamodb_table" "job_table" {
  name         = "${local.project}-jobs"
  billing_mode = "PAY_PER_REQUEST"
  lifecycle {
    # Prevent TF from changing capacities while autoscaling is active
    ignore_changes = [read_capacity, write_capacity]
  }

  hash_key  = "ID"

  attribute {
    name = "ID"
    type = "S"
  }
}

# Cron Trigger for worker
resource "aws_cloudwatch_event_rule" "schedule" {
  name                = "${local.base_name}_schedule"
  description         = "Runs the worker lambda when the next job needs to be executed"
  schedule_expression = "cron(* * 1 1 ? 1970)"
}

resource "aws_cloudwatch_event_target" "target" {
  rule      = aws_cloudwatch_event_rule.schedule.name
  target_id = "${local.base_name}_worker_target"
  arn       = aws_lambda_function.lambda.arn
}

resource "aws_lambda_permission" "permission" {
  statement_id  = "${local.base_name}_AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.schedule.arn
}


