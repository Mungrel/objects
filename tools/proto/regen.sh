#!/bin/bash

# Check if protoc is installed
if ! command -v protoc &> /dev/null
then
    echo "protoc could not be found. Please install Protocol Buffers compiler (protoc) first."
    echo "$ brew install protoc"
    exit 1
fi

# Check if protoc-gen-go plugin is installed
if ! command -v protoc-gen-go &> /dev/null
then
    echo "Error: protoc-gen-go could not be found. Please install the Go protocol buffers plugin (protoc-gen-go) first:"
    echo "$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest"
    echo "$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest"
    echo '$ export PATH="$PATH:$(go env GOPATH)/bin"'
    exit 1
fi

# Run protoc with your .proto files
protoc \
    --proto_path=internal/proto \
    --go_out=paths=source_relative:internal/proto \
    --go-grpc_out=internal/proto \
    --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
    internal/proto/**/*.proto

echo 'Done!'
