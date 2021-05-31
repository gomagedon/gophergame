package engine

type Behavior interface {
	Init(entity *Entity)
	Update(dt float64)
}

/* behaviorFunc */

type behaviorFunc struct {
	entity   *Entity
	onUpdate func(parent *Entity, dt float64)
}

func (b *behaviorFunc) Init(entity *Entity) {
	b.entity = entity
}

func (b behaviorFunc) Update(dt float64) {
	b.onUpdate(b.entity, dt)
}

func BehaviorFunc(onUpdate func(entity *Entity, dt float64)) Behavior {
	return &behaviorFunc{
		onUpdate: onUpdate,
	}
}
