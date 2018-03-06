#!/usr/bin/env bash
mkdir $GOPATH/src/proto
protoc --go_out=plugins=micro:$GOPATH/src/proto ./crm.api.proto