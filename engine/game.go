package engine

import (
	"github.com/gomagedon/gophergame/engine/canvas"
	"github.com/gomagedon/gophergame/engine/controls"
	"github.com/gomagedon/gophergame/engine/entity"
)

// Game ...
type Game struct {
	Canvas   canvas.Canvas
	Controls controls.Controls
	Entities entity.Collection
}
