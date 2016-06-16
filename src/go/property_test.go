package fzp

import (
	"testing"
)

func Test_NewProperty(t *testing.T) {
	p := NewProperty("p1", "foo")
	if p.Name != "t1" && p.Value != "foo" {
		t.Error("Property test failed")
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
