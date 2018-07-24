.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/wukongcheng/databank/cmd/swarm
	go build -o bin/geth     github.com/wukongcheng/databank/cmd/geth
	go build -o bin/bootnode github.com/wukongcheng/databank/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/wukongcheng/databank/swarm/...
