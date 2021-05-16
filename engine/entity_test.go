package engine_test

import (
	"testing"

	"github.com/gomagedon/expectate"
	"github.com/gomagedon/gophergame/engine"
	"github.com/gomagedon/gophergame/engine/canvas"
)

/* Mocks */

// MockCanvas
type MockCanvas struct{}

func (MockCanvas) SetColor(x, y, w, h int) {}
func (MockCanvas) Clear()                  {}
func (MockCanvas) Present()                {}
func (MockCanvas) DrawRect(x, y, w, h int) {}

// MockRenderer
type MockRenderer struct {
	Canvas   canvas.Canvas
	typeName string
}

func (mock MockRenderer) Type() string {
	return mock.typeName
}

func (mock *MockRenderer) OnDraw(canvas canvas.Canvas) {
	mock.Canvas = canvas
}

// MockComponent
type MockComponent struct {
	WasCalled bool
	DeltaTime float64
	typeName  string
}

func (mock MockComponent) Type() string {
	return mock.typeName
}

func (mock *MockComponent) OnUpdate(dt float64) {
	mock.DeltaTime = dt
}

/* Tests */

func TestEntity_UsesNameFromConstructor(t *testing.T) {
	expect := expectate.Expect(t)

	entity1 := engine.NewEntity("foo")
	expect(entity1.Name()).ToBe("foo")

	entity2 := engine.NewEntity("bar")
	expect(entity2.Name()).ToBe("bar")
}

func TestEntity_AddRenderer_ReturnsErr_WhenRendererHasNoType(t *testing.T) {
	expect := expectate.Expect(t)

	entity := engine.NewEntity("my entity")
	renderer := &MockRenderer{typeName: ""}

	err := entity.AddRenderer(renderer)
	expect(err).ToBe(engine.ErrRendererMustHaveType)
}

func TestEntity_AddRenderer_ReturnsErr_WithDuplicateRenderer(t *testing.T) {
	expect := expectate.Expect(t)

	entity := engine.NewEntity("my entity")

	firstRenderer := &MockRenderer{typeName: "foo"}
	duplicateRenderer := &MockRenderer{typeName: "foo"}

	err := entity.AddRenderer(firstRenderer)
	expect(err).ToBe(nil)

	err = entity.AddRenderer(duplicateRenderer)
	expect(err).ToBe(engine.ErrRendererMustBeUnique)
}

func TestEntity_Draw_PassesCanvasToRenderers(t *testing.T) {
	expect := expectate.Expect(t)

	entity := engine.NewEntity("my entity")

	renderers := []*MockRenderer{
		{typeName: "type1"},
		{typeName: "type2"},
		{typeName: "type3"},
	}

	for _, renderer := range renderers {
		err := entity.AddRenderer(renderer)
		expect(err).ToBe(nil)
	}

	canvas := new(MockCanvas)
	entity.Draw(canvas)

	for _, renderer := range renderers {
		expect(renderer.Canvas).ToBe(canvas)
	}
}

func TestEntity_AddComponent_ReturnsErr_WhenComponentHasNoType(t *testing.T) {
	expect := expectate.Expect(t)

	entity := engine.NewEntity("my entity")
	component := &MockComponent{typeName: ""}

	err := entity.AddComponent(component)
	expect(err).ToBe(engine.ErrComponentMustHaveType)
}

func TestEntity_AddComponent_ReturnsErr_WithDuplicateComponent(t *testing.T) {
	expect := expectate.Expect(t)

	entity := engine.NewEntity("my entity")

	firstComponent := &MockComponent{typeName: "foo"}
	duplicateComponent := &MockComponent{typeName: "foo"}

	err := entity.AddComponent(firstComponent)
	expect(err).ToBe(nil)
	err = entity.AddComponent(duplicateComponent)
	expect(err).ToBe(engine.ErrComponentMustBeUnique)
}

func TestEntity_Update_PassesDeltaTimeToComponents(t *testing.T) {
	expect := expectate.Expect(t)

	entity := engine.NewEntity("my entity")

	components := []*MockComponent{
		{typeName: "type1"},
		{typeName: "type2"},
		{typeName: "type3"},
	}

	for _, component := range components {
		entity.AddComponent(component)
	}

	entity.Update(123.0)
	for _, component := range components {
		expect(component.DeltaTime).ToBe(123.0)
	}

	entity.Update(99.0)
	for _, component := range components {
		expect(component.DeltaTime).ToBe(99.0)
	}
}
