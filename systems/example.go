package systems

import (
	"fmt"

	ecs "github.com/PurityLake/go-ecs"
	"github.com/portals/v2/components"
)

type ExampleSystem struct {
	// data can go here
}

func (s ExampleSystem) Setup(world *ecs.World) {
	fmt.Println("ExampleSystem setup")
	world.AddEntity("example", components.Renderable{})
}

func (s ExampleSystem) Update(world *ecs.World) {
}
