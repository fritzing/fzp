package fzp

import (
	"testing"
)

func Test_ViewLayer(t *testing.T) {
	layer := NewViewLayer("TEST-ID")
	if layer.LayerID != "TEST-ID" {
		t.Error("LayerID not equal")
	}

	// err := layer.Check()
	// if err != nil {
	// 	t.Error(err)
	// }
}
