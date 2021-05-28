package behaviors

import "github.com/gomagedon/gophergame/engine"

type Transform struct {
	x float64
	y float64
	w float64
	h float64
}

func NewTransform(x, y, w, h float64) *Transform {
	return &Transform{x, y, w, h}
}

func (Transform) Name() string {
	return "transform"
}

func (t Transform) GetPosition() engine.Vector {
	return engine.Vector{X: t.x, Y: t.y}
}

func (t Transform) GetScale() engine.Vector {
	return engine.Vector{X: t.w, Y: t.h}
}

func (t Transform) GetBox() engine.Box {
	return engine.Box{X: t.x, Y: t.y, W: t.w, H: t.h}
}

func (t *Transform) Translate(vector engine.Vector) {
	t.x += vector.X
	t.y += vector.Y
}

func (Transform) OnUpdate(parent *engine.Entity, dt float64) {}
