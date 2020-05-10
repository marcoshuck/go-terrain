package light

import (
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/math32"
)

type AmbientLight interface {
	Node() *light.Ambient
}

type ambientLight struct {
	light *light.Ambient
}

func (a ambientLight) Node() *light.Ambient {
	return a.light
}

func NewAmbientLight() AmbientLight {
	var l AmbientLight
	l = &ambientLight{
		light: light.NewAmbient(&math32.Color{R: 1.0, G: 1.0, B: 1.0}, 0.8),
	}

	return l
}