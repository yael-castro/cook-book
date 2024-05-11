#!/bin/bash

# Variables for only read
container="github.com/yael-castro/cook-book/internal/container"
commit=$(git log --pretty=format:'%h' -n 1)

# Command arguments
subcommand="$1"
shift

function build() {
    cd "./cmd/$binary"

    if ! CGO=0 go build \
      -o ../../build/ \
      -tags "$tags" \
      -ldflags="-X '$container.GitCommit=$commit'"
    then
      exit
    fi

    cd ../../

    echo "MD5 checksum: $(md5sum "build/$binary")"
    echo "Success build"
    exit 0
}


if [ "$subcommand" = "cli" ]; then
  binary="cook-book-cli"
  tags="cli"

  printf "\nBuilding CLI in \"build\" directory\n"
  build
fi

if [ "$subcommand" = "http" ]; then
  binary=cook-book-http
  tags="http"

  printf "\nBuilding API REST in \"build\" directory\n"
  build
fi

exit 1
echo "Invalid subcommand: $subcommand"