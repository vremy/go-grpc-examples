# GO GRPC Examples
I have setup this repository in response to the lack of good and working GRPC
examples in Go. This repository contains examples for all four GRPC modes like
unary, server streaming, client streaming and bidirectional streaming. This
repository contains four directories where each GRPC mode is demonstrated with
a client and server example. I have tried to document the examples in a manner
that should be straight forward.

---

## Directory explanation
Each directory contains the following files which make up the project.

***.proto**   
This file contains the description of all GRPC services that should become
available in the project.

**generate.sh**   
This script is used to generate `*_grpc.pb.go` and `*.pb.go` files that will
make the services described in the `*.proto` available in the project.

**server.go**   
This file contains all the logic to handle RPC calls from clients.

**client.go**   
This file contains all the logic for requesting RPC services available on the
server.

---

## Go package management
A go package requires a go.mod file to track all dependencies required for the
package. To initialize this file i used the command below.

```sh
go mod init github.com/vremy/go-grpc-examples
```

With the following command dependencies can be added or removed depending if
they are used in the codebase.

```sh
go mod tidy
```