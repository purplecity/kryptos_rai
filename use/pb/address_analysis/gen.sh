#! /bin/sh
protoc --go_out=plugins=grpc:. *.proto
#protoc --go-grpc_out=. *.proto
