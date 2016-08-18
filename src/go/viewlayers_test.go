package fzp

import (
	"testing"
)

func Test_NewViewLayers(t *testing.T) {
	layers := NewViewLayers()
	t.Log(layers)
	// err := layers.Check()
	// if err != nil {
	// 	t.Error(err)
	// }
}
