package fzp

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"path/filepath"
)

type Fzp struct {
	XMLName         xml.Name `xml:"module"`
	FritzingVersion string   `xml:"fritzingVersion,attr"`
	ModuleId        string   `xml:"moduleId,attr"`
	ReferenceFile   string   `xml:"referenceFile,attr"`
	Version         string   `xml:"version"`
	Title           string   `xml:"title"`
	Description     string   `xml:"description"`
	Author          string   `xml:"author"`
	Date            string   `xml:"date"`
	URL             string   `xml:"url"`
	Label           string   `xml:"label"`
	//Taxonomy        string      `xml:"taxonomy"`
	//Family          string      `xml:"family"`
	//Variant         string      `xml:"variant"`
	Tags       []string    `xml:"tags>tag"`
	Properties []Property  `xml:"properties>property"`
	Views      Views       `xml:"views"`
	Connectors []Connector `xml:"connectors>connector"`
	Buses      []Bus       `xml:"buses>bus"`
}

const FileExtensionFzp = ".fzp"

func IsFileFzp(src string) bool {
	if filepath.Ext(src) == FileExtensionFzp {
		return true
	} else {
		return false
	}
}

// ReadFzp and decode xml data
func ReadFzp(src string) (Fzp, error) {
	fzpData := Fzp{}
	// read
	fzpBytes, err := ioutil.ReadFile(src)
	if err != nil {
		return fzpData, err
	}
	// decode XML
	err = xml.Unmarshal(fzpBytes, &fzpData)
	if err != nil {
		return fzpData, err
	}
	return fzpData, nil
}

// func ReadJSON(src string) (Fzp, error) {
// }

func (f *Fzp) WriteJSON(src string) error {
	data, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(src, data, 0755)
	return err
}
