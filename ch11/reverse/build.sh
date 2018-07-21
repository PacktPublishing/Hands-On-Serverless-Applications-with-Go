#!/bin/bash

GOOS=linux go build -o main main.go
zip deployment.zip main
aws lambda create-function --function-name ReverseString \
    --role arn:aws:iam::ACCOUNT_ID:role/service-role/ReverseStringRole --handler main \
    --runtime go1.x --zip-file fileb://./deployment.zip
rm main deployment.zip