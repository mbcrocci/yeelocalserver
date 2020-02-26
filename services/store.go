package services

import (
	"sync"

	"github.com/mbcrocci/yeelocalsrv/entities"
)

type LightStore struct {
	lights []*entities.Light
	mux    *sync.Mutex
}

func (ls *LightStore) Init() {
	ls.lights = make([]*entities.Light, 0)
	ls.mux = &sync.Mutex{}
}

func (ls *LightStore) Add(l *entities.Light) {
	light := ls.Find(l.ID)
	if light == nil {
		ls.mux.Lock()
		ls.lights = append(ls.lights, light)
		ls.mux.Unlock()
	}
}

func (ls *LightStore) Find(id string) *entities.Light {
	for _, light := range ls.lights {
		if light.ID == id {
			return light
		}
	}
	return nil
}

func (ls *LightStore) Lights() []*entities.Light {
	return ls.lights
}
