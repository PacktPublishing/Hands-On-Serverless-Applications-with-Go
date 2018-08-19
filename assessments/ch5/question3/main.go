package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Movie struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
	case http.MethodGet:
		// get all movies handler
		break
	case http.MethodPost:
		// insert movie handler
		break
	case http.MethodDelete:
		// delete movie handler
		break
	case http.MethodPut:
		// update movie handler
		break
	default:
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Body:       "Unsupported HTTP method",
		}, nil
	}
}

func main() {
	lambda.Start(handler)
}
