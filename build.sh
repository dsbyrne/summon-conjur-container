#!/usr/bin/env sh
set -eo pipefail

BINARY_NAME=summon-conjur-docker
BUILD_PLATFORMS="darwin linux windows"
BUILD_ARCH="386"

mkdir -p bin

for PLATFORM in $BUILD_PLATFORMS; do
  for ARCH in $BUILD_ARCH; do
    echo "Building for $PLATFORM ($ARCH)..."

    GOOS=$PLATFORM \
    GOARCH=$ARCH \
      go build -x -o bin/$BINARY_NAME"_"$PLATFORM"."$ARCH
  done
done