.PHONY: build run clean test docker

APP_NAME := aqicloud-ai-agent-go

build:
	go build -o bin/$(APP_NAME) ./cmd/agent

run: build
	./bin/$(APP_NAME)

clean:
	rm -rf bin/
	go clean

test:
	go test ./... -v

docker:
	docker build -t $(APP_NAME) .

tidy:
	go mod tidy

fmt:
	gofmt -w .
