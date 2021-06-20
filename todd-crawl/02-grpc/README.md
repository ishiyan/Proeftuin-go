# gRPC installation

Links:

- [https://grpc.io](https://grpc.io)
- [https://grpc.io/docs/quickstart/go/](https://grpc.io/docs/quickstart/go/)

These links explain how to do an installation in an old way.
If you follow them you should be in the `%GOPATH%\src` folder.

Below is a new `Go modules` installation procedure.

## Proto compiler

Go to [https://github.com/protocolbuffers/protobuf/releases](https://github.com/protocolbuffers/protobuf/releases).
Download a latest release for your platform, e.g. `proto-x.y.z-win64.zip`.
Extract an executable (e.g, `protoc.exe`) and the `include` folder to a folder on the path, e.g. `%GOPATH%\bin`.
Check if you can access it: `protoc --version`.

## Import prot-gen-go using Go modules

Have a new project with a `tools.go` file in the root of the project.

```go
// +build tools

package main

import _ "github.com/golang/protobuf/protoc-gen-go"
```

Set `GOBIN` environment variable to something.

Run `go mod init todd-grpc`. This will create a `go.mod` file.
Examine the `go.mod` file.

Run `go mod tidy`. This will download and install packages and create a `go.sym` file.
Examine the `go.sym` file.

Run `go install http://github.com/golang/protobuf/protoc-gen-go`.

## Compile proto file

Go to the `01-proto` folder and execute two commands below.

```bash
protoc -I echo echo/echo.proto --go_out=echo
protoc -I echo echo/echo.proto --go_out=plugins=grpc:echo
```

This will generate `echo/echo.pb.go` file.
