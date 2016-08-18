package fzp

import (
	"errors"
)

// Property represet a fzp Property data object
type Property struct {
	Name  string `xml:"name,attr" json:"name"`
	Value string `xml:",chardata" json:"value"`
}

// NewProperty returns a Property object
func NewProperty(n, v string) Property {
	p := Property{}
	p.Name = n
	p.Value = v
	return p
}

// Check validate the Property data
func (p *Property) Check() error {
	errMsg := ""
	if err := p.CheckName(); err != nil {
		errMsg = err.Error()
	}
	if err := p.CheckValue(); err != nil {
		errMsg += ", " + err.Error()
		return errors.New(errMsg)
	}
	return nil
}

// CheckName validate the Property Name data
func (p *Property) CheckName() error {
	if p.Name == "" {
		return errors.New("property name undefined")
	}
	return nil
}

// CheckValue validate the Property Value data
func (p *Property) CheckValue() error {
	if p.Value == "" {
		return errors.New("property value undefined")
	}
	return nil
}
