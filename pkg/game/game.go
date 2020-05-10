package game

import (
	"fmt"
	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/window"
	"github.com/marcoshuck/go-terrain/pkg/light"
	"github.com/marcoshuck/go-terrain/pkg/models"
	"github.com/marcoshuck/go-terrain/pkg/models/box"
	"github.com/marcoshuck/go-terrain/pkg/models/tree"
	"github.com/marcoshuck/go-terrain/pkg/terrain"
	"time"
)

type Game interface {
	Run() Game
	Create() Game
	render(rend *renderer.Renderer, deltaTime time.Duration)
	onResize(eventName string, event interface{})
}

type game struct {
	application *app.Application
	scene *core.Node
	cam *camera.Camera
	terrain terrain.Terrain
	ambientLight light.AmbientLight
	models []models.Model
}

func (g *game) onResize(eventName string, event interface{}) {
	width, height := g.application.GetSize()
	g.application.Gls().Viewport(0, 0, int32(width), int32(height))
	g.cam.SetAspect(float32(width) / float32(height))
}

func (g *game) render(rend *renderer.Renderer, deltaTime time.Duration) {
	g.application.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)
	if err := rend.Render(g.scene, g.cam); err != nil {
		fmt.Errorf("[Error] %v\n", err)
	}
}

func (g *game) Run() Game {
	g.application.Run(g.render)
	return g
}

func (g *game) Create() Game {
	g.cam = camera.New(1)
	g.cam.SetPosition(100, 100, 100)

	g.scene = core.NewNode()

	gui.Manager().Set(g.scene)

	g.scene.Add(g.cam)

	camera.NewOrbitControl(g.cam)

	g.application.Subscribe(window.OnWindowSize, g.onResize)
	g.onResize("", nil)

	g.setupTerrain()
	g.setupLights()

	g.setupModels()
	g.addModels()

	g.application.Gls().ClearColor(0.529, 0.808, 0.922, 1.0)
	return g
}

func (g *game) setupTerrain() {
	g.terrain = terrain.NewTerrain()
	g.scene.Add(g.terrain.Mesh())
}

func (g *game) setupLights() {
	g.ambientLight = light.NewAmbientLight()
	g.scene.Add(g.ambientLight.Node())
}

func (g *game) setupModels() {
	b := box.NewBox(15, 15, 0, 1.5)
	g.models = append(g.models, b)

	trees := tree.GenerateTrees(500, 500, 500, 25)
	for _, tr := range trees {
		g.models = append(g.models, tr)
	}
}

func (g *game) addModels() {
	for _, m := range g.models {
		for _, mesh := range m.Mesh() {
			g.scene.Add(mesh)
		}
	}
}

func NewGame() Game {
	g := &game{
		application: app.App(),
		models: models.Models{},
	}
	return g
}