package fzp

import (
	"errors"
)

type Properties []Property

func NewProperties(fname string) Properties {
	p := Properties{}
	p = make([]Property, 0)
	p = append(p, NewProperty("family", fname))
	return p
}

func (p *Properties) Total() int {
	return len(*p)
}

func (p *Properties) AddValue(name, val string) error {
	_, err := p.GetValue(name)
	if err != nil {
		newProp := NewProperty(name, val)
		*p = append(*p, newProp)
		return nil
	}
	return errors.New("exist...")
}

func (p *Properties) GetValue(name string) (string, error) {
	for _, v := range *p {
		if v.Name == name {
			return v.Value, nil
		}
	}
	return "", errors.New("property '" + name + "' does not exist")
}

func (p *Properties) Exist(name string) error {
	// TODO:...
	return nil
}

func (p *Properties) Check() error {
	// check if each property name only exist once a time
	var tmp map[string]bool
	for _, prop := range *p {
		//TODO: check if a property with name="family" exists and is not empty
		if !tmp[prop.Name] {
			tmp[prop.Name] = true
		} else {
			return errors.New("property name '" + prop.Name + "' already exist")
		}
	}
	return nil
}
