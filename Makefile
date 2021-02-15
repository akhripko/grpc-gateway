.PHONY: grpcgen gqlgen mockgen build run lint test test_integration dockerise deploy run_postgresql run_redis start_deps stop_deps tools

tools:
	go install \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
        google.golang.org/protobuf/cmd/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc

mod:
	#go mod tidy
	go mod download
	go mod vendor

protoc:
	#protoc -I api api/hello_world.proto --go_out=plugins=grpc:api
	protoc -I ./api \
      --go_out ./api --go_opt paths=source_relative \
      --go-grpc_out ./api --go-grpc_opt paths=source_relative \
      --grpc-gateway_out ./api --grpc-gateway_opt paths=source_relative \
      ./api/helloworld/hello_world.proto

build:
	go build -o artifacts/svc ./cmd/main.go

run:
	go run ./cmd/main.go

lint:
	cd ./src && golangci-lint run

test:
	go test -cover -v `go list ./src/...`

test_integration:
	INTEGRATION_TEST=YES go test -cover -v `go list ./...`
