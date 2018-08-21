package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-lambda-go/events"
)

func TestUpdate_InvalidPayLoad(t *testing.T) {
	input := events.APIGatewayProxyRequest{
		Body: "{'name': 'avengers'}",
	}
	expected := events.APIGatewayProxyResponse{
		StatusCode: 400,
		Body:       "Invalid payload",
	}
	response, _ := update(input)
	assert.Equal(t, expected, response)
}

func TestUpdate_ValidPayload(t *testing.T) {
	input := events.APIGatewayProxyRequest{
		Body: "{\"id\":\"40\", \"name\":\"Thor\", \"description\":\"Marvel movie\", \"cover\":\"poster url\"}",
	}
	expected := events.APIGatewayProxyResponse{
		Body:       "{\"id\":\"40\", \"name\":\"Thor\", \"description\":\"Marvel movie\", \"cover\":\"poster url\"}",
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
	}
	response, _ := update(input)
	assert.Equal(t, expected, response)
}
