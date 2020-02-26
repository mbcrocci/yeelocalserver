package tests

import (
	"testing"

	"github.com/mbcrocci/yeelocalsrv/entities"
	"github.com/mbcrocci/yeelocalsrv/services"
)

type ParseResult struct {
	light *entities.Light
	err   error
}

func TestInit(t *testing.T) {
	ds := services.NewDiscoverService()

	err := ds.Init()
	if err != nil {
		t.Error("Couldn't Initilize ")
	}

	ds.Shutdown()
}

func TestShutdown(t *testing.T) {
	ds := services.NewDiscoverService()
	err := ds.Init()
	if err != nil {
		return
	}

	err = ds.Shutdown()
	if err != nil {
		t.Error("Couln't shutdown Discover Service")
	}
}

func TestEmptyParse(t *testing.T) {
	_, err := services.NewDiscoverService().ParseLight("")
	if err == nil {
		t.Error("Parse of an empty string should return an error")
	}
}
