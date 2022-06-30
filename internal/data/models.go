package data

type Models struct {
	Lights interface {
		Add(light *Light)
		Find(id string) (*Light, error)
		List() []*Light
	}
}

func NewModels() Models {
	return Models{
		Lights: NewLightModel(),
	}
}
