package main

import (
	"github.com/wukongcheng/databank/cmd/utils"
	"gopkg.in/urfave/cli.v1"
	"github.com/ipfs/go-ipfs-api"
	"fmt"
	"os"
	"bufio"
	"bytes"
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
				Action: utils.MigrateFlags(ipfsAddFile),
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
			{
				Name:   "get",
				Usage:  "Read content from ipfs to specified file",
				Action: utils.MigrateFlags(ipfsGetFile),
				Flags: []cli.Flag{
//					utils.IPFSHashFlag,
				},
				Description: ``,
			},
		},
	}
)

func ipfsAddFile(ctx *cli.Context) error {

	args := ctx.Args()

	if len(args) != 1 {
		utils.Fatalf(`Usage: xcare ipfs add <file path>`)
	}

	filePath := args[0]
	//stringToIPFS := "Hello IPFS Shell tests"
	file, err := os.Open(filePath)
	if err != nil{
		return err
	}

	s := shell.NewShell(shellUrl)

	mhash, err := s.Add(bufio.NewReader(file))
	if err != nil{
		return err
	}
	fmt.Printf("%v", mhash)
	return nil
}

func ipfsCat(ctx *cli.Context) error {

	args := ctx.Args()
	if len(args) != 1 {
		utils.Fatalf(`Usage: xcare ipfs cat <file hash>`)
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

func ipfsGetFile(ctx *cli.Context) error {

	args := ctx.Args()
	if len(args) != 2 {
		utils.Fatalf(`Usage: xcare ipfs get <file hash> <targe file path>`)
	}

	hash := args[0]

	targetFile := args[1]

	s := shell.NewShell(shellUrl)

	err := s.Get(hash,targetFile)

	if err != nil{
		return err
	}

	return nil
}