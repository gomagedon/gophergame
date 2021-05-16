package entity_test

import (
	"testing"

	"github.com/gomagedon/expectate"
	"github.com/gomagedon/gophergame/engine/canvas"
	"github.com/gomagedon/gophergame/engine/entity"
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

	entity1 := entity.New("foo")
	expect(entity1.Name()).ToBe("foo")

	entity2 := entity.New("bar")
	expect(entity2.Name()).ToBe("bar")
}

func TestEntity_AddRenderer_ReturnsErr_WhenRendererHasNoType(t *testing.T) {
	expect := expectate.Expect(t)

	myEntity := entity.New("my entity")
	renderer := &MockRenderer{typeName: ""}

	err := myEntity.AddRenderer(renderer)
	expect(err).ToBe(entity.ErrRendererMustHaveType)
}

func TestEntity_AddRenderer_ReturnsErr_WithDuplicateRenderer(t *testing.T) {
	expect := expectate.Expect(t)

	myEntity := entity.New("my entity")

	firstRenderer := &MockRenderer{typeName: "foo"}
	duplicateRenderer := &MockRenderer{typeName: "foo"}

	err := myEntity.AddRenderer(firstRenderer)
	expect(err).ToBe(nil)

	err = myEntity.AddRenderer(duplicateRenderer)
	expect(err).ToBe(entity.ErrRendererMustBeUnique)
}

func TestEntity_Draw_PassesCanvasToRenderers(t *testing.T) {
	expect := expectate.Expect(t)

	myEntity := entity.New("my entity")

	renderers := []*MockRenderer{
		{typeName: "type1"},
		{typeName: "type2"},
		{typeName: "type3"},
	}

	for _, renderer := range renderers {
		err := myEntity.AddRenderer(renderer)
		expect(err).ToBe(nil)
	}

	canvas := new(MockCanvas)
	myEntity.Draw(canvas)

	for _, renderer := range renderers {
		expect(renderer.Canvas).ToBe(canvas)
	}
}

func TestEntity_AddComponent_ReturnsErr_WhenComponentHasNoType(t *testing.T) {
	expect := expectate.Expect(t)

	myEntity := entity.New("my entity")
	component := &MockComponent{typeName: ""}

	err := myEntity.AddComponent(component)
	expect(err).ToBe(entity.ErrComponentMustHaveType)
}

func TestEntity_AddComponent_ReturnsErr_WithDuplicateComponent(t *testing.T) {
	expect := expectate.Expect(t)

	myEntity := entity.New("my entity")

	firstComponent := &MockComponent{typeName: "foo"}
	duplicateComponent := &MockComponent{typeName: "foo"}

	err := myEntity.AddComponent(firstComponent)
	expect(err).ToBe(nil)
	err = myEntity.AddComponent(duplicateComponent)
	expect(err).ToBe(entity.ErrComponentMustBeUnique)
}

func TestEntity_Update_PassesDeltaTimeToComponents(t *testing.T) {
	expect := expectate.Expect(t)

	entity := entity.New("my entity")

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
