package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func signUp(account Account) error {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return err
	}

	cognito := cognitoidentityprovider.New(cfg)
	req := cognito.SignUpRequest(&cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(os.Getenv("COGNITO_CLIENT_ID")),
		Username: aws.String(account.Username),
		Password: aws.String(account.Password),
	})
	_, err = req.Send()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	lambda.Start(signUp)
}
