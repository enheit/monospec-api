# Note: This file was copied from ChatGPT @ Roman
# Variables
GOARCH=arm64
GOOS=linux
LAMBDA_DIR=./lambdas

# Lambda functions
LAMBDA_FUNCTIONS = \
	get-me \
	get-user-appointments

# Targets
.PHONY: all build clean

all: build

# Build all Lambda functions
build: $(LAMBDA_FUNCTIONS)

# Build each Lambda function
$(LAMBDA_FUNCTIONS):
	@echo "Building $@..."
	GOARCH=$(GOARCH) GOOS=$(GOOS) go build -o $(LAMBDA_DIR)/$@/bootstrap $(LAMBDA_DIR)/$@/$@.go && \
	chmod +x $(LAMBDA_DIR)/$@/bootstrap

# Clean up built binaries
clean:
	@echo "Cleaning up..."
	for func in $(LAMBDA_FUNCTIONS); do \
		rm -f $(LAMBDA_DIR)/$$func/bootstrap; \
	done

