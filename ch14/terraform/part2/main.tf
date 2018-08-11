provider "aws" {
  region = "eu-west-1"
}

resource "aws_iam_role" "role" {
  name               = "FindAllMoviesRole"
  assume_role_policy = "${file("assume-role-policy.json")}"
}

resource "aws_iam_policy" "cloudwatch_policy" {
  name   = "PushCloudWatchLogsPolicy"
  policy = "${file("cloudwatch-policy.json")}"
}

resource "aws_iam_policy" "dynamodb_policy" {
  name   = "ScanDynamoDBPolicy"
  policy = "${file("dynamodb-policy.json")}"
}

resource "aws_iam_policy_attachment" "cloudwatch-attachment" {
  name       = "cloudwatch-lambda-attchment"
  roles      = ["${aws_iam_role.role.name}"]
  policy_arn = "${aws_iam_policy.cloudwatch_policy.arn}"
}

resource "aws_iam_policy_attachment" "dynamodb-attachment" {
  name       = "dynamodb-lambda-attchment"
  roles      = ["${aws_iam_role.role.name}"]
  policy_arn = "${aws_iam_policy.dynamodb_policy.arn}"
}

resource "aws_lambda_function" "findall" {
  function_name = "FindAllMovies"
  handler       = "main"
  filename      = "function/deployment.zip"
  runtime       = "go1.x"
  role          = "${aws_iam_role.role.arn}"

  environment {
    variables {
      TABLE_NAME = "movies"
    }
  }
}

resource "aws_dynamodb_table" "movies" {
  name           = "movies"
  read_capacity  = 5
  write_capacity = 5
  hash_key       = "ID"

  attribute {
    name = "ID"
    type = "S"
  }
}

resource "aws_dynamodb_table_item" "items" {
  table_name = "${aws_dynamodb_table.movies.name}"
  hash_key   = "${aws_dynamodb_table.movies.hash_key}"
  item       = "${file("movie.json")}"
}

resource "aws_api_gateway_rest_api" "api" {
  name = "MoviesAPI"
}

resource "aws_api_gateway_resource" "proxy" {
  rest_api_id = "${aws_api_gateway_rest_api.api.id}"
  parent_id   = "${aws_api_gateway_rest_api.api.root_resource_id}"
  path_part   = "movies"
}

resource "aws_api_gateway_method" "proxy" {
  rest_api_id   = "${aws_api_gateway_rest_api.api.id}"
  resource_id   = "${aws_api_gateway_resource.proxy.id}"
  http_method   = "GET"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "lambda" {
  rest_api_id = "${aws_api_gateway_rest_api.api.id}"
  resource_id = "${aws_api_gateway_method.proxy.resource_id}"
  http_method = "${aws_api_gateway_method.proxy.http_method}"

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = "${aws_lambda_function.findall.invoke_arn}"
}

resource "aws_api_gateway_deployment" "staging" {
  depends_on = ["aws_api_gateway_integration.lambda"]

  rest_api_id = "${aws_api_gateway_rest_api.api.id}"
  stage_name  = "staging"
}

resource "aws_lambda_permission" "apigw" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = "${aws_lambda_function.findall.arn}"
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_deployment.staging.execution_arn}/*/*"
}
