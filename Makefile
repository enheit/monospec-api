# Note: This file was copied from ChatGPT @ Roman

# Variables
GOARCH=arm64
GOOS=linux
LAMBDA_DIR=./lambdas

# Migration database URL
MIGRATE_BIN=/usr/bin/migrate
MIGRATE_DB=postgres://postgres:9HxW.CGwtuo%5E=,mOYSKD%5EwG2a==oNx@monospecapistack-rdsnestedstackrdsnest-rds34d05673-b5mbbyvdtfuv.cfiwiiwq0xla.eu-central-1.rds.amazonaws.com:5432/monospec

# Lambda functions
LAMBDA_FUNCTIONS = \
	get-me \
	get-user-appointments

# Targets
.PHONY: all build clean migrate

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

# Run migrations
migrate:
	@echo "Running migrations..."
	$(MIGRATE_BIN) -path ./postgres/migrations -database $(MIGRATE_DB) up

rollback:
	@echo "Rolling back migrations with argument: $(arg)"
	@$(MIGRATE_BIN) -path ./postgres/migrations -database $(MIGRATE_DB) down $(rc)

# rc = Rollback Count
rc := 1

