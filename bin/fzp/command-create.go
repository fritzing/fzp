package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/fritzing/fzp/src/go"
	"github.com/urfave/cli"
)

var commandCreateFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "f",
		Usage: "fzp data format (default is xml)",
		Value: "xml",
	},
	cli.StringFlag{
		Name:  "title, t",
		Usage: "fzp title",
		Value: "untitled",
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
		Usage: "", //TODO: ...
	},
	cli.StringFlag{
		Name:  "label, l",
		Usage: "", //TODO: ...
	},
	// TODO: add all flags to setup the fzp data
}

func commandCreateAction(c *cli.Context) error {
	// get the flag data
	flagFormat := c.String("f")
	flagTitle := c.String("title")
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

	// create a new fzp object
	tmpFzp := fzp.Fzp{}
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

	// format data
	tmpFzpEncoded, err := tmpFzp.Marshal(flagFormat)

	// output
	tmpArgs := c.Args()
	tmpArgsLen := len(tmpArgs)
	// fmt.Println("args", tmpArgs)
	if tmpArgsLen > 0 {
		err = ioutil.WriteFile(tmpArgs[0], tmpFzpEncoded, 0755)
		if err != nil {
			fmt.Println("WriteError:", err)
			os.Exit(129)
		}

	} else {
		fmt.Println(string(tmpFzpEncoded))
	}

	return nil
}
