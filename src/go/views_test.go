package fzp

import (
	"testing"
)

func Test_NewViews(t *testing.T) {
	testViews := NewViews()
	err := testViews.Check()
	if err != nil {
		t.Error(err)
	}
}
