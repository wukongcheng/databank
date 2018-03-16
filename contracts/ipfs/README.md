# IPFS interface

## Usage

Full documentation for integrating IPFS with xci

## Development

The SOL file in contract subdirectory implements an interface to access IPFS in xci. Through this interface, you can save file to ipfs
and read files from ipfs, what is the most exciting thing is that you don't need to handle any thing about ipfs.

The solidity source code can be found at contract/ipfs.sol.

The go bindings for IPFS contracts are generated using `abigen` via the go generator:

```shell
go generate ./contracts/ipfs
```

We provide four interfaces here. To call this functions, your need to provide your account and the password for your private key store.

* Save file to IPFS. You need encode your file by base64 schemes. The fileName should be unique, or your calling will be failed.
```shell
admin.saveDataToIpfs({your account},{passphrase},{ipfsEndpoint},{filename},{your data, base64 encode})
```
* Get how many files the specific account has submitted
```shell
admin.getIpfsFileQuantity({your account},{passphrase})
```
* Get your file name by index.
```shell
admin.getIpfsFileNameByIndex({your account},{passphrase},{your file index})
```
* Get file data from IPFS.
```shell
admin.getDataFromIpfs({your account},{passphrase},{ipfsEndpoint},{filename})
```

This module status is still experimental. To use this module, you have to follow these steps:
* Run xci geth and ipfs node
* Deploy ipfs.sol on xci maunally
* Get the contract deploy address, and replace the old value TestNetAddress in /contract/ipfs/ipfs.go
* Compile a new xci geth
* Replace the running xci geth with the new one