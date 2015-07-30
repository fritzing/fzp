package fzp

type ConnectorLayer struct {
	Layer      string `xml:"layer,attr"`
	SvgId      string `xml:"svgId,attr"`
	TerminalId string `xml:"terminalId,attr"`
}
