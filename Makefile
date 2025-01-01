include local.env

LOCAL_BIN:=$(CURDIR)/bin

LOCAL_MIGRATION_DIR:=$(MIGRATION_DIR)
LOCAL_MIGRATION_DSN:="host=localhost user=$(PG_USER) password=$(PG_PASSWORD) dbname=$(PG_DATABASE_NAME) port=$(PG_PORT_OUTER)"

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-server-api

generate-server-api:
	mkdir -p generated
	protoc -I proto proto/chat_server_v1/chat_server.proto \
	--go_out=./generated \
	--go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=./generated \
	--go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \

bin-build:
	GOOS=linux GOARCH=amd64 go build -o chat_service_linux cmd/main.go

docker-build:
	docker buildx build --no-cache --platform linux/amd64 -t chat_service:v0.0.1

local-MIGRATION-status:
	bin/goose -dir $(LOCAL_MIGRATION_DIR) postgres $(LOCAL_MIGRATION_DSN) status -v

local-MIGRATION-up:
	bin/goose -dir $(LOCAL_MIGRATION_DIR) postgres $(LOCAL_MIGRATION_DSN) up -v

local-MIGRATION-down:
	bin/goose -dir $(LOCAL_MIGRATAION_DIR) postgres $(LOCAL_MIGRATION_DSN) down -v

compose-up-local:
	docker compose --profile local --env-file local.env up -d --build  --remove-orphans

compose-down-local:
	docker compose --profile local --env-file local.env down

compose-up-db:
	docker compose --env-file test.env up -d --build pg_chat_server_test --remove-orphans

compose-down-db:
	docker compose --env-file test.env down pg_chat_server_test