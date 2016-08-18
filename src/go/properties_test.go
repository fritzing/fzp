package fzp

import (
	"testing"
)

func Test_Properties(t *testing.T) {
	props := NewProperties("test-fam")
	if props.Total() != 1 {
		t.Error("NewProperties returned object missing family property")
	}

	tmp, err := props.GetValue("family")
	if err != nil {
		t.Error(err)
	}
	if tmp != "test-fam" {
		t.Error("Properties family not equal")
	}

	err = props.AddValue("foo", "F")
	if err != nil {
		t.Error(err)
	}
	err = props.AddValue("bar", "B")
	if err != nil {
		t.Error(err)
	}

	if props.Total() != 3 {
		t.Error("Properties Total not equal")
	}

	// test if an error was returned if the prop exist...
	err = props.AddValue("foo", "F")
	if err == nil {
		t.Error("missing error")
	}

}
