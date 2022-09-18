.PHONY: setup

.PHONY: mod
mod:
	go mod tidy

.PHONY: build
build:
	go build \
			--trimpath \
			-o coinconv \
			cmd/cli/main.go

.PHONY: check
check:
	golangci-lint run -v --config .golangci.yml
