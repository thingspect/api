# Thingspect API

This repository contains the API
[Protobuf](https://developers.google.com/protocol-buffers/) and
[gRPC](https://grpc.io/) definitions for the
[Thingspect](http://www.thingspect.com/) Atlas platform. API design is based
on the Google [API Design Guide](https://cloud.google.com/apis/design).

## Getting Started

### Go

Documentation: https://pkg.go.dev/github.com/thingspect/api/go

```
go get -u github.com/thingspect/api/go
```

Example gRPC code: `go/example/`

### Python

The Python package can be copied locally from this repository or regenerated
using the Protobuf definitions. See the
[gRPC Python quick start](https://grpc.io/docs/languages/python/quickstart/) for
more information.

## Building

These instructions require
[Docker](https://docs.docker.com/get-started/overview/) to be installed.

```
make
```
