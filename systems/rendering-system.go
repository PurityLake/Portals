package systems

import (
	ecs "github.com/PurityLake/go-ecs"
	"github.com/portals/v2/components"
	"github.com/portals/v2/render"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type RenderSystem struct {
	id       int
	query    ecs.Query
	Renderer *render.Renderer
	Window   *render.Window
}

func (rs RenderSystem) Id() int {
	return rs.id
}

func (rs *RenderSystem) Setup(world *ecs.World) {
	if err := ttf.Init(); err != nil {
		panic(err)
	}
	rs.query = ecs.NewQuery(ecs.Type[components.Renderable](), ecs.Type[components.Position]())
}

func (rs RenderSystem) Update(world *ecs.World) {
	rs.Renderer.SetDrawColor(0, 0, 0, 255)
	rs.Renderer.Clear()
	componentsFound, found := world.Query(rs.query)
	if found {
		for _, componentList := range componentsFound {
			renderable := componentList[0].(components.Renderable)
			position := componentList[1].(components.Position)
			rs.Renderer.Renderer.Copy(renderable.Texture, nil,
				&sdl.Rect{
					X: int32(position.X),
					Y: int32(position.Y),
					W: int32(renderable.W),
					H: int32(renderable.H),
				})
		}
	}
	rs.Renderer.Present()
}
