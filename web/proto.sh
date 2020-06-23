#!/usr/bin/env bash

cd ../../
for file in ./go-web/web/protobuf/*/*.proto; do
   protoc -I=.  --go_out=plugins=grpc:. $file
done