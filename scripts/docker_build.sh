#!/bin/bash
set -e
# move to root directory
cd ../services/server
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ../../server .
cp config.*.json ../../

cd ../../
# docker build
docker build -t ilovelili/dongfeng-core . -f DockerFile.lite

# clean up
rm server config.*.json
echo "Bye"