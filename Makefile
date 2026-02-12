BINARY := vikunja
VERSION ?= dev
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS := -ldflags "-X github.com/jo-nike/vikunja-cli/cmd.Version=$(VERSION) -X github.com/jo-nike/vikunja-cli/cmd.Commit=$(COMMIT) -X github.com/jo-nike/vikunja-cli/cmd.Date=$(DATE)"

.PHONY: build clean test

build:
	go build $(LDFLAGS) -o $(BINARY) .

clean:
	rm -f $(BINARY)

test:
	go test ./...
