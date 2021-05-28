package platforms

import (
	"github.com/gomagedon/gophergame/canvas"
	"github.com/gomagedon/gophergame/controls"
	"github.com/veandco/go-sdl2/sdl"
)

type SDLFactory struct {
	cleanupCanvas func()
}

func NewSDLFactory() *SDLFactory {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	return new(SDLFactory)
}

func (fac *SDLFactory) GenerateCanvas(width int, height int) canvas.Canvas {
	canvas, cleanup := canvas.NewSDLCanvas(width, height)
	fac.cleanupCanvas = cleanup
	return canvas
}

func (fac SDLFactory) GenerateControls() *controls.SDLKeyboardControls {
	return controls.NewSDLKeyboardControls()
}

func (fac SDLFactory) Cleanup() {
	sdl.Quit()
	fac.cleanupCanvas()
}
