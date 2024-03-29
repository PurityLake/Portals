package components

import (
	"reflect"

	ecs "github.com/PurityLake/go-ecs"
)

type Position struct {
	ID   int
	X, Y int
}

func (p Position) Id() int {
	return p.ID
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
