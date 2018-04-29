package main

import "github.com/aws/aws-lambda-go/lambda"

func handler() (string, error) {
	return "Welcome to Serverless world", nil
}

func main() {
	lambda.Start(handler)
}
