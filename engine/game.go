package engine

import (
	"github.com/gomagedon/gophergame/engine/canvas"
	"github.com/gomagedon/gophergame/engine/controls"
)

// Game ...
type Game struct {
	Controls controls.Controls
	Canvas   canvas.Canvas
}
