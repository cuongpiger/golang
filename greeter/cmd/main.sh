#!/bin/bash

go_cmd=$(which go)

if [ -z "$go_cmd" ]; then
    echo "go not found"
    exit 1
fi

case $1 in
    build)
        $go_cmd build -o bin/greeter cmd/main.go
        ;;
    run)
        $go_cmd run cmd/main.go
        ;;
    *)
        echo "Usage: $0 {build|run}"
        exit 1
esac