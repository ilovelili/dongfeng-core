#!/bin/bash
set -e

# move to root directory
cd ../

echo "fetching public packages ..."
if ! [ -x "$(command -v dep)" ]; then
    echo "Downloading Go dependency management tool ..."
    go get -u github.com/golang/dep/cmd/dep
fi

# init when necessary
if [ ! -f Gopkg.toml ]; then
    dep init
fi

rm -f Gopkg.lock
rm -rf ./vendor

dep ensure

# compile protobuf
make proto
echo "Bye"