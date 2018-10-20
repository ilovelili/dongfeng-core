#!/bin/bash
set -e
# compile .proto files to go files in the solution
echo "fetching public packages ..."
if ! [ -x "$(command -v protoc)" ]; then
    echo "Downloading protoc ..."
    curl -OL https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip
    unzip protoc-3.6.1-linux-x86_64.zip -d protoc3      
    sudo mv ./protoc3/bin/* /usr/bin/
    sudo mv ./protoc3/include/* /usr/include
    # clean up
    sudo rm -r protoc3 protoc-3.6.1-linux-x86_64.zip
fi

for i in `find "$(cd ../services/proto; pwd)" -name "*.proto"`; do
    echo "compiling protobuf ... | source => $i"
    protoc --proto_path=$GOPATH/src:. --micro_out=$GOPATH/src --go_out=$GOPATH/src $i
done

# remove all omitempty tag
cd ../services/proto && ls *.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'

echo "done"