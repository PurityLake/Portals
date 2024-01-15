package systems

import (
	"fmt"
	"reflect"

	ecs "github.com/PurityLake/go-ecs"
	"github.com/veandco/go-sdl2/sdl"
)

type Position struct {
	id   int
	x, y int
}

func (p Position) Id() int {
	return p.id
}

func (p Position) Name() string {
	return "position"
}

func (p Position) Update() {}

func (p Position) Data() ecs.Data {
	return p
}

func (p Position) Type() reflect.Type {
	return reflect.TypeOf(p)
}

type Renderable struct {
	id      int
	texture *sdl.Texture
}

func (r Renderable) Id() int {
	return r.id
}

func (r Renderable) Name() string {
	return "renderable"
}

func (r Renderable) Update() {}

func (r Renderable) Data() ecs.Data {
	return r
}

func (r Renderable) Type() reflect.Type {
	return reflect.TypeOf(r)
}

type ExampleSystem struct{}

func (s ExampleSystem) Setup(world *ecs.World) {
	fmt.Println("ExampleSystem setup")
	world.AddEntity("example", Renderable{})
}

func (s ExampleSystem) Update(world *ecs.World) {
}
