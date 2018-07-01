package main

import (
	"fmt"
	"log"

	"github.com/djhworld/go-lambda-invoke/golambdainvoke"
)

func main() {
	response, err := golambdainvoke.Run(8001, 9)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(response))
}
