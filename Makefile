SHELL := /bin/bash

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

build:
	go build -o stori-challenge cmd/stori/main.go

lambda: 
	GOARCH=amd64 GOOS=linux go build -o stori-challenge cmd/stori/main.go
	
test:
	go test -v -cover ./...

run:
	go run main.go


.PHONY: build test run