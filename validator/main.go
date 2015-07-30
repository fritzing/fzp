package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/paulvollmer/fzp/go"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	verbose bool
)

func main() {
	verbose = false
	app := cli.NewApp()
	app.Name = "fzp-validator"
	app.Usage = "fzp validator"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "file, f",
			Value: "",
			Usage: "the fzp filepath",
		},
		cli.StringFlag{
			Name:  "dir, d",
			Value: "",
			Usage: "the fzp files directory",
		},
		cli.BoolFlag{
			Name:  "verbose, V",
			Usage: "verbose mode",
		},
	}
	app.Action = cliAction
	app.Run(os.Args)
}

func cliAction(c *cli.Context) {
	// get cli flag value
	verbose = c.Bool("verbose")
	fzpFile := c.String("file")
	fzpDir := c.String("dir")
	// process data
	if fzpFile != "" {
		validateFile(fzpFile)
		os.Exit(0)
	} else if fzpDir != "" {
		Log("read folder '%v'\n", fzpDir)
		validateFolder(fzpDir)
		os.Exit(0)
	} else {
		fmt.Println("validate a file or a directory...")
		fmt.Println("  $validate --file file/path.fzp")
		fmt.Println("  $validate --dir file/dir")
	}
}

func Log(format string, a ...interface{}) {
	if verbose {
		fmt.Printf(format, a...)
	}
}

func validateFile(src string) {
	fzpData, err := fzp.ReadFzp(src)
	if err != nil {
		fmt.Printf("validator failed @ %v\n", err)
		os.Exit(1)
	}
	Log("fzp file '%v' successful read\n", src)
	errCheck := fzp.CheckData(fzpData)
	if errCheck != nil {
		fmt.Println("Error @", src)
		for _, v := range errCheck {
			fmt.Println("=>", v)
		}
		return
	}
	Log("fzp valid\n")
}

func validateFolder(src string) {
	folderFiles, err := ioutil.ReadDir(src)
	if err != nil {
		fmt.Printf("validator failed @ read folder '%v'\n", src)
		os.Exit(2)
	}
	for _, v := range folderFiles {
		filename := v.Name()
		// fmt.Printf("file %v: %v\n", k, filename)
		// check if file is a fzp file
		if filepath.Ext(filename) == ".fzp" {
			validateFile(src + "/" + filename)
		}
	}
}
