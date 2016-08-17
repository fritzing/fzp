package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "fzp"
	app.Usage = "fzp tool (validator, encoder, formatter)"
	app.Version = "0.2.5a"
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
			Name:   "get",
			Usage:  "get data from a fzp file",
			Flags:  commandGetFlags,
			Action: commandGetAction,
		},
		{
			Name:   "create",
			Usage:  "create a new fzp dataset",
			Flags:  commandCreateFlags,
			Action: commandCreateAction,
		},
	}
	// cli.HelpPrinter = func(w io.Writer, templ string, data interface{}) {
	// 	fmt.Println("TODO: better main help")
	// }
	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
		return
	}
	app.Run(os.Args)
}

// check if the source is a directory of a file and call a handler func
func dataHandler(source string, dir, file func()) {
	finfo, err := os.Stat(source)
	if err != nil { // no such file or dir
		cli.NewExitError(err.Error(), 1)
		os.Exit(127)
	}
	if finfo.IsDir() {
		dir()
	} else {
		file()
	}
}
