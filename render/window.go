package render

import "github.com/veandco/go-sdl2/sdl"

type Window struct {
	width, height int
	title         string
	window        *sdl.Window
}

func NewWindow(title string, width, height int) *Window {
	return &Window{
		width:  width,
		height: height,
		title:  title,
	}
}

func (w Window) CreateRenderer() (*sdl.Renderer, error) {
	return sdl.CreateRenderer(w.window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
}

func (w *Window) Init() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	}

	window, err := sdl.CreateWindow("Portals", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		return err
	}

	w.window = window

	return nil
}
