package entity

type Component interface {
	Type() string
	OnUpdate(dt float64)
}
