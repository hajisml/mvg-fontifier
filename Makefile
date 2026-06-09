# Makefile for mvg-fontifier

# Build configuration
BINARY_NAME=mvg
BUILD_DIR=bin
MAIN_PACKAGE=./cmd/mvg-cli

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

.PHONY: all build test clean help tidy

all: test build

build: tidy ## Build the binary
	mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)

test: ## Run unit tests
	$(GOTEST) -v ./...

clean: ## Remove build artifacts
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

tidy: ## Tidy go modules
	$(GOMOD) tidy

run: build ## Build and run the TUI
	./$(BUILD_DIR)/$(BINARY_NAME)

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
