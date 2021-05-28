package behaviors

import "github.com/gomagedon/gophergame/engine"

const (
	TRANSFORM = "transform"
)

type Transform struct {
	x float64
	y float64
	w float64
	h float64

	entity *engine.Entity
}

func NewTransform(x, y, w, h float64) *Transform {
	return &Transform{
		x: x, y: y, w: w, h: h,
	}
}

func NewTransformFromBox(box engine.Box) *Transform {
	return &Transform{
		x: box.X, y: box.Y, w: box.W, h: box.H,
	}
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

func (t *Transform) Init(entity *engine.Entity) {
	t.entity = entity
}

func (t Transform) Update(dt float64) {
	// Needed to meet Behavior interface
}

func (t *Transform) Scale(vector engine.Vector) {
	t.w *= vector.X
	t.h *= vector.Y

	for _, transform := range t.getChildTransforms() {
		t.multiplyDistance(transform, vector)
		transform.Scale(vector)
	}
}

func (t *Transform) SetPosition(position engine.Vector) {
	t.x = position.X
	t.y = position.Y
}

func (t *Transform) SetScale(scale engine.Vector) {
	t.w = scale.X
	t.h = scale.Y
}

func (t *Transform) Translate(vector engine.Vector) {
	t.x += vector.X
	t.y += vector.Y

	for _, transform := range t.getChildTransforms() {
		transform.Translate(vector)
	}
}

func (t *Transform) multiplyDistance(transform *Transform, vector engine.Vector) {
	position := transform.GetPosition()

	xDiff := position.X - t.x
	yDiff := position.Y - t.y
	xTranslate := xDiff * (vector.X - 1)
	yTranslate := yDiff * (vector.Y - 1)

	transform.Translate(engine.Vector{X: xTranslate, Y: yTranslate})
}

func (t Transform) getChildTransforms() []*Transform {
	transforms := []*Transform{}
	if t.entity == nil {
		println("entity is nil")
	}
	for _, entity := range t.entity.Children().All() {
		if transform, ok := entity.GetBehavior(TRANSFORM).(*Transform); ok {
			transforms = append(transforms, transform)
		}
	}
	return transforms
}
