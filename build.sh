#!/bin/bash
codedir='/usr/src/codetest'
echo ">>> Compling code"
docker run --rm -e GOPATH=/usr/ -v "$PWD":"$codedir" -w "$codedir" golang:1.9 sh -c "env && go get -v ./... && GOOS=linux go build -tags netgo -installsuffix netgo -v"
echo ">>> Building docker container"
docker build -t codetest .
