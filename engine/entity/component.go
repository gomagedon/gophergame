package entity

type Component interface {
	Type() string
	OnUpdate(parent *Entity, dt float64)
}
