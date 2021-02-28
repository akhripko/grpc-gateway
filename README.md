# grpc-gateway
example of https://grpc-ecosystem.github.io/grpc-gateway/

https://github.com/googleapis/googleapis/blob/master/google/api/http.proto

# get tools
make tools

# generate proto api
make protoc

# run
make run

# send test http request
curl -X GET -k http://localhost:8090/v1/echo/abc\?data1=zxc\&data1=sdf\&data2=1\&data2=2\&data2=3 -i
curl -X POST -k http://localhost:8090/v1/echo/abc -H 'My-Header:abc' -i -d '{"data1": ["zxc", "sdf"], "data2": [1, 2, 3]}'

# build bin
make build

# run tests
make test

