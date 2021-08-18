#!/bin/sh
SRC_DIR=.
DST_DIR=.
#生成 *.pb.go 文件需要指定一下依赖软件的版本
#go get -u -v  github.com/golang/protobuf@v1.2.0
#go get -u -v  google.golang.org/grpc@v1.20.0

#./codegen.sh
#protoc -I=$SRC_DIR --go_out=$DST_DIR/ $SRC_DIR/*.proto
protoc -I=$SRC_DIR --go_out=plugins=grpc:$DST_DIR/ $SRC_DIR/schema.tl.sync_service.proto
protoc -I=$SRC_DIR --go_out=plugins=grpc:$DST_DIR/ $SRC_DIR/schema.tl.crc32.proto
protoc -I=$SRC_DIR --go_out=plugins=grpc:$DST_DIR/ $SRC_DIR/schema.tl.sync.proto
protoc -I=$SRC_DIR --go_out=plugins=grpc:$DST_DIR/ $SRC_DIR/rpc_error_codes.proto

gofmt -w *.go
