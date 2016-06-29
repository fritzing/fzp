package main

import (
	"testing"
)

func TestMain(t *testing.T) {

}

func Test_dataHandler_dir(t *testing.T) {
	dataHandler("./fixture", dirHandler, fileHandler)
}

func Test_dataHandler_file(t *testing.T) {
	dataHandler("./fixture/template.fzp", dirHandler, fileHandler)
}

func dirHandler() {}

func fileHandler() {}
