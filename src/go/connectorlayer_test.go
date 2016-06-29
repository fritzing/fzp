package fzp

import (
	"testing"
)

func Test_ConnectorLayer(t *testing.T) {
	cLayer := NewConnectorLayer()
	cLayer.Layer = "l1"
	cLayer.SvgID = "s1"
	cLayer.TerminalID = "t1"
	if cLayer.Layer != "l1" {
		t.Error("Connector test failed")
	}
}
