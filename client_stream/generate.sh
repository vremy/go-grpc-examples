#!/bin/bash

# Generate the client_stream_grpc.pb.go and client_stream.pb.go as described in
# the client_stream.proto. Everytime the client_stream.proto has been modified
# this script should be run to generate the new services and make them available
# in this client_stream project.
protoc proto/client_stream.proto --go_out=. --go-grpc_out=.