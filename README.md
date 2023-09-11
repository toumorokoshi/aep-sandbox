# aep-sandbox
@toumorokoshi's sandbox for aep.dev-related projects.

Also see [DESIGN.md](./DESIGN.md)

## User Guide

The process is as follows:

1. Author your resource definition via proto messages and `google.api` annotations.
2. Generate the service proto from the resource proto via `aepc`
3. Generate your service, grpc-gateway, and openapiv2 code from the service proto via `protoc`
4. Convert the openapiv2 schema to openapiv3 with `swagger-codegen`
5. Validate the schema with `spectral`

### Setup

- Clone submodules: `git submodule update` (needed for googleapi protos)

Install swagger codegen:

```
wget https://repo1.maven.org/maven2/io/swagger/codegen/v3/swagger-codegen-cli/3.0.43/swagger-codegen-cli-3.0.43.jar -O swagger-codegen-cli.jar
```

Install the following go tools (needed for protoc):

```bash
go install \
  github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
  github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
  google.golang.org/protobuf/cmd/protoc-gen-go \
  google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

To use the spec validator, install spectral:

```
npm install -g @stoplight/spectral-cli
```

### Generated proto files and OpenAPI specification

See [scripts/regenerate-all.sh](./scripts/regenerate-all.sh).

### Starting the servers

In this design, there is a
[grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) that handles the
HTTP bindings and OpenAPI specification. So both must be started.

In two different shells:


```bash
go run service/*.go
```

```bash
go run gateway/*.go
```

### Validate API Schema via Spectral

To ensure the API generated adheres to the AEP specification,
A [spectral](https://stoplight.io/open-source/spectral) ruleset is
provided to validate your API.

If using `aepc`, this is generally not required: the generated API will
adhere to the specification. However, this is helpful for HTTP + JSON
APIs authored via other means (e.g. a web framework like Flask).

```
spectral lint openapi/openapi.json --ruleset aep-spec/spectral-ruleset.yaml
```