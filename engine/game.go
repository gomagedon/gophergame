package engine

import (
	"github.com/gomagedon/gophergame/engine/canvas"
	"github.com/gomagedon/gophergame/engine/controls"
)

// Game ...
type Game struct {
	Canvas   canvas.Canvas
	Controls controls.Controls
	Entities EntityCollection
}
