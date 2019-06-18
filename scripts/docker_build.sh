#!/bin/bash
set -e

cd ..

# docker build
docker build -t ilovelili/dongfeng-core .

echo "Bye"