package tests

import (
	"testing"

	"github.com/mbcrocci/yeelocalsrv/internal/data"
	"github.com/mbcrocci/yeelocalsrv/internal/yeelight"
)

type ParseResult struct {
	light *data.Light
	err   error
}

func TestEmptyParse(t *testing.T) {
	_, err := yeelight.NewDiscoverService().ParseLight("")
	if err == nil {
		t.Error("Parse of an empty string should return an error")
	}
}
