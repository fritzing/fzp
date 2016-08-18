package fzp

// Connectors array
type Connectors []Connector

// NewConnectors return a new Connectors object
func NewConnectors() Connectors {
	c := Connectors{}
	c = make([]Connector, 0)
	return c
}
