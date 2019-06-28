GOPATH := $(shell go env GOPATH)
build: gen-rpc
	GO111MODULE=on go build ./...; \
	GO111MODULE=on go install ./...

init:
	GO111MODULE=on go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v1.8.5
	GO111MODULE=on go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	GO111MODULE=on go get -u github.com/golang/protobuf/protoc-gen-go

gen-rpc:
	protoc \
	-I api/v1/ \
	-I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.8.5/third_party/googleapis \
	api/v1/v1.proto --go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. --descriptor_set_out=./api/v1/api_descriptor.pb