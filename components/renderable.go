package components

import (
	"reflect"

	ecs "github.com/PurityLake/go-ecs"
	"github.com/veandco/go-sdl2/sdl"
)

type Renderable struct {
	ID      int
	W, H    int
	Texture *sdl.Texture
}

func (r Renderable) Id() int {
	return r.ID
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
