// main.go
package main

import (
	"reflect"

	"github.com/portals/v2/ecs"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Position struct {
	x, y int
}

func (p Position) Name() string {
	return "position"
}

func (p Position) Update() {}

func (p Position) Data() ecs.ComponentData {
	return p
}

func (p Position) IsA(value interface{}) bool {
	return ecs.CheckComponent[Position](value)
}

func (p Position) Type() reflect.Type {
	return reflect.TypeOf(p)
}

type Renderable struct {
	texture *sdl.Texture
}

func (r Renderable) Name() string {
	return "renderable"
}

func (r Renderable) Update() {}

func (r Renderable) Data() ecs.ComponentData {
	return r
}

func (r Renderable) IsA(value interface{}) bool {
	return ecs.CheckComponent[Renderable](value)
}

func (r Renderable) Type() reflect.Type {
	return reflect.TypeOf(r)
}

func main() {
	position := ecs.Component(Position{0, 0})
	renderable := ecs.Component(Renderable{nil})

	entity := ecs.NewEntity("player", &position, &renderable)

	if entity.HasComponents(Position{}.Type(), Renderable{}.Type()) {
		println("Has Position and Renderable")
	}

	for _, component := range entity.Components() {
		comp := *component
		switch comp.(type) {
		case Position:
			println("Position")
		case Renderable:
			println("Renderable")
		}
	}

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
