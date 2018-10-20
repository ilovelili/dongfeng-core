#!/bin/bash
set -e

# move to root directory
cd ../

echo "fetching public packages ..."
if ! [ -x "$(command -v dep)" ]; then
    echo "Downloading Go dependency management tool ..."
    go get -u github.com/golang/dep/cmd/dep
fi

rm -f Gopkg.lock Gopkg.toml
rm -rf ./vendor

dep init

# compile protobuf
make proto

echo "Bye"