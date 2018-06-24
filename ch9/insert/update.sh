#!/bin/bash

echo "Building binary"
GOOS=linux GOARCH=amd64 go build -o main main.go

echo "Building deployment package"
zip deployment.zip main

echo "Updating function's code"
aws lambda update-function-code --function-name InsertMovie --zip-file fileb://./deployment.zip

echo "Cleaning up"
rm main deployment.zip