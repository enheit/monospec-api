# Note: This file was copied from ChatGPT @ Roman

# Variables
GOARCH=arm64
GOOS=linux

# Migration database URL
MIGRATE_BIN=/usr/bin/migrate
MIGRATE_DB=postgres://postgres:9HxW.CGwtuo%5E=,mOYSKD%5EwG2a==oNx@monospecapistack-rdsnestedstackrdsnest-rds34d05673-b5mbbyvdtfuv.cfiwiiwq0xla.eu-central-1.rds.amazonaws.com:5432/monospec

# Targets
.PHONY: all build clean migrate

all: build

# Build all Lambda functions
build:
	@echo "Building all Lambda functions..."
	@find . -type f -name '*-lambda.go' | while read file; do \
		func_name=$$(basename $${file%-lambda.go}); \
		func_dir=$$(dirname "$$file"); \
		echo "Building $$func_name..."; \
		GOARCH=$(GOARCH) GOOS=$(GOOS) go build -o "$$func_dir/bootstrap" "$$file" && \
		chmod +x "$$func_dir/bootstrap"; \
	done

# Clean up built binaries
clean:
	@echo "Cleaning up..."
	@find . -name 'bootstrap' -exec rm -f {} +

# Run migrations
migrate:
	@echo "Running migrations..."
	$(MIGRATE_BIN) -path ./postgres/migrations -database $(MIGRATE_DB) up

rollback:
	@echo "Rolling back migrations with argument: $(arg)"
	@$(MIGRATE_BIN) -path ./postgres/migrations -database $(MIGRATE_DB) down $(rc)

# rc = Rollback Count
rc := 1
