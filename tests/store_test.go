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

func TestFindsLight(t *testing.T) {
	s := &services.LightStore{}
	s.Init()
	PopulateStore(s)

	light, err := s.Find("testid1")
	if err != nil {
		t.Error("Cound't find light by id")
	}

	if light.ID != "testid1" {
		t.Error("Returned light is incorrect")
	}
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

func TestAddModifiesExistingLigt(t *testing.T) {
	s := &services.LightStore{}
	s.Init()
	PopulateStore(s)

	toModify := "testid1"
	newName := "l3"
	s.Add(&entities.Light{ID: toModify, Name: newName})

	if s.Len() > 2 {
		t.Error("Shouldn't add the light")
	}

	l, err := s.Find(toModify)
	if err != nil {
		t.Error(err)

	} else if l.Name != newName {
		t.Error("Should have modified light")
	}
}
