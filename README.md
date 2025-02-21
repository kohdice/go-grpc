# go-grpc

A sample implementation of gRPC in Go

## Usage

### Install dependencies

```bash
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### Code generation

- Run the following command in the api directory.

```bash
protoc --go_out=../pkg/grpc --go_opt=paths=source_relative \
        --go-grpc_out=../pkg/grpc --go-grpc_opt=paths=source_relative \
        ./sample.proto
```
