.PHONY: install build run test

install:
	npx --yes skills add . --global --skill '*' --agent universal --yes

build:
	go build -o bin/workee ./cmd/workee

run:
	go run ./cmd/workee

test:
	go test ./...
