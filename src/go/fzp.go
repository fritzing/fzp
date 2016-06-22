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
	ReferenceFile   string   `xml:"referenceFile,attr"    json:"referencefileurl"  yaml:"referencefile"`
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
	Tags       Tags       `xml:"tags>tag"               json:"tags"`
	Properties Properties `xml:"properties>property"    json:"version"`
	Views      Views      `xml:"views"                  json:"views"`
	Connectors Connectors `xml:"connectors>connector"   json:"connectors"`
	Buses      Buses      `xml:"buses>bus"              json:"busses"`
}

// The Format type we unse for enumerate
type Format int

const (
	FormatNotSupported Format = iota
	FormatUnknown
	FormatFzp
	FormatJSON
	FormatYAML
)

func (f *Format) String() string {
	switch *f {
	case FormatNotSupported:
		return "Format not supported"
	case FormatUnknown:
		return "Format unknown"
	case FormatFzp:
		return "Format Fzp"
	case FormatJSON:
		return "Format json"
	case FormatYAML:
		return "Format yaml"
	}
	return ""
}

func GetFormat(src string) (Format, bool) {
	// is it a filename with an extension?
	tmp := filepath.Ext(src)
	isFile := true
	// fmt.Println("tmp", tmp)
	if tmp == "" {
		tmp = src
		isFile = false
	}

	switch tmp {
	case "":
		return FormatUnknown, isFile

	case "FZP", "fzp", ".FZP", ".fzp":
		return FormatFzp, isFile

	case "JSON", "json", ".JSON", ".json":
		return FormatJSON, isFile

	case "YAML", "yaml", "yml", ".YAML", ".yaml", ".yml":
		return FormatYAML, isFile

	default:
		return FormatNotSupported, isFile
	}
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
func (f *Fzp) Marshal(format Format) ([]byte, error) {
	// fmt.Println("marshal", format)

	var data []byte
	var err error
	switch format {
	case FormatFzp:
		// fmt.Println("...xml")
		data, err = f.ToXML()
		break

	case FormatJSON:
		// fmt.Println("...json")
		data, err = f.ToJSON()
		break

	case FormatYAML:
		// fmt.Println("...yaml")
		data, err = f.ToYAML()
		break

	default:
		errMsg := "format '" + string(format) + "' not supported!\n"
		errMsg += "you can choose between the following three formats:\n"
		errMsg += "- fzp (default)\n"
		errMsg += "- json\n"
		errMsg += "- yml or yaml\n"
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
	data = append([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n"), data...)
	data = append(data, []byte("\n")...)
	return data, err
}

// ToJSON encode the fzp data to json and returns a bytes array
func (f *Fzp) ToJSON() ([]byte, error) {
	data, err := json.MarshalIndent(f, "", "  ")
	data = append(data, []byte("\n")...)
	return data, err
}

// ToYAML encode the fzp data to yaml and returns a bytes array
func (f *Fzp) ToYAML() ([]byte, error) {
	data, err := yaml.Marshal(f)
	return data, err
}

// WriteXML write the fzp data formatted as xml
// func (f *Fzp) WriteXML(src string) error {
// 	data, err := f.ToXML()
// 	if err != nil {
// 		return err
// 	}
// 	err = ioutil.WriteFile(src, data, 0755)
// 	return err
// }
//
// // WriteJSON write the fzp data formatted as json
// func (f *Fzp) WriteJSON(src string) error {
// 	data, err := f.ToJSON()
// 	if err != nil {
// 		return err
// 	}
// 	err = ioutil.WriteFile(src, data, 0755)
// 	return err
// }
//
// // WriteYAML write the fzp data formatted as yaml
// func (f *Fzp) WriteYAML(src string) error {
// 	data, err := f.ToYAML()
// 	if err != nil {
// 		return err
// 	}
// 	err = ioutil.WriteFile(src, data, 0755)
// 	return err
// }
