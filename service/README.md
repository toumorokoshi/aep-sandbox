# AEP Spec Example Service

## Building

1. Install protobuf compilation tools

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

2. Compile proto

```
protoc -I=example-service/ --go_out=./example-service/generated example-service/resources.proto
```