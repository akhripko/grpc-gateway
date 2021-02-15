# grpc-gateway
example of https://grpc-ecosystem.github.io/grpc-gateway/

# get tools
make tools

# generate proto api
make protoc

# run
make run
&
curl -X POST -k http://localhost:8090/v1/example/echo -d '{"name": " hello"}'

# build bin
make build

# run tests
make test

