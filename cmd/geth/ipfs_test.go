package main

import (
	"testing"
	"os"
	"fmt"
	"bytes"
	"io"
)

func TestIPFSAdd(t *testing.T) {
	geth := runGeth(t, "ipfs", "add", "config.go")

	defer geth.ExpectExit()
	geth.Expect("QmfTDzioWPdsJLPPw6nKyo5SP3obVqWBcgcTZy1d91U4tG")
}

func TestIPFSCat(t *testing.T) {
	//geth := runGeth(t, "ipfs", "cat", "QmUfZ9rAdhV5ioBzXKdUTh2ZNsz9bzbkaLVyQ8uc8pj21F")
	geth := runGeth(t, "ipfs", "cat", "QmfTDzioWPdsJLPPw6nKyo5SP3obVqWBcgcTZy1d91U4tG")

	buf := bytes.NewBuffer(nil)
	file, _ := os.Open("config.go") // Error handling elided for brevity.
	io.Copy(buf, file)           // Error handling elided for brevity.
	file.Close()
	str := string(buf.Bytes())

	geth.Expect(str)

	defer geth.ExpectExit()
}

func TestIPFSGet(t *testing.T) {
	//geth := runGeth(t, "ipfs", "cat", "QmUfZ9rAdhV5ioBzXKdUTh2ZNsz9bzbkaLVyQ8uc8pj21F")
	targetFile := "../../../ipfs_config.go"
	geth := runGeth(t, "ipfs", "get", "QmfTDzioWPdsJLPPw6nKyo5SP3obVqWBcgcTZy1d91U4tG", targetFile)
	defer geth.ExpectExit()
	defer cleanTestFile(targetFile)
}

func cleanTestFile(path string)  {
	err := os.Remove(path)
	if err != nil {
		fmt.Errorf(err.Error())
	}
}

