package engine

// NewEntity ...
func NewEntity() *Entity {
	return new(Entity)
}

// Entity ...
type Entity struct {
	renderers  []Renderer
	components []Component
}

// AddRenderer ...
func (e *Entity) AddRenderer(renderer Renderer) error {
	if renderer.Type() == "" {
		return ErrRendererMustHaveType
	}
	if e.isRendererAlreadyAttached(renderer) {
		return ErrRendererMustBeUnique
	}

	e.renderers = append(e.renderers, renderer)
	return nil
}

// AddComponent ...
func (e *Entity) AddComponent(component Component) error {
	if component.Type() == "" {
		return ErrComponentMustHaveType
	}
	if e.isComponentAlreadyAttached(component) {
		return ErrComponentMustBeUnique
	}

	e.components = append(e.components, component)
	return nil
}

func (e Entity) isRendererAlreadyAttached(newRenderer Renderer) bool {
	for _, renderer := range e.renderers {
		if newRenderer.Type() == renderer.Type() {
			return true
		}
	}
	return false
}

func (e Entity) isComponentAlreadyAttached(newComponent Component) bool {
	for _, component := range e.components {
		if newComponent.Type() == component.Type() {
			return true
		}
	}
	return false
}

// Draw ...
func (e Entity) Draw(canvas Canvas) {
	for _, renderer := range e.renderers {
		renderer.OnDraw(canvas)
	}
}

// Update ...
func (e Entity) Update(dt float64) {
	for _, component := range e.components {
		component.OnUpdate(dt)
	}
}
