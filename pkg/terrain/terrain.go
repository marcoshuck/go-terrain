package terrain

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

type Terrain interface {
	Mesh() *graphic.Mesh
}

type terrain struct {
	geometry *geometry.Geometry
	material *material.Standard
	mesh *graphic.Mesh
}

func (t terrain) Mesh() *graphic.Mesh {
	return t.mesh
}

func NewTerrain() Terrain {
	var t Terrain
	geo := geometry.NewPlane(500, 500)
	mat := material.NewStandard(math32.NewColor("lightgreen"))
	t = &terrain{
		geometry: geo,
		material: mat,
		mesh:     graphic.NewMesh(geo, mat),
	}
	return t
}