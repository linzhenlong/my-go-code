#!/usr/bin/env bash

CPAHT=$PWD

#进入protobuf目录
cd $CPAHT/protobuf

#生成文件
protoc --go_out=plugins=grpc:../server/services *.proto
protoc --go_out=plugins=grpc:../client/services *.proto

protoc --grpc-gateway_out=logtostderr=true:../server/services *.proto
protoc --grpc-gateway_out=logtostderr=true:../client/services *.proto

echo "success"





