#!/bin/bash

GOOS=linux go build -o main main.go
zip deployment.zip main
aws lambda create-function --function-name ReverseString \
    --role arn:aws:iam::305929695733:role/service-role/lambda-role-execute --handler main \
    --runtime go1.x --zip-file fileb://./deployment.zip
rm main deployment.zip