package tests

import (
	 "github.com/mbcrocci/yeelocalsrv/entities"
	"testing"
)

func TestLightEquality(t *testing.T) {
	l1 := &entities.Light{ID: "testid"}
	l2 := &entities.Light{ID: "testid"}

	if !l1.Equal(l2) {
		t.Error("Lights should be equal")
	}
}