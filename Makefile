BINARY_NAME := awesomeai
GO_VERSION := 1.21.4
VERSION := 1.0.0
BUILD_DIR := build
PLATFORMS := linux windows darwin
ARCHITECTURES := amd64 arm64

all: build

build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) .

run: build
	@echo "Running $(BINARY_NAME) with MODEL_DATA_URL=$(MODEL_DATA_URL)..."
	@MODEL_DATA_URL=$(MODEL_DATA_URL) ./$(BUILD_DIR)/$(BINARY_NAME)

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)

build-all: clean
	@echo "Building for all platforms..."
	@for GOOS in $(PLATFORMS); do \
		for GOARCH in $(ARCHITECTURES); do \
			export GOOS=$$GOOS GOARCH=$$GOARCH; \
			OUTPUT=$(BUILD_DIR)/$(BINARY_NAME)-$$GOOS-$$GOARCH; \
			if [ "$$GOOS" = "windows" ]; then OUTPUT=$$OUTPUT.exe; fi; \
			echo "Building $$OUTPUT..."; \
			go build -o $$OUTPUT .; \
		done \
	done

deps:
	@echo "Checking dependencies..."
	@go mod download

check-go:
	@echo "Checking Go version..."
	@go version | grep -q "go$(GO_VERSION)" || (echo "Go version $(GO_VERSION) required" && exit 1)

fmt:
	@echo "Formatting code..."
	@go fmt ./...

lint:
	@echo "Linting code..."
	@golangci-lint run

test:
	@echo "Running tests..."
	@go test -v ./...

help:
	@echo "Available targets:"
	@echo "  all       - Build for current platform (default)"
	@echo "  build     - Build for current platform"
	@echo "  build-all - Build for all platforms"
	@echo "  run       - Build and run"
	@echo "  clean     - Clean build artifacts"
	@echo "  deps      - Install dependencies"
	@echo "  check-go  - Check Go version"
	@echo "  fmt       - Format code"
	@echo "  lint      - Lint code"
	@echo "  test      - Run tests"
	@echo "  help      - Show this help"

.PHONY: all build build-all run clean deps check-go fmt lint test help
