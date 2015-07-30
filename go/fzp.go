package fzp

import (
	"encoding/xml"
	"errors"
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

// ReadFzp and decode xml data
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

func (f *Fzp) CheckTags() error {
	if len(f.Tags) != 0 {
		for _, tag := range f.Tags {
			if tag == "" {
				return errors.New("tag value undefined!")
			}
		}
	}
	return nil
}

func (f *Fzp) CheckProperties() error {
	if len(f.Properties) != 0 {
		for _, property := range f.Properties {
			if err := property.CheckName(); err != nil {
				return err
			}
			if err := property.CheckValue(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (f *Fzp) CheckViews() error {
	// TODO: ...
	return nil
}

func (f *Fzp) CheckConnectors() error {
	if len(f.Connectors) != 0 {
		for _, connector := range f.Connectors {
			if err := connector.Check(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (f *Fzp) CheckBuses() error {
	if len(f.Buses) != 0 {
		for _, bus := range f.Buses {
			if err := bus.Check(); err != nil {
				return err
			}
		}
	}
	return nil
}

// check all the data
func (f *Fzp) Check() []error {
	var errList []error

	if err := f.CheckTags(); err != nil {
		errList = append(errList, err)
	}
	if err := f.CheckProperties(); err != nil {
		errList = append(errList, err)
	}
	if err := f.CheckViews(); err != nil {
		errList = append(errList, err)
	}
	if err := f.CheckConnectors(); err != nil {
		errList = append(errList, err)
	}
	if err := f.CheckBuses(); err != nil {
		errList = append(errList, err)
	}

	return errList
}
