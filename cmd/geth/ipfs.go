package main

import (
	"github.com/xcareteam/xci/cmd/utils"
	"gopkg.in/urfave/cli.v1"
	"github.com/ipfs/go-ipfs-api"
	"fmt"
	"bytes"
	"strings"
)

var (
	shellUrl     = "localhost:5001"

	ipfsCommand = cli.Command{
		Name:     "ipfs",
		Usage:    "Interactive with ipfs",
		Category: "IPFS COMMANDS",
		Description: ``,
		Subcommands: []cli.Command{
			{
				Name:   "add",
				Usage:  "Save content to ipfs",
				Action: utils.MigrateFlags(ipfsAdd),
				Flags: []cli.Flag{
//					utils.IPFSStringFlag,
				},
				Description: ``,
			},
			{
				Name:   "cat",
				Usage:  "Read content from ipfs",
				Action: utils.MigrateFlags(ipfsCat),
				Flags: []cli.Flag{
//					utils.IPFSHashFlag,
				},
				Description: ``,
			},
		},
	}
)

func ipfsAdd(ctx *cli.Context) error {

	args := ctx.Args()

	if len(args) != 1 {
		utils.Fatalf(`Usage: geth ipfs add <string>`)
	}

	stringToIPFS := args[0]
	//stringToIPFS := "Hello IPFS Shell tests"

	s := shell.NewShell(shellUrl)

	mhash, err := s.Add(strings.NewReader(stringToIPFS))
	if err != nil{
		return err
	}
	fmt.Printf("%v", mhash)
	return nil
}

func ipfsCat(ctx *cli.Context) error {

	args := ctx.Args()
	if len(args) != 1 {
		utils.Fatalf(`Usage: geth ipfs cat <file hash>`)
	}

	hash := args[0]

	s := shell.NewShell(shellUrl)

	rc, err := s.Cat(fmt.Sprintf("/ipfs/%s", hash))

	if err != nil{
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(rc)
	fileBytes := buf.String()

	fmt.Printf("%+v", fileBytes)

	return nil
}