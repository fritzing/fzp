package fzp

import (
	"testing"
)

func Test_Connector(t *testing.T) {
	con1 := NewConnector("Test-Connector")
	err := con1.Check()
	if err != nil && con1.ID == "Test-Connector" {
		t.Error("Connector test failed")
	}
}
