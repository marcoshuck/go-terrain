package models

import "github.com/g3n/engine/graphic"

type Model interface {
	Mesh() []*graphic.Mesh
}

type Models []Model