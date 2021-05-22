#!/bin/sh

# 1. service dir
SERVICE_DIR="$PWD"
echo "LOGGER SERVICE_DIR: $SERVICE_DIR"

# 2. build
echo "begin clean history resource"
rm -rf AmService
echo "clean history resource success"

echo "begin build service"
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
go build -o AmService ./src
echo "build service success"

# 3. docker build
echo "build docker build"
docker build -t fuhui/golang-am:2.0 .
echo "docker build success"
