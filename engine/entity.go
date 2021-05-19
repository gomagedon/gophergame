package engine

type Entity struct {
	children  *Collection
	behaviors map[string]Behavior
	name      string
}

func NewEntity(name string) *Entity {
	return &Entity{
		children:  NewCollection(),
		behaviors: map[string]Behavior{},
		name:      name,
	}
}

func (entity *Entity) AddBehavior(behavior Behavior) error {
	err := entity.validateNewBehavior(behavior)
	if err != nil {
		return err
	}

	entity.behaviors[behavior.Name()] = behavior
	return nil
}

func (entity Entity) GetBehavior(name string) Behavior {
	return entity.behaviors[name]
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
		behavior.OnUpdate(entity, dt)
	}
}

func (entity Entity) hasBehavior(name string) bool {
	_, ok := entity.behaviors[name]
	return ok
}

func (entity Entity) validateNewBehavior(behavior Behavior) error {
	if behavior.Name() == "" {
		return ErrBehaviorMustHaveType
	}
	if entity.hasBehavior(behavior.Name()) {
		return ErrBehaviorMustBeUnique
	}
	return nil
}
