package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Movie struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var movies = []Movie{
	Movie{
		ID:   1,
		Name: "Avengers",
	},
	Movie{
		ID:   2,
		Name: "Ant-Man",
	},
	Movie{
		ID:   3,
		Name: "Thor",
	},
	Movie{
		ID:   4,
		Name: "Hulk",
	},
	Movie{
		ID:   5,
		Name: "Doctor Strange",
	},
}

func insert(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var movie Movie
	err := json.Unmarshal([]byte(req.Body), &movie)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Invalid payload",
		}, nil
	}

	movies = append(movies, movie)

	response, err := json.Marshal(movies)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(response),
	}, nil
}

func main() {
	lambda.Start(insert)
}
