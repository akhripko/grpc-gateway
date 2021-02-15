# grpc-gateway
example of https://grpc-ecosystem.github.io/grpc-gateway/

# get tools
make tools

# generate proto api
make protoc

# run
make run

# send test http request
curl -X POST -k http://localhost:8090/v1/example/echo -H 'My-Header:abc' -i -d '{"name": " hello"}'

# build bin
make build

# run tests
make test

