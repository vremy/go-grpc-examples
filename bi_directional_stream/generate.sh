#!/bin/bash

# Generate the bi_directional_stream_grpc.pb.go and bi_directional_stream.pb.go
# as described in the bi_directional_stream.proto. Everytime the
# bi_directional_stream.proto has been modified this script should be run to
# generate the new services and make them available in this
# bi_directional_stream project.
protoc proto/bi_directional_stream.proto --go_out=. --go-grpc_out=.