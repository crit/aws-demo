BUILD := $(shell date -u +%m%d%H%M)

default: run

.PHONY: default run

run:
				./env

test:
				go test ./...

install:
				go build -o bin/darwin/aws-demo -ldflags "-X main.BUILD $(BUILD)" .
				GOOS=linux GOARCH=amd64 go build -o bin/linux/aws-demo -ldflags "-X main.BUILD $(BUILD)" .

dependency:
				go get ./...
