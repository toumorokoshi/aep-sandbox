# aep-sandbox
@toumorokoshi's sandbox for aep.dev-related projects.


## User Guide

The process is as follows:

1. Author your resource definition via proto messages and `google.api` annotations.
2. Generate the service proto from the resource proto via `aepc`
3. Generate your service grpc code from the service proto via `protoc`

### Generate service code

```bash
protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative example-service/resources.proto
```