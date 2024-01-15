// main.go
package main

import (
	ecs "github.com/PurityLake/go-ecs"
	"github.com/portals/v2/systems"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	world := ecs.World{}
	world.AddSystem(systems.ExampleSystem{})
	world.Start()

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Portals", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	if err := ttf.Init(); err != nil {
		panic(err)
	}
	defer ttf.Quit()

	font, err := ttf.OpenFont("assets/fonts/square.ttf", 20)
	if err != nil {
		panic(err)
	}
	defer font.Close()

	running := true
	dirty := true
	for running {
		if dirty {
			renderer.SetDrawColor(0, 0, 0, 255)
			renderer.Clear()
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
