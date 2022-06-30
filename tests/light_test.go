package tests

import (
	"testing"

	"github.com/mbcrocci/yeelocalsrv/internal/data"
)

func TestLightEquality(t *testing.T) {
	l1 := &data.Light{ID: "testid"}
	l2 := &data.Light{ID: "testid"}

	if !l1.Equal(l2) {
		t.Error("Lights should be equal")
	}
}

func TestLightSupportsMethod(t *testing.T) {
	l := &data.Light{Support: "set_power toggle"}

	if l.Supports("set_rgb") {
		t.Error("Light shouldn't support command")
	}

	if !l.Supports("toggle") {
		t.Error("Light should support command")
	}
}
