// main.go
package main

import (
	ecs "github.com/PurityLake/go-ecs"
	// "github.com/portals/v2/mapgen"
	"github.com/portals/v2/components"
	"github.com/portals/v2/render"
	"github.com/portals/v2/systems"

	// "github.com/portals/v2/systems"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
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

	if err := ttf.Init(); err != nil {
		panic(err)
	}

	world.AddSystem(&systems.PlayerSystem{Renderer: renderer})
	world.Start()

	running := true
	query := ecs.NewQuery(ecs.Type[components.Renderable](), ecs.Type[components.Position]())
	for running {
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()
		componentsFound, found := world.Query(query)
		if found {
			for _, componentList := range componentsFound {
				renderable := componentList[0].(components.Renderable)
				position := componentList[1].(components.Position)
				renderer.Copy(renderable.Texture, nil,
					&sdl.Rect{
						X: int32(position.X),
						Y: int32(position.Y),
						W: int32(renderable.W),
						H: int32(renderable.H),
					})
			}
		}
		renderer.Present()
		world.Update()
	}
}
