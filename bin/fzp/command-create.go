package main

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/fritzing/fzp/src/go"
	"github.com/urfave/cli"
)

var commandCreateFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "title, t",
		Usage: "fzp title",
		Value: "untitled",
	},
	cli.StringFlag{
		Name:  "family, f",
		Usage: "set a familyname property for the part",
	},
	cli.StringFlag{
		Name:  "version, v",
		Usage: "fzp version",
		Value: "0.1.0",
	},
	cli.StringFlag{
		Name:  "desc, d",
		Usage: "fzp description",
	},
	cli.StringFlag{
		Name:  "author, a",
		Usage: "author of the fritzing part",
		Value: "unknown",
	},
	cli.StringFlag{
		Name:  "date, D",
		Usage: "creation date of the part",
	},
	cli.StringFlag{
		Name:  "url, u",
		Usage: "parts url",
	},
	cli.StringFlag{
		Name:  "label, l",
		Usage: "set a label for the part",
	},
	cli.StringFlag{
		Name:  "tags, T",
		Usage: "a list of part tags",
	},
	// TODO: add all flags to setup the fzp data
}

func commandCreateAction(c *cli.Context) error {
	// get the flag data
	flagTitle := c.String("title")
	flagFamily := c.String("family")
	flagVersion := c.String("version")
	flagDesc := c.String("desc")
	flagAuthor := c.String("author")
	flagDate := c.String("date")
	if flagDate == "" {
		now := time.Now()
		flagDate = now.Format(time.UnixDate)
	}
	flagURL := c.String("url")
	flagLabel := c.String("label")
	// flagTags := c.String("tags")
	// flagProperties := c.String("properties")
	// flagViews := c.String("views")
	// flagConnectors := c.String("connectors")
	// flagBusses := c.String("busses")

	// create a new fzp object
	tmpFzp := fzp.NewFzp("unknown", flagTitle, flagFamily)
	// tmpFzp.FritzingVersion =
	// tmpFzp.ModuleId =
	// tmpFzp.ReferenceFile =
	tmpFzp.Title = flagTitle
	tmpFzp.Version = flagVersion
	tmpFzp.Description = flagDesc
	tmpFzp.Author = flagAuthor
	tmpFzp.Date = flagDate
	tmpFzp.URL = flagURL
	tmpFzp.Label = flagLabel
	// tmpFzp.Tags = flagTags
	// tmpFzp.Properties = flagProperties
	//fmt.Println("fzp", tmpFzp)
	// tmpFzp.PrettyPrint()

	// output
	tmpArgs := c.Args()
	tmpArgsLen := len(tmpArgs)
	// fmt.Println("args", tmpArgs)

	format := fzp.FormatFzp
	isFile := false
	if tmpArgsLen > 0 {
		format, isFile = fzp.GetFormat(tmpArgs[0])
	}

	// format data
	tmpFzpEncoded, err := tmpFzp.Marshal(format)
	if err != nil {
		cli.NewExitError("Encode Error: "+err.Error(), 127)
	}

	if !isFile {
		os.Stdout.Write(tmpFzpEncoded)
		return nil
	}

	err = ioutil.WriteFile(tmpArgs[0], tmpFzpEncoded, 0755)
	if err != nil {
		cli.NewExitError("WriteError: "+err.Error(), 127)
	}

	return nil
}
