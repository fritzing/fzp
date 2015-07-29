package main

import (
	"encoding/xml"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/kr/pretty"
	"io/ioutil"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "fzp-validator"
	app.Usage = "fzp validator"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "file, f",
			Value: "*.fzp",
			Usage: "the fzp filepath",
		},
	}
	app.Action = cliAction
	app.Run(os.Args)
}

func cliAction(c *cli.Context) {
	// get cli flag value
	fzpFile := c.String("file")

	// readFile
	fzpBytes, err := ioutil.ReadFile(fzpFile)
	if err != nil {
		panic(err)
	}

	// decode XML
	fzp := Fzp{}
	xml.Unmarshal(fzpBytes, &fzp)

	fmt.Printf("%# v", pretty.Formatter(fzp))
}
