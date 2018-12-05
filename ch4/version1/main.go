package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	StatusCode      int               `json:"statusCode"`
	Headers         map[string]string `json:"headers"`
	Body            string            `json:"body"`
	IsBase64Encoded bool              `json:"isBase64Encoded,omitempty"`
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
