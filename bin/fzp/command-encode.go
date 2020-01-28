package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/fritzing/fzp/src/go"
	"github.com/urfave/cli"
)

var commandEncodeFlags = []cli.Flag{
	&cli.StringFlag{
		Name:  "input, i",
		Usage: "the input file. supported formats: xml, json and yaml",
	},
	&cli.StringFlag{
		Name:  "output, o",
		Usage: "the output file or format. supported formats: xml, json and yaml",
	},
}

func commandEncodeAction(c *cli.Context) error {
	flagInput := c.String("input")
	flagOutput := c.String("output")
	fmt.Println("in", flagInput, "out", flagOutput)

	if flagInput == "" {
		fmt.Println("missing -input flag")
		os.Exit(1)
	}
	if flagOutput == "" {
		fmt.Println("missing -output flag")
		os.Exit(1)
	}

	encodeDirHandler := func() {
		if err := encodeDir(flagInput, flagOutput); err != nil {
			fmt.Println("Error", err)
			os.Exit(127)
		}
	}

	encodeFileHandler := func() {
		if err := encodeFile(flagInput, flagOutput); err != nil {
			fmt.Println("Error", err)
			os.Exit(127)
		}
	}

	dataHandler(flagInput, encodeDirHandler, encodeFileHandler)
	return nil
}

func encodeFile(source, out string) error {
	// fmt.Println("encode file", source, out)

	// read
	format, _ := fzp.GetFormat(source)
	if format != fzp.FormatFzp {
		return errors.New("Error: input file '" + source + "' is not a fzp file!")
	}
	loadedFzp, _, err := fzp.ReadFzp(source)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// write := false

	// format
	outFormat, isOutFile := fzp.GetFormat(out)
	// fmt.Println("outFormat", outFormat)

	formatted, err := loadedFzp.Marshal(outFormat)
	if err != nil {
		return err
	}

	if isOutFile {
		err := ioutil.WriteFile(out, formatted, 0755)
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf("file %q successful written\n", out)
		return nil
	}

	// if not write, print to stdout
	fmt.Println(string(formatted))
	return nil
}

func encodeDir(source, format string) error {
	dirData, err := ioutil.ReadDir(source)
	if err != nil {
		return err
	}
	for i, v := range dirData {
		fmt.Println(i, v.Name())
		tmpFile := path.Join(source, v.Name())
		format, _ := fzp.GetFormat(tmpFile)
		if format == fzp.FormatFzp {
			err := encodeFile(tmpFile, "format") //, runWrite, runDiff)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}
