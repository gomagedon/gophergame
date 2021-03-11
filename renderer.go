package engine

// Renderer ...
type Renderer interface {
	Type() string
	OnDraw(cvs Canvas)
}
