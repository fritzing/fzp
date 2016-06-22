package fzp

// Connector represet a fzp Connector data object
type Connector struct {
	ID             string           `xml:"id,attr"                json:"id"`
	Name           string           `xml:"name,attr"              json:"name"`
	Type           string           `xml:"type,attr"              json:"type"`
	Description    string           `xml:"description"            json:"description"`
	BreadboardView []ConnectorLayer `xml:"views>breadboardView>p" json:"breadboard"`
	PcbView        []ConnectorLayer `xml:"views>pcbView>p"        json:"pcb"`
	SchematicView  []ConnectorLayer `xml:"views>schematicView>p"  json:"schematic"`
}

// NewConnector returns a Connector object
func NewConnector(id string) Connector {
	con := Connector{}
	con.ID = id
	con.BreadboardView = make([]ConnectorLayer, 0)
	con.PcbView = make([]ConnectorLayer, 0)
	con.SchematicView = make([]ConnectorLayer, 0)
	return con
}

// Check validate the Connector data
func (c *Connector) Check() error {
	// TODO: create check rules
	return nil
}
