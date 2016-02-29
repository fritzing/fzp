package fzp

import (
	"testing"
)

func Test_ConnectorLayer(t *testing.T) {
	cLayer := NewConnectorLayer()
	cLayer.Layer = "l1"
	cLayer.SvgId = "s1"
	cLayer.TerminalId = "t1"
	if cLayer.Layer != "l1" {
		t.Error("Connector test failed")
	}
}
