package fzp

type Connectors []Connector

func NewConnectors() *Connectors {
	c := Connectors{}
	c = make([]Connector, 0)
	return &c
}
