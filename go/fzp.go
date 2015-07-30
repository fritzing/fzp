package fzp

import (
	"encoding/xml"
	"io/ioutil"
)

type Fzp struct {
	XMLName         xml.Name    `xml:"module"`
	FritzingVersion string      `xml:"fritzingVersion,attr"`
	ModuleId        string      `xml:"moduleId,attr"`
	ReferenceFile   string      `xml:"referenceFile,attr"`
	Version         string      `xml:"version"`
	Title           string      `xml:"title"`
	Description     string      `xml:"description"`
	Author          string      `xml:"author"`
	Date            string      `xml:"date"`
	URL             string      `xml:"url"`
	Label           string      `xml:"label"`
	Taxonomy        string      `xml:"taxonomy"`
	Family          string      `xml:"family"`
	Variant         string      `xml:"variant"`
	Tags            []string    `xml:"tags>tag"`
	Properties      []Property  `xml:"properties>property"`
	Views           Views       `xml:"views"`
	Connectors      []Connector `xml:"connectors>connector"`
	Buses           []Bus       `xml:"buses>bus"`
}

func ReadFzp(src string) (Fzp, error) {
	fzpData := Fzp{}
	// read
	fzpBytes, errRead := ioutil.ReadFile(src)
	if errRead != nil {
		return fzpData, errRead
	}
	// decode XML
	errDecode := xml.Unmarshal(fzpBytes, &fzpData)
	if errDecode != nil {
		return fzpData, errDecode
	}
	return fzpData, nil
}
