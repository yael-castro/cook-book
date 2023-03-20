#!/bin/bash

# Global variables
declare -r module=github.com/yael-castro/cb-search-engine-api
declare -r commit=$(git log --pretty=format:'%h' -n 1)

# Command arguments
subcommand="$1"
shift

if [ "$subcommand" = "cli" ]; then
  cd cmd/cli
  path="$module/cmd/cli/container"

  printf "DSN:"
  if ! read -r dsn ; then
    exit
  fi

  printf "DB:"
  if ! read -r db ; then
    exit
  fi

  echo "Building CLI in \"/build\" directory"

  if ! CGO=0 go build \
    -o ../../build/ \
    -ldflags="-X '$path.mongoDSN=$dsn' -X '$path.mongoDB=$db' -X '$path.GitCommit=$commit'"
  then
    exit
  fi

  echo "Success build"
  exit
fi

if [ "$subcommand" = "server" ]; then
  cd cmd/server
  path="$module/cmd/server/container"

  echo "Building API REST in \"/build\" directory"

  if ! CGO=0 go build \
    -o ../../build/ \
    -ldflags="-X '$path.GitCommit=$commit'"
  then
    exit
  fi

  echo "Success build"
  exit
fi

echo "subcommand \"$subcommand\" is not valid"