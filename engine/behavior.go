package engine

type Behavior interface {
	Name() string
	OnUpdate(parent *Entity, dt float64)
}

type behaviorFunc struct {
	name     string
	onUpdate func(parent *Entity, dt float64)
}

func (b behaviorFunc) Name() string {
	return b.name
}

func (b behaviorFunc) OnUpdate(parent *Entity, dt float64) {
	b.onUpdate(parent, dt)
}

func BehaviorFunc(name string, onUpdate func(parent *Entity, dt float64)) Behavior {
	return &behaviorFunc{
		name:     name,
		onUpdate: onUpdate,
	}
}
