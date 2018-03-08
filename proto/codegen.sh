#!/usr/bin/env bash
mkdir $GOPATH/src/proto/crm
protoc --go_out=plugins=grpc:$GOPATH/src/proto/crm ./crm.api.proto