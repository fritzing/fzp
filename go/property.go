package fzp

import "errors"

type Property struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

func NewProperty(n, v string) Property {
	p := Property{}
	p.Name = n
	p.Value = v
	return p
}

func (p *Property) CheckName() error {
	if p.Name == "" {
		return errors.New("property name Undefined")
	} else {
		return nil
	}
}

func (p *Property) CheckValue() error {
	if p.Value == "" {
		return errors.New("property value undefined")
	} else {
		return nil
	}
}
