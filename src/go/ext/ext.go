package ext

import (
	"errors"
	"path/filepath"
)

const (
	// ExtFzp store the .fzp extension as constant variable
	ExtFzp = ".fzp"
	// ExtJSON store the .json extension as constant variable
	ExtJSON = ".json"
	// ExtYAML store the .yml extension as constant variable
	ExtYAML = ".yml"
)

var errorUseLowercase = errors.New("Not common file extension! Please use lowercase for the extension")

// IsValid checks if a filename extension is valid. if not, return a description about the issue
func IsValid(filename string) (bool, error) {
	fileExt := filepath.Ext(filename)
	switch fileExt {

	case ExtFzp, ExtJSON, ExtYAML:
		return true, nil

	case ".FZP":
		return false, errorUseLowercase
	case ".JSON":
		return false, errorUseLowercase
	case ".YML", ".YAML", ".yaml":
		return false, errorUseLowercase
	case ".XML", ".xml":
		return false, errors.New("Not common file extension! Use .fzp instead of " + fileExt)

	default:
		return false, errors.New("Unknown file extension! (" + fileExt + ")")
	}
}
