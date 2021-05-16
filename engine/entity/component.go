package entity

// Component ...
type Component interface {
	Type() string
	OnUpdate(dt float64)
}
