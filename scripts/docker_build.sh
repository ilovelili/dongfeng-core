#!/bin/bash
set -e
# move to root directory
cd ..
# docker build
docker build -t dongfeng-core . -f Dockerfile
echo "Bye"