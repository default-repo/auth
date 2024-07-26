help:
	@echo [1] Befor start of using '"make"' command, you need to set up '"ENV"' variable with one of the values:"\n"\
		- PROD"\n"\
		- LOCAL"\n"
	@echo [2] A list of available commands"\n"\
		- migration-gen tag= : create a new migration file"\n" \
		- migration-up : apply all available migrations"\n"\
		- migration-down: roll back a single migration from the current version"\n"\
		- migration-status: print the status of all migrations"\n"\
		- install-db-deps: set dependencies 'for' working with the database"\n"\
		- install-proto-deps: install dependencies 'for' proto* "\n"\
 		- get-deps: get all dependencies"\n"\
 		- gen: generate API \"*.pb.go files\""\n"\
 		- install-lint: install linter"\n"\
        - lint: inspect code"\n"\
        - vendor: tidy code and get dependencies"\n"\
        - format: format imports"\n"\
        - run-local: run project with LOCAL variables"\n"\
        - run-prod: run project with PROD variables

#================================================================================================================

# PROD || LOCAL
ENV=PROD

ifeq ($(ENV), PROD)
	include .env.prod
else
	include .env.local
endif

LOCAL_BIN:=$(CURDIR)/bin
GOOSE := $(LOCAL_BIN)/goose

# ---------------------- DB ------------------------------

migration-gen:
	@[ "${tag}" ] || (echo "Error: tag is not set"; exit 1)
	$(GOOSE) -dir $(MIGRATION_DIR) create ${tag} sql

migration-up:
	$(GOOSE) -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} up -v

migration-down:
	$(GOOSE) -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} down -v

migration-status:
	$(GOOSE) -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} status -v

install-db-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

# ----------------------- INFRA ---------------------------

install-proto-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

gen:
	make generate-auth-api

generate-auth-api:
	rm -rf pkg/proto/auth_v1  && mkdir -p pkg/proto/auth_v1
	protoc --proto_path api/proto/auth_v1 \
	--go_out=pkg/proto/auth_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/proto/auth_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--experimental_allow_proto3_optional \
	api/proto/auth_v1/auth.proto

install-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... --config .golangci.pipeline.yaml

vendor:
	go mod tidy && go mod vendor

format:
	goimports -w ./internal/ && goimports -w ./cmd/

run-local:
	@echo "Running in development mode!" && go run cmd/auth/main.go -config-path=.env.local

run-prod:
	@echo "Running in production mode!" && go run cmd/auth/main.go -config-path=.env.prod
