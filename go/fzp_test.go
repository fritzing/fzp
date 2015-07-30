package fzp

import (
	"testing"
)

func Test_ReadFzp_Ok(t *testing.T) {
	f, err := ReadFzp("../sample.fzp")
	if err != nil {
		t.Error("Fzp.ReadFzp failed")
	}

	errCheck := f.Check()
	if errCheck != nil {
		t.Error("Fzp.Check broken:", errCheck)
	}
}

func Test_ReadFzp_Failed(t *testing.T) {
	_, err := ReadFzp("../not.found")
	if err == nil {
		t.Error("fzp ReadFzp failed")
	}
}
