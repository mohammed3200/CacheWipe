.PHONY: help build run test clean fmt lint coverage

help:
	@echo "Cache Wipe - Build Commands"
	@echo "============================"
	@echo "  make build    - Build the application"
	@echo "  make run      - Run the application"
	@echo "  make test     - Run tests"
	@echo "  make clean    - Clean build files"
	@echo "  make fmt      - Format code"
	@echo "  make lint     - Run linter"
	@echo "  make coverage - Generate coverage report"

build:
	go build -o cache-wipe

run:
	go run main.go

test:
	go test -v ./...

clean:
	go clean
	rm -f cache-wipe cache-wipe.exe

fmt:
	go fmt ./...

lint:
	golangci-lint run

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html