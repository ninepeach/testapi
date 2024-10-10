# Define the Go binary name
BINARY_NAME=testapi

# Build output directories
BUILD_DIR=build

# Define GOOS and GOARCH for different platforms
LINUX_GOOS=linux
OSX_GOOS=darwin
LINUX_GOARCH=amd64
OSX_GOARCH=arm64

# Paths for output binaries
LINUX_OUTPUT=$(BUILD_DIR)/$(BINARY_NAME)_linux_$(LINUX_GOARCH)
OSX_OUTPUT=$(BUILD_DIR)/$(BINARY_NAME)_osx_$(OSX_GOARCH)

# Default target: build both Linux and macOS binaries
all: linux osx

# Clean previous builds
clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)

# Create the build directory if it doesn't exist
$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

# Build for Linux
linux: $(BUILD_DIR)
	@echo "Building Linux binary..."
	GOOS=$(LINUX_GOOS) GOARCH=$(LINUX_GOARCH) go build -o $(LINUX_OUTPUT)
	@echo "Linux binary created: $(LINUX_OUTPUT)"

# Build for macOS (arm64)
osx: $(BUILD_DIR)
	@echo "Building macOS binary (arm64)..."
	GOOS=$(OSX_GOOS) GOARCH=$(OSX_GOARCH) go build -o $(OSX_OUTPUT)
	@echo "macOS binary created: $(OSX_OUTPUT)"

# Clean and build all
rebuild: clean all

.PHONY: all clean linux osx rebuild
