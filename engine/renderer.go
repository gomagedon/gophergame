package engine

import "github.com/gomagedon/gophergame/engine/canvas"

// Renderer ...
type Renderer interface {
	Type() string
	OnDraw(cvs canvas.Canvas)
}
