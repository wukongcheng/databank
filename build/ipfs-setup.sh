#!/bin/sh

if [ -x "$(command -v ipfs)" ]; then
  echo 'Wanning: ipfs is installed.' >&2
  exit 1
fi

tar -xvf go-ipfs_v0.4.14_linux-386.tar.gz
cd go-ipfs
sudo ./install.sh

cd ..
rm -rf go-ipfs