#!/bin/bash

# Global variables
module=github.com/yael-castro/cook-book
commit=$(git log --pretty=format:'%h' -n 1)

# Command arguments
subcommand="$1"
shift

if [ "$subcommand" = "cli" ]; then
  cd cmd/cook-book-cli || exit
  path="$module/internal/container"

  printf "\nTo compile the CLI the following variables are required.\n\n"

  read -rp "DB (test): " db
  read -rp "DSN (mongodb://localhost:27017): " dsn

  printf "\nBuilding CLI in \"/build\" directory\n"

  if ! CGO=0 go build \
    -o ../../build/ \
    -tags cli \
    -ldflags="-X '$path.mongoDSN=$dsn' -X '$path.mongoDB=$db' -X '$path.gitCommit=$commit'"
  then
    exit
  fi

  cd ../../

  echo "MD5 checksum: $(md5sum build/cook-book-cli)"
  echo "Success build"

  exit
fi

if [ "$subcommand" = "http" ]; then
  cd cmd/cook-book-http || exit
  path="$module/internal/container"

  printf "\nBuilding API REST in \"/build\" directory\n"

  if ! CGO=0 go build \
    -o ../../build/ \
    -tags http \
    -ldflags="-X '$path.GitCommit=$commit'"
  then
    exit
  fi

   cd ../../

  echo "MD5 checksum: $(md5sum build/cook-book-http)"
  echo "Success build"

  exit
fi

echo "subcommand \"$subcommand\" is not valid"