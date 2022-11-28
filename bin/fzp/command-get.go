package main

import (
	"fmt"
	"github.com/fritzing/fzp/src/go"
	"github.com/urfave/cli"
	"os"
	"strconv"
)

var commandGetFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:  "title, t",
		Usage: "get the title data",
	},
	&cli.BoolFlag{
		Name:  "version, v",
		Usage: "get the version data",
	},
	&cli.BoolFlag{
		Name:  "desc, d",
		Usage: "get the description data",
	},
	&cli.BoolFlag{
		Name:  "author, a",
		Usage: "get the author data",
	},
	&cli.BoolFlag{
		Name:  "date, D",
		Usage: "get the date data",
	},
	&cli.BoolFlag{
		Name:  "url, u",
		Usage: "get the url data",
	},
	&cli.BoolFlag{
		Name:  "label, l",
		Usage: "get the label data",
	},
	&cli.StringFlag{
		Name:  "tags, T",
		Usage: "get the tags data",
	},
	&cli.StringFlag{
		Name:  "props, p",
		Usage: "get the properties data",
	},
	&cli.StringFlag{
		Name:  "views, V",
		Usage: "get the views data",
	},
	&cli.StringFlag{
		Name:  "connectors, c",
		Usage: "get the connectorsr data",
	},
	&cli.StringFlag{
		Name:  "buses, b",
		Usage: "get the buses data",
	},
}

func commandGetAction(c *cli.Context) error {
	tmpArgs := c.Args()
	// fmt.Println("ARGS", tmpArgs)

	if len(tmpArgs) == 0 {
		fmt.Println("Missing fritzing part filepath")
		os.Exit(1)
	}

	getActionFile := func(src string) {
		// fmt.Println("get file data...")
		part, _, err := fzp.ReadFzp(src)
		if err != nil {
			fmt.Println(err)
			os.Exit(127)
		}
		// fmt.Printf("%# v\n", part)

		if c.Bool("title") {
			fmt.Println(part.Title)
		}

		if c.Bool("version") {
			fmt.Println(part.Version)
		}

		if c.Bool("descritino") {
			fmt.Println(part.Description)
		}

		if c.Bool("author") {
			fmt.Println(part.Author)
		}

		if c.Bool("date") {
			fmt.Println(part.Date)
		}

		if c.Bool("url") {
			fmt.Println(part.URL)
		}

		if c.Bool("label") {
			fmt.Println(part.Label)
		}

		flagTags := c.String("tags")
		if flagTags != "" {
			totalTags := len(part.Tags)
			if flagTags == "all" {
				for i := 0; i < totalTags; i++ {
					fmt.Println(i, part.Tags[i])
				}
			} else {
				// is the flag an indice
				flagTagsInt, err := strconv.Atoi(flagTags)
				if err != nil {
					fmt.Println("tags flag is not an integer. please use numbers to query a tag item")
					os.Exit(1)
				}
				if flagTagsInt > totalTags {
					fmt.Println("tag does not exist! total number of tags:", totalTags)
				}
				fmt.Println(part.Tags[flagTagsInt])
			}
		}

		flagProps := c.String("props")
		if flagProps != "" {
			if flagProps == "all" {
				for i := 0; i < len(part.Properties); i++ {
					fmt.Println(i, part.Properties[i].Name, part.Properties[i].Value)
				}
			} else {
				tmp, err := part.Properties.GetValue(flagProps)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				fmt.Println(tmp)
			}
		}

		flagViews := c.String("views")
		if flagViews != "" {
			// for i := 0; i < len(part.Views); i++ {
			// 	fmt.Println(i, part.Views[i])
			// }
		}

		flagConnectors := c.String("connectors")
		if flagConnectors != "" {
			fmt.Println(part.Connectors)
		}

		flagBuses := c.String("buses")
		if flagBuses != "" {
			fmt.Println(part.Buses)
		}
	}

	getActionDir := func(src string) {
		fmt.Println("TODO: get folder data...")
	}

	dataHandler(tmpArgs.Get(0), func() {
		// if the source is a directory...
		getActionDir(tmpArgs.Get(0))

	}, func() {
		// if the source is a file...
		getActionFile(tmpArgs.Get(0))
	})

	return nil
}
