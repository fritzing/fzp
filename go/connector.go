package fzp

type Connector struct {
	Id             string           `xml:"id,attr"`
	Name           string           `xml:"name,attr"`
	Type           string           `xml:"type,attr"`
	Description    string           `xml:"description"`
	BreadboardView []ConnectorLayer `xml:"views>breadboardView>p"`
	PcbView        []ConnectorLayer `xml:"views>pcbView>p"`
	SchematicView  []ConnectorLayer `xml:"views>schematicView>p"`
}

func (b *Connector) Check() error {
	return nil
}
