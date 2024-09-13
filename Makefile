# Используем bin в текущей директории для установки плагинов protoc
LOCAL_BIN:=$(CURDIR)/bin
LOCAL_MIGRATION_DIR="./internal/app/migrations/"
LOCAL_MIGRATION_DSN="host=localhost port=5433 dbname=tenders user=postgres password=password sslmode=disable"


.PHONY: dc run test lint

dc:
	docker-compose up  --remove-orphans --build

build:
	go build -race -o app cmd/main.go

run:
	go build -race -o app cmd/main.go && \
	SERVER_ADDRESS=:8085 \
	DEBUG_ERRORS=1 \
	POSTGRES_CONN="postgres://postgres:password@127.0.0.1:5432/tenders?sslmode=disable" \
	MIGRATIONS_PATH="file://./internal/app/migrations" \
	./app

test:
	go test -race ./...

install-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.57.2

lint:
	golangci-lint run ./...

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

generate:
	go generate ./...

migrate:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up

migrate-down:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v