package entity

import "github.com/gomagedon/gophergame/engine/canvas"

type Renderer interface {
	Type() string
	OnDraw(cvs canvas.Canvas)
}
