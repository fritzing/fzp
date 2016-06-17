package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var (
	verbose bool
)

func main() {
	verbose = false
	app := cli.NewApp()
	app.Name = "fzp"
	app.Usage = "fzp tool (validator, encoder)"
	app.Version = "0.2.4"
	app.Email = "https://github.com/fritzing/fzp"
	app.Commands = []cli.Command{
		{
			Name:   "validate",
			Usage:  "validate fzp file/files",
			Flags:  commandValidateFlags,
			Action: commandValidateAction,
		},
		{
			Name:   "encode",
			Usage:  "read a fzp file and encode to json",
			Flags:  commandEncodeFlags,
			Action: commandEncodeAction,
		},
		{
			Name:   "format",
			Usage:  "read a fzp file and format it",
			Flags:  commandFormatFlags,
			Action: commandFormatAction,
		},
		{
			Name:   "create",
			Usage:  "create a new template fzp file",
			Flags:  commandCreateFlags,
			Action: commandCreateAction,
		},
	}
	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
		return
	}
	app.Run(os.Args)
}

func Logf(format string, a ...interface{}) {
	if verbose {
		fmt.Printf(format, a...)
	}
}
