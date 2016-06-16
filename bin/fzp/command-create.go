package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/fritzing/fzp/src/go"
)

var commandCreateFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "title, t",
		Usage: "fzp title",
	},
	cli.StringFlag{
		Name:  "version, v",
		Usage: "fzp version",
	},
	cli.StringFlag{
		Name:  "desc",
		Usage: "fzp desc",
	},
}

func commandCreateAction(c *cli.Context) {
	fmt.Println("create fzp file")

	// input
	tmpArgs := c.Args()
	tmpArgsLen := len(tmpArgs)
	fmt.Println("args", tmpArgs)

	tmpTitle := c.String("title")
	if tmpTitle == "" {
		tmpTitle = "input here"
	}

	// process
	// create
	tmpFzp := fzp.Fzp{}
	tmpFzp.Title = tmpTitle
	fmt.Println("fzp", tmpFzp)

	// tmpFzp.PrettyPrint()
	fmt.Println("FritzingVersion =", tmpFzp.FritzingVersion)
	fmt.Println("ModuleId        =", tmpFzp.ModuleId)
	fmt.Println("ReferenceFile   =", tmpFzp.ReferenceFile)
	fmt.Println("Version         =", tmpFzp.Version)
	fmt.Println("Title           =", tmpFzp.Title)
	fmt.Println("Description     =", tmpFzp.Description)
	fmt.Println("Author          =", tmpFzp.Author)

	// output
	if tmpArgsLen > 0 {
		fmt.Println("write to ", tmpArgs)
		// fzp.Write()
	}
}
