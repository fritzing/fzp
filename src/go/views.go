package fzp

import (
	"errors"
	"strings"
)

// Views represet a fzp Views data object
type Views struct {
	Icon       ViewLayers `xml:"iconView>layers"       json:"icon"`
	Breadboard ViewLayers `xml:"breadboardView>layers" json:"breadboard"`
	Pcb        ViewLayers `xml:"pcbView>layers"        json:"pcb"`
	Schematic  ViewLayers `xml:"schematicView>layers"  json:"schematic"`
}

// NewViews returns a Views object
func NewViews() Views {
	v := Views{}
	v.Icon = NewViewLayers()
	v.Breadboard = NewViewLayers()
	v.Pcb = NewViewLayers()
	v.Schematic = NewViewLayers()
	return v
}

// Check validate the Views data. this calls the icon, breadboard, pcb and
// schematic check funcs and concatendate the error messages
func (v *Views) Check() error {
	var errConcat []string

	if err := v.Icon.Check(); err != nil {
		errConcat = append(errConcat, err.Error())
	}
	if err := v.Breadboard.Check(); err != nil {
		errConcat = append(errConcat, err.Error())
	}
	if err := v.Pcb.Check(); err != nil {
		errConcat = append(errConcat, err.Error())
	}
	if err := v.Schematic.Check(); err != nil {
		errConcat = append(errConcat, err.Error())
	}

	if len(errConcat) != 0 {
		return errors.New(strings.Join(errConcat, "\n"))
	}
	return nil
}
