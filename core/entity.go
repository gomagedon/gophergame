package core

type Entity struct {
	behaviors map[string]Behavior
	children  *Children
	transform *Transform
	name      string
}

func NewEntity(name string) *Entity {
	entity := buildEntity(name)
	entity.addTransform()
	return entity
}

func buildEntity(name string) *Entity {
	return &Entity{
		children:  NewChildren(),
		behaviors: map[string]Behavior{},
		name:      name,
	}
}

func (entity *Entity) addTransform() {
	entity.transform = NewTransform(1, 1, 1, 1)
	entity.AddBehavior("transform", entity.transform)
}

func (entity *Entity) AddBehavior(name string, behavior Behavior) error {
	if err := entity.validateBehaviorName(name); err != nil {
		return err
	}

	entity.behaviors[name] = behavior
	behavior.Init(entity)

	return nil
}

func (entity Entity) GetBehavior(name string) Behavior {
	return entity.behaviors[name]
}

func (entity Entity) Transform() *Transform {
	return entity.transform
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
		return ErrBehaviorIsNotUnique
	}
	return nil
}
