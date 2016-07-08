#!/usr/bin/env bash
set -eo pipefail

GOLANG_VER=1.7-alpine

docker run --rm -v $PWD:/go/src/app -w /go/src/app golang:$GOLANG_VER ./build.sh