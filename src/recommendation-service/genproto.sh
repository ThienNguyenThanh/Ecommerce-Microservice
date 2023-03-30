#!/bin/bash

# requires gRPC tools:
#   pip install -r requirements.txt

python -m grpc_tools.protoc -I../../pb --python_out=./genproto  --grpc_python_out=./genproto ../../pb/general.proto
