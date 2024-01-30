.PHONY: build
build:
	go build -v

.PHONY: run
run:
	./apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build