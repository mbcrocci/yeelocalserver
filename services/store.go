package services

import (
	"errors"
	"sync"

	"github.com/mbcrocci/yeelocalsrv/entities"
)

type LightStore struct {
	lights []*entities.Light
	mux    sync.Mutex
}

func (ls *LightStore) Len() int {
	return len(ls.lights)
}

func (ls *LightStore) Init() {
	ls.lights = make([]*entities.Light, 0)
	ls.mux = sync.Mutex{}
}

func (ls *LightStore) Add(l *entities.Light) {
	light, _ := ls.Find(l.ID)
	if light == nil {
		ls.mux.Lock()
		ls.lights = append(ls.lights, l)
		ls.mux.Unlock()
	}
}

func (ls *LightStore) Find(id string) (*entities.Light, error) {
	for _, light := range ls.lights {
		if light.ID == id {
			return light, nil
		}
	}
	return nil, errors.New("Couldn't find light")
}

func (ls *LightStore) Lights() []*entities.Light {
	return ls.lights
}
