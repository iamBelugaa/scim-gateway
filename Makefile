BINARY_NAME := scim-gateway
MAIN_PACKAGE := ./cmd/scim-gateway/main.go
SERVICE_MODULE := github.com/iamBelugaa/scim-gateway

BUILD_DIR := ./dist
GOA_GEN_DIR := ./gen
BUILD_FLAGS := -v -ldflags="-s -w"

.PHONY: tidy deps install-goa fmt clean gen-goa build run all

## Tidy Go modules
tidy:
	@echo "Tidying Go modules..."
	@go mod tidy
	@echo "Go modules tidied."

## Download Go dependencies
deps:
	@echo "Downloading Go modules..."
	@go mod download
	@go mod verify
	@echo "Go modules downloaded and verified."

## Install Goa framework
install-goa:
	@echo "Installing Goa framework..."
	@go install goa.design/goa/v3/cmd/goa@latest
	@echo "Goa installation complete."

## Format codebase
fmt:
	@echo "Formatting Go code..."
	@go fmt ./...
	@echo "Formatting complete."

## Clean build and generated files
clean:
	@echo "Removing build artifacts..."
	@go clean
	@rm -rf $(BUILD_DIR)
	@rm -rf $(GOA_GEN_DIR)
	@echo "Clean complete."

## Generate Goa code from design definitions
gen-goa:
	@echo "Generating Goa code..."
	@goa gen $(SERVICE_MODULE)/internal/design
	@echo "Goa code generation complete."

## Build the binary
build: clean gen-goa
	@echo "Building $(BINARY_NAME) for $(shell go env GOOS)/$(shell go env GOARCH)..."
	GOOS=$(shell go env GOOS) GOARCH=$(shell go env GOARCH) go build $(BUILD_FLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

## Run the built service
run: build
	@echo "Running $(BINARY_NAME)..."
	@$(BUILD_DIR)/$(BINARY_NAME)

## Run full cycle: clean, build, run
all: run