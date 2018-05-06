package main

import "github.com/aws/aws-lambda-go/lambda"

type Response struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
}

func handler() (Response, error) {
	return Response{
		StatusCode: 200,
		Body:       "Welcome to Serverless world",
	}, nil
}

func main() {
	lambda.Start(handler)
}
