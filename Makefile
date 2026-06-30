BINARY := vikunja-cli
SKILL_DIR := .claude/skills/vikunja-cli
VERSION ?= dev
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS := -ldflags "-X github.com/jo-nike/vikunja-cli/cmd.Version=$(VERSION) -X github.com/jo-nike/vikunja-cli/cmd.Commit=$(COMMIT) -X github.com/jo-nike/vikunja-cli/cmd.Date=$(DATE)"

.PHONY: build skill clean test

build:
	go build $(LDFLAGS) -o $(BINARY) .

skill: ## Build the host-arch (darwin/arm64) binary into the vendored skill's bin/
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 \
		go build $(LDFLAGS) -o "$(SKILL_DIR)/bin/$(BINARY)" .

clean:
	rm -f $(BINARY)
	rm -rf dist build

test:
	go test ./...
