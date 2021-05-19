package engine

type Collection struct {
	children map[string]*Entity
}

func NewCollection() *Collection {
	return &Collection{
		children: map[string]*Entity{},
	}
}

func (collection *Collection) AddChild(entity *Entity) error {
	if collection.hasChild(entity.name) {
		return ErrDuplicateEntity
	}
	collection.children[entity.Name()] = entity
	return nil
}

func (collection Collection) GetChildren() []*Entity {
	children := []*Entity{}
	for _, child := range collection.children {
		children = append(children, child)
	}
	return children
}

func (collection Collection) GetChild(name string) *Entity {
	child, ok := collection.children[name]
	if !ok {
		return nil
	}
	return child
}

func (collection Collection) RemoveChild(name string) error {
	if !collection.hasChild(name) {
		return ErrNoSuchEntity
	}
	delete(collection.children, name)
	return nil
}

func (collection Collection) hasChild(name string) bool {
	_, ok := collection.children[name]
	return ok
}
