package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func handler(input string) (string, error) {
	log.Println("Before:", input)
	output := reverse(input)
	log.Println("After:", output)
	return output, nil
}

func main() {
	lambda.Start(handler)
}
