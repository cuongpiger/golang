#!/bin/bash

docker_build() {
  docker buildx build --push \
    --platform linux/amd64,linux/arm64 \
    --tag manhcuong8499/multi_arch_sample:buildx-latest .
}

case $1 in
  "build")
    docker_build
    ;;
  *)
    echo "Usage: $0 build"
    ;;
esac
