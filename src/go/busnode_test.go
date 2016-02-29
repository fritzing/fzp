package fzp

import (
	"testing"
)

func Test_BusNode(t *testing.T) {
	bus1 := NewBusNode("Test-BusNode")

	err := bus1.Check()
	if err != nil {
		t.Error("BusNode.Check test failed")
	}

	bus1.ConnectorId = ""
	err = bus1.Check()
	if err == nil {
		t.Error("BusNode.Check test failed")
	}
}
