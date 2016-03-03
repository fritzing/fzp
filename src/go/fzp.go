package fzp

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"path/filepath"
)

type FZP struct {
	XMLName         xml.Name    `xml:"module" json:"-"`
	FritzingVersion string      `xml:"fritzingVersion,attr"`
	ModuleId        string      `xml:"moduleId,attr" json:"module_id"`
	ReferenceFile   string      `xml:"referenceFile,attr"`
	Version         string      `xml:"version"`
	Title           string      `xml:"title"`
	Description     string      `xml:"description"`
	Author          string      `xml:"author"`
	Date            string      `xml:"date"`
	URL             string      `xml:"url"`
	Label           string      `xml:"label"`
	Tags            []string    `xml:"tags>tag"`
	Properties      []Property  `xml:"properties>property"`
	Views           Views       `xml:"views"`
	Connectors      []Connector `xml:"connectors>connector"`
	Buses           []Bus       `xml:"buses>bus"`
}

const FileExtension = ".fzp"

func IsFileFzp(src string) bool {
	if filepath.Ext(src) == FileExtension {
		return true
	} else {
		return false
	}
}

// ReadFzp and decode xml data
func ReadFzp(src string) (FZP, error) {
	fzpData := FZP{}
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

func (f *FZP) WriteFzp(src string) error {
	data, err := xml.Marshal(f)
	if err != nil {
		return err
	}
	filename := src
	if !IsFileFzp(src) {
		filename = src + FileExtension
	}
	err = ioutil.WriteFile(filename, data, 0777)
	if err != nil {
		return err
	}
	return nil
}

// func ReadJSON(src string) (FZP, error) {
// }

func (f *FZP) WriteJSON(src string) error {
	data, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(src+".json", data, 0777)
	if err != nil {
		return err
	}
	return nil
}
