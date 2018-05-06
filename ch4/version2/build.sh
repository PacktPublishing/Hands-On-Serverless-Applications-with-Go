#!/bin/bash

echo "Build the binary"
GOOS=linux GOARCH=amd64 go build -o main main.go

echo "Create a zip file"
zip deployment.zip main

echo "Cleaning up"
rm main