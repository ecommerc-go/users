include .env
LOCAL_BIN:=$(CURDIR)/bin

LOCAL_MIGRATION_DIR=$(MIGRATION_DIR)
LOCAL_MIGRATION_DSN="host=localhost port=$(PG_PORT) dbname=$(PG_DATABASE_NAME) user=$(PG_USER) password=$(PG_PASSWORD) sslmode=disable"

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0
	GOBIN=$(LOCAL_BIN) go install golang.org/x/tools/cmd/goimports@v0.18.0
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml



migration-status:
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

migration-up:
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

migration-down:
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v