package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/fritzing/fzp/src/go"
	"github.com/urfave/cli"
)

var commandValidateFlags = []cli.Flag{
	// input flags
	cli.StringFlag{
		Name:  "file, f",
		Usage: "fzp filepath",
	},
	cli.StringFlag{
		Name:  "dir, d",
		Usage: "fzp directory",
	},

	// data check settings
	cli.BoolFlag{
		Name:  "no-check-fritzingversion, nf",
		Usage: "disable fritzingVersion check",
	},
	cli.BoolFlag{
		Name:  "no-check-moduleid, nm",
		Usage: "disable moduleid check",
	},
	cli.BoolFlag{
		Name:  "no-check-referencefile, nr",
		Usage: "disable referenceFile check",
	},
	cli.BoolFlag{
		Name:  "no-check-version, nv",
		Usage: "disable <version> check",
	},
	cli.BoolFlag{
		Name:  "no-check-title, nt",
		Usage: "disable <title> check",
	},
	cli.BoolFlag{
		Name:  "no-check-description, nd",
		Usage: "disable <description> check",
	},
	cli.BoolFlag{
		Name:  "no-check-author, na",
		Usage: "disable <author> check",
	},
	// cli.BoolFlag{
	// 	Name:  "no-check-date, nD",
	// 	Usage: "disable <date> check",
	// },
	// cli.BoolFlag{
	// 	Name:  "no-check-url, nD",
	// 	Usage: "disable <url> check",
	// },
	// cli.BoolFlag{
	// 	Name:  "no-check-label, nD",
	// 	Usage: "disable <label> check",
	// },
	// cli.BoolFlag{
	// 	Name:  "no-check-taxonomy, nD",
	// 	Usage: "disable <taxonomy> check",
	// },
	cli.BoolFlag{
		Name:  "no-check-family, nD",
		Usage: "disable <family> check",
	},
	// cli.BoolFlag{
	// 	Name:  "no-check-variant, nD",
	// 	Usage: "disable <variant> check",
	// },
	cli.BoolFlag{
		Name:  "no-check-tags, nT",
		Usage: "disable <tags> check",
	},
	cli.BoolFlag{
		Name:  "no-check-properties, np",
		Usage: "disable <properties> check",
	},
	cli.BoolFlag{
		Name:  "no-check-views, nV",
		Usage: "disable <views> check",
	},
	cli.BoolFlag{
		Name:  "no-check-connectors, nc",
		Usage: "disable <connectors> check",
	},
	cli.BoolFlag{
		Name:  "no-check-buses, nb",
		Usage: "disable <buses> check",
	},
	// utils
	cli.BoolFlag{
		Name:  "verbose, V",
		Usage: "verbose mode",
	},
}

func commandValidateAction(c *cli.Context) error {
	fmt.Println("validate fzp file")
	// get cli flag value
	fzpFile := c.String("file")
	fzpDir := c.String("dir")

	// process data
	if fzpFile != "" {
		if err := validateFile(c, fzpFile); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	} else if fzpDir != "" {
		// Logf("read folder '%v'\n", fzpDir)
		if err := validateFolder(c, fzpDir); err != nil {
			os.Exit(1)
		}
		os.Exit(0)
	} else {
		cli.ShowSubcommandHelp(c)
		fmt.Println("USAGE-SAMPLES")
		fmt.Println("")
		fmt.Printf("   $ %v validate --file file/path.fzp\n", c.App.Name)
		fmt.Println("   # or")
		fmt.Printf("   $ %v validate -f file/path.fzp\n", c.App.Name)
		fmt.Println("")
		fmt.Printf("   $ %v validate --dir file/dir\n", c.App.Name)
		fmt.Println("   # or")
		fmt.Printf("   $ %v validate -d file/dir\n", c.App.Name)
		fmt.Println("")
		fmt.Println("also you can combine the other flags (no-check, verbose)")
	}
	return nil
}

func validateFile(c *cli.Context, src string) error {
	fzpData, _, err := fzp.ReadFzp(src)
	if err != nil {
		fmt.Printf("validator failed @ %v\n", err)
		os.Exit(1)
	}
	// Logf("fzp file '%v' successful read\n", src, fzpData)

	errCounter := checkData(c, fzpData)
	if errCounter != 0 {
		return errors.New(strconv.Itoa(errCounter) + " Errors @ " + src)
	}

	// Logf("fzp valid\n")
	return nil
}

func validateFolder(c *cli.Context, src string) []error {
	var errList []error
	folderFiles, err := ioutil.ReadDir(src)
	if err != nil {
		errList = append(errList, errors.New("validator failed @ read folder '"+src+"'"))
		return errList
	}

	for _, v := range folderFiles {
		filename := v.Name()
		// fmt.Printf("file %v: %v\n", k, filename)
		// check if file is a fzp file
		// if fzp.HasExtFzp(filename) {
		if err := validateFile(c, src+"/"+filename); err != nil {
			errList = append(errList, err)
			fmt.Println(err)
			// }
		}
	}
	return errList
}

func checkData(c *cli.Context, fzpData fzp.Fzp) int {
	checkErrorCounter := 0

	/*if !c.Bool("no-check-fritzingversion") {
		if err := fzpData.CheckFritzingVersion(); err != nil {
			fmt.Println("=>", err)
			checkErrorCounter++
		}
	}*/

	if !c.Bool("no-check-moduleid") {
		if err := fzpData.CheckModuleID(); err != nil {
			fmt.Println("=>", err)
			checkErrorCounter++
		}
	}

	// ReferenceFile

	/*if !c.Bool("no-check-version") {
		if err := fzpData.CheckVersion(); err != nil {
			fmt.Println("=>", err)
			checkErrorCounter++
		}
	}*/

	if !c.Bool("no-check-title") {
		if err := fzpData.CheckTitle(); err != nil {
			fmt.Println("=>", err)
			checkErrorCounter++
		}
	}

	/*if !c.Bool("no-check-family") {
		if err := fzpData.CheckFamily(); err != nil {
			fmt.Println("=>", err)
			checkErrorCounter++
		}
	}*/

	if !c.Bool("no-check-description") {
		if err := fzpData.CheckDescription(); err != nil {
			fmt.Println("=>", err)
			checkErrorCounter++
		}
	}

	if !c.Bool("no-check-author") {
		if err := fzpData.CheckAuthor(); err != nil {
			fmt.Println("=>", err)
			checkErrorCounter++
		}
	}

	// Check Date ?
	// Check URL ?
	// Check Label ?
	// Check Taxonomy ?
	// Check Family ?
	// Check Variant ?

	/*if !c.Bool("no-check-tags") {
		if err := fzpData.CheckTags(); err != nil {
			fmt.Println("=>", err)
			checkErrorCounter++
		}
	}*/

	if !c.Bool("no-check-properties") {
		if err := fzpData.CheckProperties(); err != nil {
			fmt.Println("=>", err)
			checkErrorCounter++
		}
	}

	if !c.Bool("no-check-buses") {
		if err := fzpData.CheckBuses(); err != nil {
			fmt.Println("=>", err)
			checkErrorCounter++
		}
	}

	return checkErrorCounter
}
