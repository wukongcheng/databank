#!/bin/sh

tar -xvf go-ipfs_v0.4.14_linux-386.tar.gz

cd go-ipfs

sudo ./install.sh

cd ..

rm -rf go-ipfs
