.DEFAULT_GOAL = build

.PHONY: build
build:
	go build -race -v -o aoc cli/main.go

.PHONY: test
test:
	go test -race -v ./aoc2019 ./cli ./utils