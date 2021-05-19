package entity

type Entity struct {
	children   *Collection
	components map[string]Component
	name       string
}

func New(name string) *Entity {
	return &Entity{
		children:   NewCollection(),
		components: map[string]Component{},
		name:       name,
	}
}

func (entity *Entity) AddComponent(component Component) error {
	err := entity.validateNewComponent(component)
	if err != nil {
		return err
	}

	entity.components[component.Name()] = component
	return nil
}

func (entity Entity) GetComponent(name string) Component {
	return entity.components[name]
}

func (entity Entity) Name() string {
	return entity.name
}

func (entity Entity) RemoveComponent(name string) error {
	if !entity.hasComponent(name) {
		return ErrComponentDoesNotExist
	}
	delete(entity.components, name)
	return nil
}

func (entity *Entity) Update(dt float64) {
	for _, component := range entity.components {
		component.OnUpdate(entity, dt)
	}
}

func (entity Entity) hasComponent(name string) bool {
	_, ok := entity.components[name]
	return ok
}

func (entity Entity) validateNewComponent(component Component) error {
	if component.Name() == "" {
		return ErrComponentMustHaveType
	}
	if entity.hasComponent(component.Name()) {
		return ErrComponentMustBeUnique
	}
	return nil
}
