package core

type Children struct {
	entities map[string]*Entity
}

func NewChildren() *Children {
	return &Children{
		entities: map[string]*Entity{},
	}
}

func (children *Children) Add(entity *Entity) error {
	if children.hasChild(entity.name) {
		return ErrDuplicateEntity
	}
	children.entities[entity.Name()] = entity
	return nil
}

func (children Children) All() []*Entity {
	entities := []*Entity{}
	for _, child := range children.entities {
		entities = append(entities, child)
	}
	return entities
}

func (children Children) Get(name string) *Entity {
	child, ok := children.entities[name]
	if !ok {
		return nil
	}
	return child
}

func (children Children) Remove(name string) error {
	if !children.hasChild(name) {
		return ErrNoSuchEntity
	}
	delete(children.entities, name)
	return nil
}

func (children Children) hasChild(name string) bool {
	_, ok := children.entities[name]
	return ok
}
