#!/usr/bin/env bash

current_path=$PWD

cd $current_path/protofiles


protoc --go_out=plugins=grpc:../server/services --validate_out=lang=go:../server/services *.proto
protoc --grpc-gateway_out=logtostderr=true:../server/services --validate_out=lang=go:../server/services *.proto

protoc --go_out=plugins=grpc:../client/services --validate_out=lang=go:../client/services *.proto
protoc --grpc-gateway_out=logtostderr=true:../client/services --validate_out=lang=go:../client/services *.proto