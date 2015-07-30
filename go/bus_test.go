package fzp

import (
	"testing"
)

func Test_Bus(t *testing.T) {
	bus := Bus{}
	bus.Id = "test-id"

	errCheckId := bus.CheckId()
	if errCheckId != nil {
		t.Error("Bus.CheckId failed")
	}
}
