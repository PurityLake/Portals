// main.go
package main

import (
	ecs "github.com/PurityLake/go-ecs"
	// "github.com/portals/v2/mapgen"

	"github.com/portals/v2/render"
	"github.com/portals/v2/systems"
	// "github.com/portals/v2/systems"
)

func main() {
	world := ecs.World{}

	window := render.NewWindow("Portals", 800, 600)
	err := window.Init()
	if err != nil {
		panic(err)
	}

	renderer, err := window.CreateRenderer()
	if err != nil {
		panic(err)
	}
	world.AddSystem(&systems.RenderSystem{Renderer: &render.Renderer{Renderer: renderer}})
	world.AddSystem(&systems.PlayerSystem{Renderer: renderer})
	world.Start()

	running := true
	for running {
		world.Update()
	}
}
