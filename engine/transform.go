package engine

type Transform struct {
	x float64
	y float64
	w float64
	h float64

	entity *Entity
}

func NewTransform(x, y, w, h float64) *Transform {
	return &Transform{
		x: x, y: y, w: w, h: h,
	}
}

func NewTransformFromBox(box Box) *Transform {
	return &Transform{
		x: box.X, y: box.Y, w: box.W, h: box.H,
	}
}

/* Behavior methods */

func (t *Transform) Init(entity *Entity) {
	t.entity = entity
}

func (t Transform) Update(dt float64) {
	// Needed to meet Behavior interface
}

/* Public methods */

func (t Transform) GetPosition() Vector {
	return Vector{X: t.x, Y: t.y}
}

func (t Transform) GetSize() Vector {
	return Vector{X: t.w, Y: t.h}
}

func (t Transform) GetBox() Box {
	return Box{X: t.x, Y: t.y, W: t.w, H: t.h}
}

func (t *Transform) Scale(vector Vector) {
	t.w *= vector.X
	t.h *= vector.Y
}

func (t *Transform) SetPosition(newPosition Vector) {
	t.x = newPosition.X
	t.y = newPosition.Y
}

func (t *Transform) SetSize(newScale Vector) {
	t.w = newScale.X
	t.h = newScale.Y
}

func (t *Transform) Translate(vector Vector) {
	t.x += vector.X
	t.y += vector.Y
}
