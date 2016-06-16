package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	// "github.com/fritzing/fzp/src/go"
)

var commandEncodeFlags = []cli.Flag{}

func commandEncodeAction(c *cli.Context) {
	fmt.Println("encode fzp file")
}
