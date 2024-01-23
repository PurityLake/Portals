package render

import "github.com/veandco/go-sdl2/sdl"

type Renderer struct {
	Renderer *sdl.Renderer
}

func (rend Renderer) SetDrawColor(r, g, b, a uint8) {
	rend.Renderer.SetDrawColor(r, g, b, a)
}

func (rend Renderer) Clear() {
	rend.Renderer.Clear()
}

func (rend Renderer) Present() {
	rend.Renderer.Present()
}
