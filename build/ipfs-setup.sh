#!/bin/sh

if [ x`uname`x = xDarwinx ]; then
  echo 'Wanning: macOS is not supported.' >&2
  exit 1
fi

if [ x`uname`x != xLinuxx ]; then
  echo 'Wanning: only linux is supported.' >&2
  exit 1
fi

if [ -x "$(command -v ipfs)" ]; then
  echo 'Wanning: ipfs is installed.' >&2
  exit 1
fi

tar -xvf go-ipfs_v0.4.14_linux-386.tar.gz
cd go-ipfs
sudo ./install.sh

cd ..
rm -rf go-ipfs
