#!/usr/bin/env bash


for file in ./web/protobuf/*/*.proto; do
     protoc -I=. -I=../  --go_out=plugins=grpc:../ $file
done