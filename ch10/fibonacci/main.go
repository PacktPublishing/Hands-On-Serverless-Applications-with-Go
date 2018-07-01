package main

import "github.com/aws/aws-lambda-go/lambda"

func fib(n int64) int64 {
	if n > 2 {
		return fib(n-1) + fib(n-2)
	}
	return 1
}

func handler(n int64) (int64, error) {
	return fib(n), nil
}

func main() {
	lambda.Start(handler)
}
