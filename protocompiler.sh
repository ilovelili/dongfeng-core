#!/bin/sh
set -e
# compile .proto files to go files in the solution
PKG_OK=$(command -v protoc)
if [ "" = "$PKG_OK" ]; then
    echo "No protobuf. Setting up..."
    # Make sure you grab the latest version
    curl -OL https://github.com/google/protobuf/releases/download/v3.2.0/protoc-3.2.0-linux-x86_64.zip
    unzip protoc-3.2.0-linux-x86_64.zip -d protoc3
      
    sudo mv ./protoc3/bin/* /usr/bin/
    sudo mv ./protoc3/include/* /usr/include

    # clean up
    sudo rm -r protoc3 protoc-3.2.0-linux-x86_64.zip
fi

for i in `find "$(cd services; pwd)" -name "*.proto"`; do
    echo "compiling protobuf ... | source => $i"
    protoc --proto_path=$GOPATH/src:. --micro_out=$GOPATH/src --go_out=$GOPATH/src $i
done

echo "done"