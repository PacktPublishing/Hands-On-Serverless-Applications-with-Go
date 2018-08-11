package main

import "github.com/aws/aws-lambda-go/lambda"

func handler() (string, error) {
	return "First Lambda function with Terraform", nil
}
func main() {
	lambda.Start(handler)
}
