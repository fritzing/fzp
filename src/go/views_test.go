package fzp

import (
	"testing"
)

func Test_NewViews(t *testing.T) {
	testViews := NewViews()
	t.Log(testViews)
	// err := testViews.Check()
	// if err != nil {
	// 	t.Error(err)
	// }
}
