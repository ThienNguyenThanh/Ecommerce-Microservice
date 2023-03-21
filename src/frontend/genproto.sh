#!/usr/bin/bash

PATH=$PATH:$GOPATH/bin
protodir=../../pb

protoc --go_out=./genproto --go_opt=paths=source_relative \
    --go-grpc_out=./genproto --go-grpc_opt=paths=source_relative \
     $protodir/general.proto