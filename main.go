// main.go
package main

import (
	ecs "github.com/PurityLake/go-ecs"
	// "github.com/portals/v2/mapgen"
	"github.com/portals/v2/components"
	"github.com/portals/v2/render"
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

	font, err := ttf.OpenFont("assets/fonts/square.ttf", 20)
	if err != nil {
		panic(err)
	}

	surface, err := font.RenderUTF8Blended("@", sdl.Color{R: 255, G: 255, B: 255, A: 255})
	if err != nil {
		panic(err)
	}

	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		panic(err)
	}

	world.AddEntity("player", components.Renderable{Texture: texture, W: 20, H: 20}, components.Position{X: 100, Y: 100})

	running := true
	dirty := true
	for running {
		if dirty {
			renderer.SetDrawColor(0, 0, 0, 255)
			renderer.Clear()
			componentsFound, found := world.Query(components.Renderable{}.Type(), components.Position{}.Type())
			if found {
				for _, componentList := range componentsFound {
					position := componentList[1].(components.Position)
					renderable := componentList[0].(components.Renderable)
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
			dirty = false
		}
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event := event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
			case *sdl.KeyboardEvent:
				if event.Keysym.Sym == sdl.K_q {
					println("Quit")
					running = false
				}
			}
		}
	}
}
