package fzp

import (
	"testing"
)

func Test_NewBus(t *testing.T) {
	b := NewBus("Test")
	if b.Id != "Test" {
		t.Error("Bus NewBus test faild")
	}
}

func Test_Bus_CheckId(t *testing.T) {
	// test without data (the empty state)
	b := Bus{}
	err := b.CheckId()
	if err == nil {
		t.Error("Bus test failed")
	}

	// test with data
	b.Id = "test"
	err = b.CheckId()
	if err != nil {
		t.Error("Bus test failed")
	}
}

func Test_Bus_CheckNodeMembers(t *testing.T) {
	b := Bus{}
	err := b.CheckNodeMembers()
	if err == nil {
		t.Error("Bus test failed")
	}

	b.AddNodeMember("node1")
	if len(b.NodeMember) != 1 && b.NodeMember[0].ConnectorId != "node1" {
		t.Error("AddNodeMember broken...")
	}
	err = b.CheckNodeMembers()
	if err != nil {
		t.Error("Bus test failed")
	}
}
