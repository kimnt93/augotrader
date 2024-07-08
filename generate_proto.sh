#!/bin/bash

# Create the proto directory if it doesn't exist
mkdir -p proto

# Generate the protobuf files for Python
python -m grpc_tools.protoc -I. --python_out=proto --grpc_python_out=proto service.proto

# Generate the protobuf files for Go
protoc -I. --go_out=proto --go-grpc_out=proto service.proto
