package ext

import (
	"errors"
	"path/filepath"
)

const (
	// ExtFzp store the .fzp extension ast constant variable
	ExtFzp = ".fzp"
	// ExtJSON store the .fzp extension ast constant variable
	ExtJSON = ".json"
	// ExtYAML store the .fzp extension ast constant variable
	ExtYAML = ".yml"
)

func FilenameEqualToExt(filename, ext string) bool {
	if filepath.Ext(filename) == ext {
		return true
	}
	return false
}

// HasExtFzp returns true if the filename has a .fzp suffix
func HasExtFzp(filename string) bool {
	return FilenameEqualToExt(filename, ExtFzp)
}

// HasExtJSON returns true if the filename has a .json suffix
func HasExtJSON(filename string) bool {
	return FilenameEqualToExt(filename, ExtJSON)
}

// HasExtYAML returns true if the filename has a .yml suffix
func HasExtYAML(filename string) bool {
	return FilenameEqualToExt(filename, ExtYAML)
}

func IsFormatSupported(filename string) error {
	if !HasExtFzp(filename) && !HasExtJSON(filename) && !HasExtYAML(filename) {
		return errors.New("format not supported")
	}
	return nil
}
