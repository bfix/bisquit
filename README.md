
[![Build Status](https://travis-ci.org/bfix/bisquit.svg?branch=master)](https://travis-ci.org/bfix/bisquit)
[![Go Report Card](https://goreportcard.com/badge/github.com/bfix/bisquit)](https://goreportcard.com/report/github.com/bfix/bisquit)
[![GoDoc](https://godoc.org/github.com/bfix/bisquit?status.svg)](https://godoc.org/github.com/bfix/bisquit)

Bisquit: Bisq API client
========================

(c) 2021-2023 Bernd Fix <brf@hoi-polloi.org>   >Y<

bisquit is free software: you can redistribute it and/or modify it
under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License,
or (at your option) any later version.

bisquit is distributed in the hope that it will be useful, but
WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.

SPDX-License-Identifier: AGPL3.0-or-later

## Prerequisites

This library is intended to be used with Go1.20+ and might not work
correctly on previous versions.

### Protobuf compiler (optional)

You need to have a newer protobuf compiler for Go installed on your
system; make sure you installed necessary dependencies:

```bash
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go
```

### Protobuf definitions (update optional)

You also need the Protobuf definitions from the Bisq source tree (found at
`proto/src/main/proto/grpc.proto` and `proto/src/main/proto/pb.proto` to
generate Go stubs for the gRPC process. These files are copied to this
repository and are kept up to date if possible; if I miss to update them
you can just copy them over from the newest Bisq source tree and modify
both of them by adding

```
option go_package = ".;bisquit";
```

in the `option` section at the beginning of the `*.proto` files.

## Building

### Generating Protobuf stubs (optional)

The Go source code generated from the Protobuf definitions are included in
the repository, so this step is optional and only required if newer Protobuf
definitions are available.

To generate the Go code from the Protobuf definitions, run the following
commands in the base directory:

```bash
protoc --go_out=. --go_opt=paths=source_relative pb.proto
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc.proto
```

This should generate three (new) Go files: `grpc_grpc.pb.go`,
`grpc.pb.go` and `pb.pb.go`.

### Compiling the library

Make sure all dependencies are installed by running

```bash
go mod tidy
```

and then build the library with

```bash
go build
```

## Testing

### Bisq daemon

You need a running Bisq daemon (v1.9.10+) with enabled gRPC and API password
to run the unit tests. Use

```bash
export BISQ_API_PASSWORD="my_secret"
export BISQ_API_HOST="localhost:9998"
```

to pass the API password and host settings to the tests.
