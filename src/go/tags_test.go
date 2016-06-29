package fzp

import (
	"testing"
)

func Test_Tags(t *testing.T) {
	tags := NewTags()
	if tags.Total() != 0 {
		t.Error("Tags Total not equal")
	}

	tags.Add("helllo")
	tags.Add("world")
	if tags.Total() != 2 {
		t.Error("Tags Total not equal")
	}

	err := tags.Add("world")
	if err == nil {
		t.Error("Tags Add 'world' missing error (world already exist)")
	}
}

func Test_Tags_Check(t *testing.T) {
	tags := Tags{}
	tags = []string{"hello", "world", "foo", "hello"}
	err := tags.Check()
	// t.Log(err)
	if err == nil {
		t.Error("missing error (hello duplicated)", err)
	}

}
