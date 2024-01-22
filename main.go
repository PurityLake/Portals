// main.go
package main

import (
	"fmt"
	"log"

	ecs "github.com/PurityLake/go-ecs"
	// "github.com/portals/v2/mapgen"
	"github.com/portals/v2/components"
	"github.com/portals/v2/render"

	// "github.com/portals/v2/systems"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var (
	UpPressed    bool = false
	UpReleased   bool = true
	DownPressed  bool = false
	DownReleased bool = true
	LeftPressed  bool = false
	RightPressed bool = false
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

	world.AddEntity("player",
		components.Renderable{ID: 1, Texture: texture, W: 20, H: 20},
		components.Position{ID: 2, X: 100, Y: 100})

	running := true
	dirty := true
	query := ecs.NewQuery(ecs.Type[components.Renderable](), ecs.Type[components.Position]())
	playerQuery := ecs.NewQuery(ecs.Type[components.Position]())
	for running {
		if dirty {
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
			dirty = false
		}
		if UpPressed || DownPressed {
			comps, found := world.QueryMut(playerQuery)
			if found {
				for _, componentList := range comps {
					for _, comp := range componentList {
						position, ok := (comp).(components.Position)
						if !ok {
							log.Fatal("Could not cast component to Position")
						}
						if UpPressed {
							fmt.Println("Move up")
							UpPressed = false
							position.Y -= 20
							dirty = true
						}
						if DownPressed {
							fmt.Println("Move down")
							DownPressed = false
							position.Y += 20
							dirty = true
						}
					}
				}
			}
		}
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event := event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
			case *sdl.KeyboardEvent:
				switch event.Keysym.Sym {
				case sdl.K_q:
					println("Quit")
					running = false
				case sdl.K_w:
					if event.State == sdl.PRESSED && !UpPressed && UpReleased {
						fmt.Println("W pressed")
						UpPressed = true
						UpReleased = false
					} else if event.State == sdl.RELEASED {
						fmt.Println("W released")
						UpPressed = false
						UpReleased = true
					}
				case sdl.K_s:
					if event.State == sdl.PRESSED && !DownPressed && DownReleased {
						fmt.Println("S pressed")
						DownPressed = true
						DownReleased = false
					} else if event.State == sdl.RELEASED {
						fmt.Println("S release")
						DownPressed = false
						DownReleased = true
					}
				}
			}
		}
	}
}
