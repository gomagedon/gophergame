package engine

type EntityCollection struct {
	children map[string]*Entity
}

func NewEntityCollection() *EntityCollection {
	return &EntityCollection{
		children: map[string]*Entity{},
	}
}

func (collection *EntityCollection) AddChild(entity *Entity) error {
	if collection.hasChild(entity.name) {
		return ErrDuplicateEntity
	}
	collection.children[entity.Name()] = entity
	return nil
}

func (collection EntityCollection) GetChildren() []*Entity {
	children := []*Entity{}
	for _, child := range collection.children {
		children = append(children, child)
	}
	return children
}

func (collection EntityCollection) GetChild(name string) *Entity {
	child, ok := collection.children[name]
	if !ok {
		return nil
	}
	return child
}

func (collection EntityCollection) RemoveChild(name string) error {
	if !collection.hasChild(name) {
		return ErrNoSuchEntity
	}
	delete(collection.children, name)
	return nil
}

func (collection EntityCollection) hasChild(name string) bool {
	_, ok := collection.children[name]
	return ok
}
