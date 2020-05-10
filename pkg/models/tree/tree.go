package tree

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/marcoshuck/go-terrain/pkg/models"
)

type Tree interface {
	models.Model
	GetPosition() *math32.Vector3
	DistanceTo(v *math32.Vector3) float32
}

type Trees []Tree

type leaves struct {
	geometry *geometry.Geometry
	material *material.Standard
	mesh *graphic.Mesh
	size float32
}

func newLeaves(x, y, z, size float32, trunkHeight float32) *leaves {
	geo := geometry.NewCube(size)
	mat := material.NewStandard(math32.NewColor("darkgreen"))
	l := &leaves{
		geometry: geo,
		material: mat,
		mesh:     graphic.NewMesh(geo, mat),
		size: 	  size,
	}
	l.mesh.SetPosition(x, y, z + (size / 2) + trunkHeight)
	return l
}

type trunk struct {
	geometry *geometry.Geometry
	material *material.Standard
	mesh *graphic.Mesh
	height float32
}

func newTrunk(x, y, z, height float32) *trunk {
	geo := geometry.NewBox(1, 1, height)
	mat := material.NewStandard(math32.NewColor("brown"))
	t := &trunk{
		geometry: geo,
		material: mat,
		mesh:     graphic.NewMesh(geo, mat),
		height:   height,
	}
	t.mesh.SetPosition(x, y, z + height / 2)
	return t
}

type tree struct {
	leaves *leaves
	trunk *trunk
	position *math32.Vector3
}

func (t *tree) DistanceTo(v *math32.Vector3) float32 {
	return t.position.DistanceTo(v)
}

func (t *tree) GetPosition() *math32.Vector3 {
	return t.position
}

func (t *tree) Mesh() []*graphic.Mesh {
	return []*graphic.Mesh{
		t.trunk.mesh,
		t.leaves.mesh,
	}
}

func NewTree(x, y, z float32) Tree {
	return &tree{
		leaves:   newLeaves(x, y, z, 3, 3),
		trunk:    newTrunk(x, y, z, 3),
		position: math32.NewVector3(x, y, z),
	}
}