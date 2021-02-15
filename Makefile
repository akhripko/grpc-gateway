.PHONY: grpcgen gqlgen mockgen build run lint test test_integration dockerise deploy run_postgresql run_redis start_deps stop_deps

mod:
	go mod download
	go mod vendor

protoc:
	protoc -I api api/hello_world.proto --go_out=plugins=grpc:api

build:
	go build -o artifacts/svc ./cmd/svc/main.go

run:
	go run ./cmd/main.go

lint:
	cd ./src && golangci-lint run

test:
	go test -cover -v `go list ./src/...`

test_integration:
	INTEGRATION_TEST=YES go test -cover -v `go list ./...`
