package engine

type Entity struct {
	children  *Children
	behaviors map[string]Behavior
	name      string
}

func NewEntity(name string) *Entity {
	return &Entity{
		children:  NewChildren(),
		behaviors: map[string]Behavior{},
		name:      name,
	}
}

func (entity *Entity) AddBehavior(name string, behavior Behavior) error {
	err := entity.validateBehaviorName(name)
	if err != nil {
		return err
	}

	entity.behaviors[name] = behavior
	behavior.Init(entity)

	return nil
}

func (entity Entity) GetBehavior(name string) Behavior {
	return entity.behaviors[name]
}

func (entity Entity) Children() *Children {
	return entity.children
}

func (entity Entity) Name() string {
	return entity.name
}

func (entity Entity) RemoveBehavior(name string) error {
	if !entity.hasBehavior(name) {
		return ErrBehaviorDoesNotExist
	}
	delete(entity.behaviors, name)
	return nil
}

func (entity *Entity) Update(dt float64) {
	for _, behavior := range entity.behaviors {
		behavior.Update(dt)
	}
}

func (entity Entity) hasBehavior(name string) bool {
	_, ok := entity.behaviors[name]
	return ok
}

func (entity Entity) validateBehaviorName(name string) error {
	if name == "" {
		return ErrBehaviorMustHaveName
	}
	if entity.hasBehavior(name) {
		return ErrBehaviorMustHaveUniqueName
	}
	return nil
}
