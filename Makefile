.PHONY: run
run:
	go run ./cmd/app/main.go

.PHONY: build
build:
	go build -v ./cmd/app/

.DEFAULT_GOAL := build