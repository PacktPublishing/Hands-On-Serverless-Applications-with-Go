package main

import (
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestFindAll_WithoutIAMRole(t *testing.T) {
	expected := events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       "Error while scanning DynamoDB",
	}
	response, err := findAll()
	assert.IsType(t, nil, err)
	assert.Equal(t, expected, response)
}

func TestFindAll_WithIAMRole(t *testing.T) {
	response, err := findAll()
	assert.IsType(t, nil, err)
	assert.NotNil(t, response.Body)
}
