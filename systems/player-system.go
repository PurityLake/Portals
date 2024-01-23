package systems

import (
	"log"

	ecs "github.com/PurityLake/go-ecs"
	"github.com/portals/v2/components"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var (
	UpPressed     bool = false
	UpReleased    bool = true
	DownPressed   bool = false
	DownReleased  bool = true
	LeftPressed   bool = false
	LeftReleased  bool = true
	RightPressed  bool = false
	RightReleased bool = true
)

type PlayerSystem struct {
	id       int
	query    ecs.Query
	Renderer *sdl.Renderer
}

func (ps *PlayerSystem) Id() int {
	return ps.id
}

func (ps *PlayerSystem) Setup(world *ecs.World) {
	ps.query = ecs.NewQuery(
		ecs.Type[components.Player](),
		ecs.Type[components.Position]())

	font, err := ttf.OpenFont("assets/fonts/square.ttf", 20)
	if err != nil {
		panic(err)
	}

	surface, err := font.RenderUTF8Blended("@", sdl.Color{R: 255, G: 255, B: 255, A: 255})
	if err != nil {
		panic(err)
	}

	texture, err := ps.Renderer.CreateTextureFromSurface(surface)
	if err != nil {
		panic(err)
	}

	world.AddEntity("player",
		components.Position{X: 100, Y: 100},
		components.Renderable{Texture: texture, W: 20, H: 20},
		components.Player{})
}

func (ps *PlayerSystem) Update(world *ecs.World) {
	if UpPressed || DownPressed || LeftPressed || RightPressed {
		entities, comps, found := world.QueryWithEntityMut(ps.query)
		if found {
			for i, entity := range entities {
				for _, comp := range comps[i] {
					c, err := entity.GetComponent(comp)
					if err != nil {
						log.Fatal(err)
					}
					switch c := (*c).(type) {
					case components.Position:
						if UpPressed {
							UpPressed = false
							c.Y -= 20
						}
						if DownPressed {
							DownPressed = false
							c.Y += 20
						}
						if LeftPressed {
							LeftPressed = false
							c.X -= 20
						}
						if RightPressed {
							RightPressed = false
							c.X += 20
						}
						entity.SetComponent(comp, c)
					}
				}
			}
		}
	}
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event := event.(type) {
		case *sdl.QuitEvent:
			log.Fatal("Ok")
		case *sdl.KeyboardEvent:
			switch event.Keysym.Sym {
			case sdl.K_w:
				if event.State == sdl.PRESSED && !UpPressed && UpReleased {
					UpPressed = true
					UpReleased = false
				} else if event.State == sdl.RELEASED {
					UpPressed = false
					UpReleased = true
				}
			case sdl.K_s:
				if event.State == sdl.PRESSED && !DownPressed && DownReleased {
					DownPressed = true
					DownReleased = false
				} else if event.State == sdl.RELEASED {
					DownPressed = false
					DownReleased = true
				}
			case sdl.K_a:
				if event.State == sdl.PRESSED && !LeftPressed {
					LeftPressed = true
					LeftReleased = false
				} else if event.State == sdl.RELEASED {
					LeftPressed = false
					LeftReleased = true
				}
			case sdl.K_d:
				if event.State == sdl.PRESSED && !RightPressed {
					RightPressed = true
					RightReleased = false
				} else if event.State == sdl.RELEASED {
					RightPressed = false
					RightReleased = true
				}
			}
		}
	}
}
