#!/bin/bash

# Global variables
module=github.com/yael-castro/cb-search-engine-api
commit=$(git log --pretty=format:'%h' -n 1)

# Command arguments
subcommand="$1"
shift

if [ "$subcommand" = "cli" ]; then
  cd cmd/cli || exit
  path="$module/cmd/cli/container"

  printf "\nTo compile the CLI the following variables are required.\n\n"

  read -rp "DB (test): " db
  read -rp "DSN (mongodb://localhost:27017): " dsn

  printf "\nBuilding CLI in \"/build\" directory\n"

  if ! CGO=0 go build \
    -o ../../build/ \
    -ldflags="-X '$path.mongoDSN=$dsn' -X '$path.mongoDB=$db' -X '$path.gitCommit=$commit'"
  then
    exit
  fi

  cd ../../

  echo "MD5 checksum: $(md5sum build/cli)"
  echo "Success build"

  exit
fi

if [ "$subcommand" = "server" ]; then
  cd cmd/server || exit
  path="$module/cmd/server/container"

  printf "\nBuilding API REST in \"/build\" directory\n"

  if ! CGO=0 go build \
    -o ../../build/ \
    -ldflags="-X '$path.GitCommit=$commit'"
  then
    exit
  fi

   cd ../../

  echo "MD5 checksum: $(md5sum build/server)"
  echo "Success build"

  exit
fi

echo "subcommand \"$subcommand\" is not valid"