package entity

type Behavior interface {
	Name() string
	OnUpdate(parent *Entity, dt float64)
}
