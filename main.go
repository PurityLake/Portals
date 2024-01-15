// main.go
package main

import (
	"fmt"
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

func (r Renderable) Type() reflect.Type {
	return reflect.TypeOf(r)
}

func main() {
	system := ecs.System{}
	position := ecs.Component(Position{0, 0})
	position2 := ecs.Component(Position{0, 0})
	renderable := ecs.Component(Renderable{nil})

	entity := ecs.NewEntity("player", position, renderable)
	entity2 := ecs.NewEntity("test", position2)
	system.AddEntity(entity)
	system.AddEntity(entity2)
	entities, components := system.QueryWithEntity(Position{}.Type(), Renderable{}.Type())

	for i, entity := range entities {
		fmt.Println(entity.Name())
		componentList := components[i]
		for _, component := range componentList {
			if ecs.ComponentIsA[Position](component) {
				fmt.Println("Position")
			}
			if ecs.ComponentIsA[Renderable](component) {
				fmt.Println("Renderable")
			}
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
