BINARY := bin/hexlet-path-size

.PHONY: build run test lint lint-fix check

build:
	go build -o $(BINARY) ./cmd/hexlet-path-size

run: build
	./$(BINARY) $(ARGS)

test:
	go test -v ./...

lint:
	golangci-lint run ./...

lint-fix:
	golangci-lint run --fix ./...

check: lint test
