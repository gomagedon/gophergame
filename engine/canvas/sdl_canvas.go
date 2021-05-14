package canvas

import "github.com/veandco/go-sdl2/sdl"

// SDLCanvas ...
type SDLCanvas struct {
	renderer *sdl.Renderer
}

// NewSDLCanvas ...
func NewSDLCanvas(renderer *sdl.Renderer) *SDLCanvas {
	canvas := new(SDLCanvas)
	canvas.renderer = renderer
	return canvas
}

// SetColor ...
func (cvs SDLCanvas) SetColor(r, g, b, a int) {
	cvs.renderer.SetDrawColor(uint8(r), uint8(g), uint8(b), uint8(a))
}

// Clear ...
func (cvs SDLCanvas) Clear() {
	cvs.renderer.Clear()
}

// Present ...
func (cvs SDLCanvas) Present() {
	cvs.renderer.Present()
}

// DrawRect ...
func (cvs SDLCanvas) DrawRect(x, y, w, h int) {
	cvs.renderer.DrawRect(&sdl.Rect{
		X: int32(x),
		Y: int32(y),
		W: int32(w),
		H: int32(h),
	})
}
