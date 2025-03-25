# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet
GOGET=$(GOCMD) get
BINARY_NAME=myapp
BINARY_DIR=bin
BINARY_LINUX=$(BINARY_DIR)/$(BINARY_NAME)_linux
BINARY_WINDOWS=$(BINARY_DIR)/$(BINARY_NAME)_windows.exe
MAIN_SRC=cmd/server/main.go

# All target
all: test vet build

# Build the project
build:
	$(GOBUILD) -o $(BINARY_DIR)/$(BINARY_NAME) -v $(MAIN_SRC)

# Run tests
test:
	$(GOTEST) -v ./...

# Vet the code
vet:
	$(GOVET) ./...

# Clean build artifacts
clean:
	$(GOCLEAN)
	rm -rf $(BINARY_DIR)

# Run the application
run:
	$(GOBUILD) -o $(BINARY_DIR)/$(BINARY_NAME) -v $(MAIN_SRC)
	./$(BINARY_DIR)/$(BINARY_NAME)

# Install dependencies
dep:
	$(GOGET) -u ./...

# Cross compile for Linux
build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_LINUX) -v $(MAIN_SRC)

# Cross compile for Windows
build-windows:
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_WINDOWS) -v $(MAIN_SRC)

.PHONY: all build clean test vet run dep build-linux build-windows
