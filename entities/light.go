package entities

import (
	"errors"
	"fmt"
	"strings"
)

type Light struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Model           string `json:"model"`
	Power           string `json:"power"`
	FirmwareVersion int    `json:"fw_ver,string"`
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

func (l *Light) Supports() []string {
	return strings.Split(l.Support, " ")
}

// func (l *Light) RGB() (uint32, uint32, uint32, uint32) {
// 	return color.Alpha{A: l.Alpha}.RGBA()
// }
