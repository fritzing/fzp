package fzp

import (
	"testing"
)

func TestFzp(t *testing.T) {
	f, err := ReadFzp("../sample.fzp")
	if err != nil {
		t.Error("fzp ReadFzp failed")
	}

	if f.Version != "x.y.z" {
		t.Error("fzp Version not equal")
	}
	if f.Title != "part-name" {
		t.Error("fzp Title not equal")
	}
}
