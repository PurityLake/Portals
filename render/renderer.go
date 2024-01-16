package render

import "github.com/veandco/go-sdl2/sdl"

type Renderer struct {
	renderer *sdl.Renderer
}

func (rend Renderer) SetDrawColor(r, g, b, a uint8) {
	rend.renderer.SetDrawColor(r, g, b, a)
}

func (rend Renderer) Clear() {
	rend.renderer.Clear()
}

func (rend Renderer) Present() {
	rend.renderer.Present()
}
