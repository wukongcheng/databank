package main

import (
	"github.com/cheekybits/is"
	"bytes"
	"testing"
	"github.com/ipfs/go-ipfs-api"
)

const (
//	shellUrl     = "localhost:5001"
)

func TestAdd(t *testing.T) {
	is := is.New(t)
	s := shell.NewShell(shellUrl)

	mhash, err := s.Add(bytes.NewBufferString("Hello IPFS Shell tests"))
	is.Nil(err)
	is.Equal(mhash, "QmUfZ9rAdhV5ioBzXKdUTh2ZNsz9bzbkaLVyQ8uc8pj21F")
}

func TestIPFSAdd(t *testing.T) {
	geth := runGeth(t, "ipfs", "add", "Hello IPFS Shell tests")
	defer geth.ExpectExit()
	geth.Expect("QmUfZ9rAdhV5ioBzXKdUTh2ZNsz9bzbkaLVyQ8uc8pj21F")
}

func TestIPFSCat(t *testing.T) {
	geth := runGeth(t, "ipfs", "cat", "QmUfZ9rAdhV5ioBzXKdUTh2ZNsz9bzbkaLVyQ8uc8pj21F")
	defer geth.ExpectExit()
	geth.Expect("Hello IPFS Shell tests")
}

