package components

import (
	"reflect"

	"github.com/PurityLake/go-ecs"
	"github.com/veandco/go-sdl2/sdl"
)

type Renderable struct {
	id      int
	Texture *sdl.Texture
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
