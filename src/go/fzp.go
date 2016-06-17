package fzp

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Fzp represet a fzp data object
type Fzp struct {
	XMLName         xml.Name `xml:"module"                json:"-"                 yaml:"-"`
	FritzingVersion string   `xml:"fritzingVersion,attr"  json:"fritzingversion"   yaml:"fritzingversion"`
	ModuleID        string   `xml:"moduleId,attr"         json:"moduleid"          yaml:"moduleid"`
	ReferenceFile   string   `xml:"referenceFile,attr"    json:"referencefileurl"  yaml:"referencefile" yaml:"referencefile"`
	Version         string   `xml:"version"               json:"version"           json:"version"`
	Title           string   `xml:"title"                 json:"title"             json:"title"`
	Description     string   `xml:"description"           json:"description"       json:"description"`
	Author          string   `xml:"author"                json:"author"            json:"author"`
	Date            string   `xml:"date"                  json:"date"              json:"date"`
	URL             string   `xml:"url"                   json:"url"               json:"url"`
	Label           string   `xml:"label"                 json:"label"             json:"label"`
	//Taxonomy        string      `xml:"taxonomy"`
	//Family          string      `xml:"family"`
	//Variant         string      `xml:"variant"`
	Tags       []string    `xml:"tags>tag"               json:"tags"`
	Properties []Property  `xml:"properties>property"    json:"version"`
	Views      Views       `xml:"views"                  json:"views"`
	Connectors []Connector `xml:"connectors>connector"   json:"connectors"`
	Buses      []Bus       `xml:"buses>bus"              json:"busses"`
}

const (
	// ExtFzp store the .fzp extension ast constant variable
	ExtFzp = ".fzp"
	// ExtJSON store the .fzp extension ast constant variable
	ExtJSON = ".json"
	// ExtYAML store the .fzp extension ast constant variable
	ExtYAML = ".yml"
)

func hasExt(ext, filename string) bool {
	if filepath.Ext(filename) == ext {
		return true
	}
	return false
}

// HasExtFzp returns true if the filename has a .fzp suffix
func HasExtFzp(filename string) bool {
	return hasExt(ExtFzp, filename)
}

// HasExtJSON returns true if the filename has a .json suffix
func HasExtJSON(filename string) bool {
	return hasExt(ExtJSON, filename)
}

// HasExtYAML returns true if the filename has a .yml suffix
func HasExtYAML(filename string) bool {
	return hasExt(ExtYAML, filename)
}

// ReadFzp and decode xml data
func ReadFzp(src string) (Fzp, []byte, error) {
	fzpData := Fzp{}
	// read
	fzpBytes, err := ioutil.ReadFile(src)
	if err != nil {
		return fzpData, []byte{}, err
	}
	// decode XML
	err = xml.Unmarshal(fzpBytes, &fzpData)
	if err != nil {
		return fzpData, []byte{}, err
	}
	return fzpData, fzpBytes, nil
}

// func ReadXML(src string) (Fzp, error) {
// }

// func ReadJSON(src string) (Fzp, error) {
// }

// func ReadYAML(src string) (Fzp, error) {
// }

// Marshal the Fzp data object
func (f *Fzp) Marshal(format string) ([]byte, error) {
	// fmt.Println("marshal", format)

	var data []byte
	var err error
	switch format {
	case "xml":
		// fmt.Println("...xml")
		data, err = f.ToXML()
		break

	case "json":
		// fmt.Println("...json")
		data, err = f.ToJSON()
		break

	case "yaml":
		// fmt.Println("...yaml")
		data, err = f.ToYAML()
		break

	default:
		errMsg := "format '" + format + "' not supported!\n"
		errMsg += "you can choose between the following three formats:"
		errMsg += "- xml (default)"
		errMsg += "- json"
		errMsg += "- yaml"
		return []byte(""), errors.New(errMsg)
	}

	if err != nil {
		return []byte(""), errors.New("Format Error:" + err.Error())
	}

	return data, err
}

// ToXML encode the fzp data to yaml and returns a bytes array
func (f *Fzp) ToXML() ([]byte, error) {
	data, err := xml.MarshalIndent(f, "", "  ")
	return data, err
}

// ToJSON encode the fzp data to json and returns a bytes array
func (f *Fzp) ToJSON() ([]byte, error) {
	data, err := json.MarshalIndent(f, "", "  ")
	return data, err
}

// ToYAML encode the fzp data to yaml and returns a bytes array
func (f *Fzp) ToYAML() ([]byte, error) {
	data, err := yaml.Marshal(f)
	return data, err
}

// WriteXML write the fzp data formatted as xml
func (f *Fzp) WriteXML(src string) error {
	data, err := f.ToXML()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(src, data, 0755)
	return err
}

// WriteJSON write the fzp data formatted as json
func (f *Fzp) WriteJSON(src string) error {
	data, err := f.ToJSON()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(src, data, 0755)
	return err
}

// WriteYAML write the fzp data formatted as yaml
func (f *Fzp) WriteYAML(src string) error {
	data, err := f.ToYAML()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(src, data, 0755)
	return err
}
