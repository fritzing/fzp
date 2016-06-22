package fzp

import (
	"testing"
)

func Test_Properties(t *testing.T) {
	props := NewProperties()
	if len(*props) != 0 {
		t.Error("NewProperties returned object not equal")
	}
	err := props.AddValue("foo", "F")
	if err != nil {
		t.Error(err)
	}
	err = props.AddValue("bar", "B")
	if err != nil {
		t.Error(err)
	}
}
