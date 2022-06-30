package data

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

type Light struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Model           string `json:"model"`
	FirmwareVersion int    `json:"fw_ver,string"`
	Location        string `json:"location"`
	Power           string `json:"power"`
	Brightness      int    `json:"bright,string"`
	ColorMode       string `json:"color_mode"`
	Ct              string `json:"ct"`
	Alpha           int    `json:"rgb,string"`
	Hue             int    `json:"hue,string"`
	Saturation      int    `json:"sat,string"`
	Support         string `json:"support"`
}

func NewLightFromString(str string) (*Light, error) {
	if len(str) == 0 {
		return nil, errors.New("empty string to parse")
	}

	return &Light{}, nil
}

func (l Light) String() string {
	return fmt.Sprintf("%s: {\n\tpower: %s,\n}", l.Name, l.Power)
}

func (l Light) Equal(l2 *Light) bool {
	return l.ID == l2.ID
}

func (l *Light) Supports(cmd string) bool {
	for _, c := range strings.Split(l.Support, " ") {
		if c == cmd {
			return true
		}
	}

	return false
}

// func (l *Light) RGB() (uint32, uint32, uint32, uint32) {
// 	return color.Alpha{A: l.Alpha}.RGBA()
// }

type LightModel struct {
	lights []*Light
	mux    sync.Mutex
}

func NewLightModel() *LightModel {
	return &LightModel{
		lights: make([]*Light, 0),
		mux:    sync.Mutex{},
	}
}

func (l *LightModel) Len() int {
	return len(l.lights)
}

func (l *LightModel) List() []*Light {
	return l.lights
}

func (l *LightModel) Add(light *Light) {
	l.mux.Lock()
	defer l.mux.Unlock()

	fl, err := l.Find(light.ID)
	if err != nil {
		l.lights = append(l.lights, light)
	} else {
		*fl = *light
	}
}

func (l *LightModel) Find(id string) (*Light, error) {
	for _, light := range l.lights {
		if light.ID == id {
			return light, nil
		}
	}
	return nil, errors.New("light not found")
}
