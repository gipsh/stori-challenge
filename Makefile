SHELL := /bin/bash

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

build:
	go build -o stori-challenge cmd/local/main.go

lambda: 
	GOOS=linux go build -o stori-challenge-lambda cmd/lambda/main.go
	
test:
	go test -v -cover ./...

run:
	go run main.go


.PHONY: build test run