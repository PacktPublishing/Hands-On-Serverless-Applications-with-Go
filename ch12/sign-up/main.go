package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/aws"
)

func main() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatal(err)
	}

	cognito := cognitoidentityprovider.New(cfg)
	req := cognito.SignUpRequest(&cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(os.Getenv("COGNITO_CLIENT_ID")),
		Username: aws.String("EMAIL"),
		Password: aws.String("PASSWORD"),
	})
	_, err = req.Send()
	if err != nil {
		log.Fatal(err)
	}
}
