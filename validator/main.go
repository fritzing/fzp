package main

import (
	"encoding/xml"
	"fmt"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
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
			Name:  "folder, d",
			Value: "",
			Usage: "the path to a folder of fzp files",
		},
	}
	app.Action = cliAction
	app.Run(os.Args)
}

func cliAction(c *cli.Context) {
	// get cli flag value
	fzpFile := c.String("file")
	fzpFolder := c.String("folder")
	// read fzp file
	if fzpFile != "" {
		validateFile(fzpFile)
		os.Exit(0)
	}
	// read folder
	if fzpFolder != "" {
		validateFolder(fzpFolder)
		os.Exit(0)
	}
}

func validateFile(src string) {
	// read
	fzpBytes, err := ioutil.ReadFile(src)
	if err != nil {
		fmt.Printf("Read fzp file '%v' Error: %v\n", src, err)
		os.Exit(10)
	}
	// decode XML
	fzp := Fzp{}
	errDecode := xml.Unmarshal(fzpBytes, &fzp)
	if errDecode != nil {
		fmt.Printf("Decode Error at file '%v': %v\n", src, errDecode)
		os.Exit(11)
	}
}

func validateFolder(src string) {
	folderFiles, err := ioutil.ReadDir(src)
	if err != nil {
		fmt.Printf("Error read folder '%v'\n", src)
		os.Exit(20)
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
