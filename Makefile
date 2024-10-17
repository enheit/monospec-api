# Note: This file was copied from ChatGPT @ Roman

# Variables
GOARCH=arm64
GOOS=linux

# Migration database URL
MIGRATE_BIN=/usr/bin/migrate
MIGRATE_DB=postgres://postgres:vkwWOJh9610DxDTWisD8K,6-e5BBLN@monospecapistack-rdsnestedstackrdsnest-rds34d05673-lwt37qlx9xe6.cfiwiiwq0xla.eu-central-1.rds.amazonaws.com:5432/monospec

# Targets
.PHONY: all build clean migrate rollback deploy-roman

all: build

# Build all Lambda functions
build:
	@echo "Building all Lambda functions..."
	@find . -type d \( -name 'cdk.out' \) -prune -o -type f -name '*-lambda.go' -print | while read file; do \
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

deploy-roman: build
	@echo "Deploying to roman..."
	npx cdk deploy MonospecApiStack --profile ms-roman

