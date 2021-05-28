package canvas

import "github.com/veandco/go-sdl2/sdl"

// SDLCanvas is the SDL2 implementation of Canvas
type SDLCanvas struct {
	color    struct{ r, g, b, a int }
	renderer *sdl.Renderer
}

func NewSDLCanvas(width int, height int) (canvas *SDLCanvas, cleanup func()) {
	canvas = new(SDLCanvas)
	window, err := sdl.CreateWindow(
		"test",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(width), int32(height),
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		panic(err)
	}
	canvas.renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	return canvas, func() {
		window.Destroy()
		canvas.renderer.Destroy()
	}
}

func (cvs *SDLCanvas) SetColor(r, g, b, a int) {
	cvs.color = struct {
		r int
		g int
		b int
		a int
	}{r, g, b, a}
}

func (cvs SDLCanvas) Clear() {
	cvs.renderer.SetDrawColor(0, 0, 0, 255)
	cvs.renderer.Clear()
}

func (cvs SDLCanvas) Present() {
	cvs.renderer.Present()
}

func (cvs SDLCanvas) DrawRect(x, y, w, h int) {
	cvs.renderer.SetDrawColor(
		uint8(cvs.color.r), uint8(cvs.color.g), uint8(cvs.color.b), uint8(cvs.color.a),
	)
	cvs.renderer.DrawRect(&sdl.Rect{
		X: int32(x),
		Y: int32(y),
		W: int32(w),
		H: int32(h),
	})
}
