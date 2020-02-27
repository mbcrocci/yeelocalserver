package tests

import (
	"testing"

	"github.com/mbcrocci/yeelocalsrv/entities"
	"github.com/mbcrocci/yeelocalsrv/services"
)

func PopulateStore(s *services.LightStore) {
	s.Add(&entities.Light{ID: "testid1", Name: "l1"})
	s.Add(&entities.Light{ID: "testid2", Name: "l2"})
}

func TestAddsLight(t *testing.T) {
	s := &services.LightStore{}
	s.Init()

	if s.Len() != 0 {
		t.Error("Store should be empty")
	}

	PopulateStore(s)

	if s.Len() != 2 {
		t.Error("Store should have 2 lights")
	}
}

func TestFindsLight(t *testing.T) {
	s := &services.LightStore{}
	s.Init()
	PopulateStore(s)

	light := s.Find("testid1")
	if light == nil {
		t.Error("Cound't find light by id")
	}
}
