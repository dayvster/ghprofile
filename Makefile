.PHONY: build install clean test help

VERSION ?= dev
LDFLAGS := -ldflags "-X main.Version=$(VERSION)"

help:
	@echo "ghprofile build targets:"
	@echo "  make build          Build ghprofile binary"
	@echo "  make install        Build and install to $$GOPATH/bin"
	@echo "  make test           Run tests"
	@echo "  make clean          Remove build artifacts"
	@echo "  make fmt            Format code"
	@echo ""
	@echo "Customize with: make build VERSION=1.0.0"

build:
	go build -o ghprofile ./cmd/main.go

install: build
	go install ./cmd/main.go

test:
	go test -v ./...

clean:
	rm -f ghprofile

fmt:
	go fmt ./...

.DEFAULT_GOAL := help
