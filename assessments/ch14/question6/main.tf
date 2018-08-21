provider "aws" {
  region = "${var.aws_region}"
}

resource "aws_iam_role" "roles" {
  count              = "${length(var.functions)}"
  name               = "${element(var.functions, count.index)}Role"
  assume_role_policy = "${file("policies/assume-role-policy.json")}"
}

resource "aws_iam_policy" "policies" {
  count  = "${length(var.functions)}"
  name   = "${element(var.functions, count.index)}Policy"
  policy = "${file("policies/${element(var.functions, count.index)}-policy.json")}"
}

resource "aws_iam_policy_attachment" "policy-attachments" {
  count      = "${length(var.functions)}"
  name       = "${element(var.functions, count.index)}Attachment"
  roles      = ["${element(aws_iam_role.roles.*.name, count.index)}"]
  policy_arn = "${element(aws_iam_policy.policies.*.arn, count.index)}"
}

resource "aws_lambda_function" "functions" {
  count         = "${length(var.functions)}"
  function_name = "${element(var.functions, count.index)}"
  handler       = "main"
  filename      = "functions/${element(var.functions, count.index)}.zip"
  runtime       = "go1.x"
  role          = "${element(aws_iam_role.roles.*.arn, count.index)}"

  environment {
    variables {
      TABLE_NAME = "${var.table_name}"
    }
  }
}

resource "aws_dynamodb_table" "table" {
  name           = "${var.table_name}"
  read_capacity  = 5
  write_capacity = 5
  hash_key       = "ID"

  attribute {
    name = "ID"
    type = "S"
  }
}

resource "aws_api_gateway_rest_api" "api" {
  name = "MoviesAPI"
}

resource "aws_api_gateway_resource" "proxy" {
  rest_api_id = "${aws_api_gateway_rest_api.api.id}"
  parent_id   = "${aws_api_gateway_rest_api.api.root_resource_id}"
  path_part   = "movies"
}

resource "aws_api_gateway_deployment" "staging" {
  depends_on = ["aws_api_gateway_integration.integrations"]

  rest_api_id = "${aws_api_gateway_rest_api.api.id}"
  stage_name  = "staging"
}

resource "aws_api_gateway_method" "proxies" {
  count         = "${length(var.functions)}"
  rest_api_id   = "${aws_api_gateway_rest_api.api.id}"
  resource_id   = "${aws_api_gateway_resource.proxy.id}"
  http_method   = "${lookup(var.methods, element(var.functions, count.index))}"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "integrations" {
  count       = "${length(var.functions)}"
  rest_api_id = "${aws_api_gateway_rest_api.api.id}"
  resource_id = "${element(aws_api_gateway_method.proxies.*.resource_id, count.index)}"
  http_method = "${element(aws_api_gateway_method.proxies.*.http_method, count.index)}"

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = "${element(aws_lambda_function.functions.*.invoke_arn, count.index)}"
}

resource "aws_lambda_permission" "permissions" {
  count         = "${length(var.functions)}"
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = "${element(aws_lambda_function.functions.*.arn, count.index)}"
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_deployment.staging.execution_arn}/*/*"
}
