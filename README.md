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
curl --location --request GET 'http://localhost:8090/v1/echo/abc?data1=z&data1=q&data2=1&data2=3&em_id.id=123' -i \
--data-raw '{"name": "1 hello", "em_id": {"id": "12"}}'

curl -X POST -k 'http://localhost:8090/v1/echo/abc' -H 'My-Header:abc' -i -d '{"data1": ["zxc", "sdf"], "data2": [1, 2, 3], "em_id": {"id": "12"}}'

# build bin
make build

# run tests
make test

