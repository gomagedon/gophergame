package entity

type Entity struct {
	children   *Collection
	components []Component
	name       string
}

func New(name string) *Entity {
	return &Entity{
		children: NewCollection(),
		name:     name,
	}
}

func (entity *Entity) AddComponent(component Component) error {
	if component.Type() == "" {
		return ErrComponentMustHaveType
	}
	if entity.hasComponent(component) {
		return ErrComponentMustBeUnique
	}

	entity.components = append(entity.components, component)
	return nil
}

func (entity Entity) Name() string {
	return entity.name
}

func (entity *Entity) Update(dt float64) {
	for _, component := range entity.components {
		component.OnUpdate(entity, dt)
	}
}

func (e Entity) hasComponent(newComponent Component) bool {
	for _, component := range e.components {
		if newComponent.Type() == component.Type() {
			return true
		}
	}
	return false
}
