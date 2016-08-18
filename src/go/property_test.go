package fzp

import (
	"testing"
)

func Test_NewProperty(t *testing.T) {
	p := NewProperty("p1", "foo")
	if p.Name != "p1" && p.Value != "foo" {
		t.Error("Property data not equal")
	}
}

func Test_Property_Check(t *testing.T) {
	// empty
	p1 := NewProperty("", "")
	err1 := p1.Check()
	if err1 == nil {
		t.Error("Property.Check missing error")
	}

	// with data
	p2 := NewProperty("foo", "bar")
	err2 := p2.Check()
	if err2 != nil {
		t.Error(err2)
	}
}

func Test_Property_CheckName(t *testing.T) {
	p := Property{}
	// empty
	err := p.CheckName()
	if err == nil {
		t.Error("Property.CheckName test failed")
	}
	// with data
	p.Name = "test"
	err = p.CheckName()
	if err != nil {
		t.Error("Property.CheckName test failed - ERROR:", err)
	}
}

func Test_Property_CheckValue(t *testing.T) {
	p := Property{}
	// empty
	err := p.CheckValue()
	if err == nil {
		t.Error("Property.CheckValue test failed")
	}
	// with data
	p.Value = "test"
	err = p.CheckValue()
	if err != nil {
		t.Error("Property.CheckValue test failed - ERROR:", err)
	}
}
