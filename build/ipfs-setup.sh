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
else
  echo 'Installing ipfs ... ' >&1

  tar -xvf go-ipfs_v0.4.14_linux-386.tar.gz
  cd go-ipfs
  sudo ./install.sh

  cd ..
  rm -rf go-ipfs

  echo 'Done!' >&1

  ipfs init
  echo 'Init ipfs ... ' >&1

  cp ./swarm.key ~/.ipfs
  echo 'Set ipfs swarm.key ... ' >&1

  ipfs bootstrap rm --all
  ipfs bootstrap add /ip4/104.131.131.82/tcp/4001/ipfs/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ
  echo 'Reset ipfs bootstrap peers ... ' >&1

  nohup ipfs daemon &
  echo 'Start ipfs daemon ... ' >&1

fi

