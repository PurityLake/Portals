package components

import (
	"reflect"

	ecs "github.com/PurityLake/go-ecs"
)

type Player struct {
	ID      int
	Surname string
}

func (p Player) Id() int {
	return p.ID
}

func (p Player) Name() string {
	return "player"
}

func (p Player) Update() {}

func (p Player) Data() ecs.Data {
	return p
}

func (p Player) Type() reflect.Type {
	return reflect.TypeOf(p)
}
