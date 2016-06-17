package fzp

// ConnectorLayer represet a fzp ConnectorLayer data object
type ConnectorLayer struct {
	Layer      string `xml:"layer,attr"      json:"layer"`
	SvgID      string `xml:"svgId,attr"      json:"svgId"`
	TerminalID string `xml:"terminalId,attr" json:"terminalId"`
}

// NewConnectorLayer returns a ConnectorLayer object
func NewConnectorLayer() ConnectorLayer {
	cLayer := ConnectorLayer{}
	return cLayer
}

// Check validate the ConnectorLayer data
func (c *ConnectorLayer) Check() error {
	// TODO: create check rules
	return nil
}
