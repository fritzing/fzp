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
	verbose   bool
	checkData bool
)

func main() {
	verbose = false
	app := cli.NewApp()
	app.Name = "validator"
	app.Usage = "fzp validator"
	app.Version = "0.1.1"
	app.Author = "paul vollmer"
	app.Email = "https://github.com/paulvollmer/fzp"
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
			Name:  "no-check, c",
			Usage: "not check the data",
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
	fzpFile := c.String("file")
	fzpDir := c.String("dir")
	verbose = c.Bool("verbose")
	checkData = !c.Bool("no-check")
	// process data
	if fzpFile != "" {
		validateFile(fzpFile)
		os.Exit(0)
	} else if fzpDir != "" {
		Log("read folder '%v'\n", fzpDir)
		validateFolder(fzpDir)
		os.Exit(0)
	} else {
		cli.ShowAppHelp(c)
		fmt.Println("USAGE-SAMPLES")
		fmt.Println("")
		fmt.Printf("   $ %v --file file/path.fzp\n", c.App.Name)
		fmt.Println("   # or")
		fmt.Printf("   $ %v -f file/path.fzp\n", c.App.Name)
		fmt.Println("")
		fmt.Printf("   $ %v --dir file/dir\n", c.App.Name)
		fmt.Println("   # or")
		fmt.Printf("   $ %v -d file/dir\n", c.App.Name)
		fmt.Println("")
		fmt.Println("also you can combine the other flags (no-check, verbose)")
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

	if checkData {
		errCheckTags := fzpData.CheckTags()
		if errCheckTags != nil {
			fmt.Println("Check 'tags' Error @", src)
			fmt.Println(errCheckTags)
		}

		errCheck := fzp.CheckData(fzpData)
		if errCheck != nil {
			fmt.Println("Error @", src)
			for _, v := range errCheck {
				fmt.Println("=>", v)
			}
			return
		}
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
		if isExtFzp(filename) {
			validateFile(src + "/" + filename)
		}
	}
}

func isExtFzp(src string) bool {
	if filepath.Ext(src) == ".fzp" {
		return true
	} else {
		return false
	}
}
