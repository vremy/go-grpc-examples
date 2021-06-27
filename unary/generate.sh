#!/bin/bash

# Generate the unary_grpc.pb.go and unary.pb.go as described in the unary.proto.
# Everytime the unary.proto has been modified this script should be run to 
# generate the new services and make them available in this unary project.
protoc proto/unary.proto --go_out=. --go-grpc_out=.