package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/fritzing/fzp/src/go"
	"github.com/pmezard/go-difflib/difflib"
	"github.com/urfave/cli"
)

var commandFormatFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:  "write, w",
		Usage: "write result to file instead of stdout",
	},
	&cli.BoolFlag{
		Name:  "diff, d",
		Usage: "display diffs instead of stdout or write file",
	},
}

// the format action is similar to the gofmt command.
// it can be used to format the source of a fzp file.
func commandFormatAction(c *cli.Context) error {
	tmpArgs := c.Args()
	if tmpArgs.Len() == 0 {
		fmt.Println("missing source to format. try run")
		fmt.Println("  fzp format part.fzp")
		os.Exit(127)
	}

	source := tmpArgs.Get(0)

	formatDirHandler := func() {
		err := formatFile(source, c.Bool("write"), c.Bool("diff"))
		if err != nil {
			fmt.Println("Error", err)
			os.Exit(127)
		}
	}

	formatFileHandler := func() {
		err := formatDir(source, c.Bool("write"), c.Bool("diff"))
		if err != nil {
			fmt.Println("Error", err)
			os.Exit(127)
		}
	}

	dataHandler(source, formatFileHandler, formatDirHandler)
	return nil
}

func formatFile(filepath string, runWrite, runDiff bool) error {
	format, _ := fzp.GetFormat(filepath)

	if format != fzp.FormatFzp {
		return errors.New("at the moment only fzp format supported")
	}

	fzpFile, fzpBytes, err := fzp.ReadFzp(filepath)
	if err != nil {
		fmt.Println(err)
		return err
	}

	formattedXML, err := fzpFile.ToXML()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if runWrite {
		err := ioutil.WriteFile(filepath, formattedXML, 0755)
		if err != nil {
			fmt.Println("Error", err)
			os.Exit(127)
		}
		return nil
	}

	// diff
	if runDiff {
		diff := difflib.ContextDiff{
			A:        difflib.SplitLines(string(fzpBytes)),
			B:        difflib.SplitLines(string(formattedXML)),
			FromFile: filepath,
			ToFile:   "Current",
			Context:  3,
			Eol:      "\n",
		}
		result, _ := difflib.GetContextDiffString(diff)
		fmt.Printf(strings.Replace(result, "\t", " ", -1))
		return nil
	}

	fmt.Println(string(formattedXML))
	return nil
}

func formatDir(filepath string, runWrite, runDiff bool) error {
	dirData, err := ioutil.ReadDir(filepath)
	if err != nil {
		return err
	}
	for _, v := range dirData {
		fmt.Println("i", v.Name())
		formatFile(path.Join(filepath, v.Name()), runWrite, runDiff)
	}
	return nil
}
