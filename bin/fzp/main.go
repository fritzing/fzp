package main

import (
	// "errors"
	// "fmt"
	"github.com/codegangsta/cli"
	// "github.com/fritzing/fzp/src/go"
	// "io/ioutil"
	// "strconv"
	"os"
)

var (
	verbose bool
)

func main() {
	verbose = false
	app := cli.NewApp()
	app.Name = "fzp"
	app.Usage = "fzp tool (validator, encoder)"
	app.Version = "0.2.3"
	app.Email = "https://github.com/fritzing/fzp"
	app.Commands = []cli.Command{
		{
			Name:   "validate",
			Usage:  "validate fzp file/files",
			Flags:  commandValidateFlags,
			Action: commandValidateAction,
		},
		{
			Name:   "create",
			Usage:  "create a new template fzp file",
			Flags:  commandCreateFlags,
			Action: commandCreateAction,
		},
		{
			Name:   "encode",
			Usage:  "read fzp file and encode to json",
			Flags:  commandEncodeFlags,
			Action: commandEncodeAction,
		},
	}
	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
		os.Exit(0)
	}
	app.Run(os.Args)
}
