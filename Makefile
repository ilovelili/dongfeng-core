default: build

all: build

local:
	cd ./scripts && ./local_setup.sh

build:
	cd ./scripts && ./docker_build.sh
	
proto:
	cd ./scripts && ./proto_compiler.sh

test:
	cd ./scripts && ./test.sh

.PTHONY: all local build proto test