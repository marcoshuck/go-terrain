package box

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/marcoshuck/go-terrain/pkg/models"
)

type Box interface {
	models.Model
}

type box struct {
	geometry *geometry.Geometry
	material *material.Standard
	mesh *graphic.Mesh
}

func (t *box) Mesh() []*graphic.Mesh {
	return []*graphic.Mesh{
		t.mesh,
	}
}

func NewBox(x, y, z, size float32) Box {
	geo := geometry.NewCube(size)
	mat := material.NewStandard(math32.NewColor("darkorange"))
	b := &box{
		geometry: geo,
		material: mat,
		mesh:     graphic.NewMesh(geo, mat),
	}
	b.mesh.SetPosition(x, y, z + size / 2)
	return b
}