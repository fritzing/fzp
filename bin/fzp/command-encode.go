package main

import (
	"fmt"
	"os"

	"github.com/fritzing/fzp/src/go"
	"github.com/urfave/cli"
)

var commandEncodeFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "input, i",
		Usage: "the input file. supported formats: xml, json and yaml",
	},
	cli.StringFlag{
		Name:  "output, o",
		Usage: "the output file or format. supported formats: xml, json and yaml",
	},
}

func commandEncodeAction(c *cli.Context) error {
	fmt.Println("encode file")

	flagInput := c.String("input")
	flagOutput := c.String("output")
	fmt.Println("in", flagInput, "out", flagOutput)
	// read
	loadedFzp, _, err := fzp.ReadFzp(flagInput)
	if err != nil {
		return err
	}

	// format
	var data []byte
	switch flagOutput {
	case "":
		fmt.Println("missing out format or filepath. try xml, json or yaml")
		os.Exit(127)

	case "xml":
		data, err = loadedFzp.ToXML()
		break

	case "json":
		data, err = loadedFzp.ToJSON()
		break

	case "yaml":
		data, err = loadedFzp.ToYAML()
		break
	}

	// return nil
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
