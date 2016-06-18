package main

import (
	"fmt"
	"os"

	"github.com/fritzing/fzp/src/go"
	"github.com/urfave/cli"
)

var commandGetFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "author",
		Usage: "get the author data",
	},
}

func commandGetAction(c *cli.Context) error {
	tmpArgs := c.Args()
	fmt.Println("ARGS", tmpArgs)

	if len(tmpArgs) == 0 {
		fmt.Println("Missing fritzing part filepath")
		os.Exit(1)
	}
	dataHandler(tmpArgs[0], func() {
		getActionDir(tmpArgs[0])
	}, func() {
		getActionFile(tmpArgs[0], c.Bool("author"))
	})

	return nil
}

func getActionFile(src string, a bool) {
	fmt.Println("get file data...")
	part, _, err := fzp.ReadFzp(src)
	if err != nil {
		fmt.Println(err)
		os.Exit(127)
	}
	fmt.Printf("%# v\n", part)
}

func getActionDir(src string) {
	fmt.Println("get folder data...")
}
