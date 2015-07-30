package fzp

import (
	"testing"
)

func Test_ReadFzp_Ok(t *testing.T) {
	f, err := ReadFzp("../sample.fzp")
	if err != nil {
		t.Error("fzp ReadFzp failed")
	}

	errCheck := CheckData(f)
	if errCheck != nil {
		t.Error("fzp CheckData broken")
	}

	if f.Version != "x.y.z" {
		t.Error("fzp Version not equal")
	}
	if f.Title != "part-name" {
		t.Error("fzp Title not equal")
	}
}

func Test_ReadFzp_Failed(t *testing.T) {
	_, err := ReadFzp("../not.found")
	if err == nil {
		t.Error("fzp ReadFzp failed")
	}
}
