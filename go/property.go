package fzp

import (
	"errors"
)

type Property struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

func (p *Property) CheckName() error {
	if p.Name == "" {
		return errors.New("Missing Property.Name")
	} else {
		return nil
	}
}

func (p *Property) CheckValue() error {
	if p.Value == "" {
		return errors.New("Missing Property.Value")
	} else {
		return nil
	}
}
