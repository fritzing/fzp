package fzp

import (
	"testing"
)

func TestFzp(t *testing.T) {
	f := Fzp{}
	err := f.ReadFile("../sample.fzp")
	if err != nil {
		t.Error("fzp readfile failed")
	}
}
