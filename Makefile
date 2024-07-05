LOCAL_BIN:=$(CURDIR)/bin

install-deps:
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

vendor:
	go mod vendor && go mod tidy
