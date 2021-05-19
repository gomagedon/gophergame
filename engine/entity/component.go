package entity

type Component interface {
	Name() string
	OnUpdate(parent *Entity, dt float64)
}
