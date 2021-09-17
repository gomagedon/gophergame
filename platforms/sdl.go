package platforms

import (
	"github.com/gomagedon/gophergame/canvas"
	"github.com/gomagedon/gophergame/controls"
	"github.com/veandco/go-sdl2/sdl"
)

type SDL struct {
	cleanupCanvas func()
}

func NewSDL() *SDL {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	return new(SDL)
}

func (fac *SDL) GenerateCanvas(width int, height int) canvas.Canvas {
	canvas, cleanup := canvas.NewSDLCanvas(width, height)
	fac.cleanupCanvas = cleanup
	return canvas
}

func (fac SDL) GenerateControls() *controls.SDLKeyboardControls {
	return controls.NewSDLKeyboardControls()
}

func (fac SDL) Cleanup() {
	sdl.Quit()
	fac.cleanupCanvas()
}
