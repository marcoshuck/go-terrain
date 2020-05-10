package tree

import (
	"github.com/g3n/engine/math32"
	"math/rand"
)

func GenerateTrees(width, height float32, amount uint, distance float32) Trees {
	var generated Trees
	for i := 1; uint(i) <= amount; i++ {
		generateTree(&generated, width, height, distance)
	}
	return generated
}

func generateTree(generated *Trees, width, height float32, distance float32) {
	var x float32
	var y float32
	for _, t := range *generated {
		var valid bool
		for !valid {
			x = (rand.Float32() - 0.5) * width
			y = (rand.Float32() - 0.5) * height
			v := math32.NewVector3(x, y, 0)
			if t.DistanceTo(v) >= distance {
				valid = true
			}
		}

	}
	created := NewTree(x, y, 0)
	*generated = append(*generated, created)
}