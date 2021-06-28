#!/bin/bash

# Generate the server_stream_grpc.pb.go and server_stream.pb.go as described in
# the server_stream.proto. Everytime the server_stream.proto has been modified
# this script should be run to generate the new services and make them available
# in this server_stream project.
protoc proto/server_stream.proto --go_out=. --go-grpc_out=.