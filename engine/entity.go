package engine

import "github.com/gomagedon/gophergame/engine/canvas"

type Entity struct {
	renderers  []Renderer
	components []Component
}

func NewEntity() *Entity {
	return new(Entity)
}

func (entity *Entity) AddRenderer(renderer Renderer) error {
	if renderer.Type() == "" {
		return ErrRendererMustHaveType
	}
	if entity.isRendererAlreadyAttached(renderer) {
		return ErrRendererMustBeUnique
	}

	entity.renderers = append(entity.renderers, renderer)
	return nil
}

func (entity *Entity) AddComponent(component Component) error {
	if component.Type() == "" {
		return ErrComponentMustHaveType
	}
	if entity.isComponentAlreadyAttached(component) {
		return ErrComponentMustBeUnique
	}

	entity.components = append(entity.components, component)
	return nil
}

func (entity Entity) Draw(canvas canvas.Canvas) {
	for _, renderer := range entity.renderers {
		renderer.OnDraw(canvas)
	}
}

func (entity Entity) Update(dt float64) {
	for _, component := range entity.components {
		component.OnUpdate(dt)
	}
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
