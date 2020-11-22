.DEFAULT_GOAL = build

.PHONY: build
build:
	go build -race -v -o aoc cli/main.go