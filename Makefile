depend:
	go install github.com/bufbuild/buf/cmd/buf@v1.4.0
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.10.0
	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	google.golang.org/protobuf/cmd/protoc-gen-go
	google.golang.org/grpc/cmd/protoc-gen-go-grpc

build:
	buf generate --path api/**/**/

all: depend build