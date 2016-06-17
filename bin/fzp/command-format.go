package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fritzing/fzp/src/go"
	"github.com/pmezard/go-difflib/difflib"
	"github.com/urfave/cli"
)

var commandFormatFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "write, w",
		Usage: "write result to file instead of stdout",
	},
	cli.BoolFlag{
		Name:  "diff, d",
		Usage: "display diffs instead of stdout or write file",
	},
}

// the format action is similar to the gofmt command.
// it can be used to format the source of a fzp file.
func commandFormatAction(c *cli.Context) error {
	tmpArgs := c.Args()
	if len(tmpArgs) == 0 {
		fmt.Println("missing source to format. try run")
		fmt.Println("  fzp format part.fzp")
		os.Exit(127)
	}

	// check if the source has an fzp suffix
	source := tmpArgs[0]
	if !fzp.HasExtFzp(source) {
		fmt.Printf("the source %q is not a %s file\n", source, fzp.ExtFzp)
		os.Exit(127)
	}

	fzpFile, fzpBytes, err := fzp.ReadFzp(source)
	if err != nil {
		fmt.Println(err)
		return err
	}

	formattedXML, err := fzpFile.ToXML()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if c.Bool("write") {
		err := ioutil.WriteFile(source, formattedXML, 0755)
		if err != nil {
			fmt.Println("Error", err)
			os.Exit(127)
		}
		return nil
	}

	// diff
	if c.Bool("diff") {
		diff := difflib.ContextDiff{
			A:        difflib.SplitLines(string(fzpBytes)),
			B:        difflib.SplitLines(string(formattedXML)),
			FromFile: source,
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
