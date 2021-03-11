package engine_test

import (
	"testing"

	"github.com/gomagedon/expectate"
	"github.com/steve-kaufman/sdl-game/engine"
)

type DummyRenderer struct {
	wasCalled bool
	rType     string
}

func (dummy DummyRenderer) Type() string {
	return dummy.rType
}

func (dummy *DummyRenderer) OnDraw(engine.Canvas) {
	dummy.wasCalled = true
}

type DummyComponent struct {
	wasCalled       bool
	wasCalledWithDT float64
	rType           string
}

func (dummy DummyComponent) Type() string {
	return dummy.rType
}

func (dummy *DummyComponent) OnUpdate(dt float64) {
	dummy.wasCalled = true
	dummy.wasCalledWithDT = dt
}

type MockCanvas struct{}

func (MockCanvas) SetColor(x, y, w, h int) {}
func (MockCanvas) Clear()                  {}
func (MockCanvas) Present()                {}
func (MockCanvas) DrawRect(x, y, w, h int) {}

func TestEntity_AddRenderer_MustHaveType(t *testing.T) {
	expect := expectate.Expect(t)

	entity := engine.NewEntity()
	dummy := new(DummyRenderer)

	err := entity.AddRenderer(dummy)

	expect(err).ToBe(engine.ErrRendererMustHaveType)
}

func TestEntity_AddRenderer_MustBeUniqueType(t *testing.T) {
	expect := expectate.Expect(t)

	entity := engine.NewEntity()

	dummy1 := new(DummyRenderer)
	dummy1.rType = "type1"
	dummy2 := new(DummyRenderer)
	dummy2.rType = "type1"

	entity.AddRenderer(dummy1)
	err := entity.AddRenderer(dummy2)

	expect(err).ToBe(engine.ErrRendererMustBeUnique)
}

func TestEntity_CallsOnDraw_WithOneRenderer(t *testing.T) {
	expect := expectate.Expect(t)

	entity := engine.NewEntity()
	dummy := new(DummyRenderer)
	dummy.rType = "foo"
	entity.AddRenderer(dummy)

	entity.Draw(new(MockCanvas))

	expect(dummy.wasCalled).ToBe(true)
}

func TestEntity_CallsOnDraw_WithManyRenderers(t *testing.T) {
	expect := expectate.Expect(t)

	entity := engine.NewEntity()

	dummy1 := new(DummyRenderer)
	dummy1.rType = "type1"
	dummy2 := new(DummyRenderer)
	dummy2.rType = "type2"
	dummy3 := new(DummyRenderer)
	dummy3.rType = "type3"
	dummy4 := new(DummyRenderer)
	dummy4.rType = "type4"

	entity.AddRenderer(dummy1)
	entity.AddRenderer(dummy2)
	entity.AddRenderer(dummy3)
	entity.AddRenderer(dummy4)

	entity.Draw(new(MockCanvas))

	expect(dummy1.wasCalled).ToBe(true)
	expect(dummy2.wasCalled).ToBe(true)
	expect(dummy3.wasCalled).ToBe(true)
	expect(dummy4.wasCalled).ToBe(true)
}

func TestEntity_AddComponent_FailsWithNoType(t *testing.T) {
	expect := expectate.Expect(t)

	entity := engine.NewEntity()
	dummy := new(DummyComponent)

	err := entity.AddComponent(dummy)

	expect(err).ToBe(engine.ErrComponentMustHaveType)
}

func TestEntity_AddComponent_SucceedsWithType(t *testing.T) {
	expect := expectate.Expect(t)

	entity := engine.NewEntity()
	dummy := new(DummyComponent)
	dummy.rType = "foo"

	err := entity.AddComponent(dummy)

	expect(err).ToBe(nil)
}

func TestEntity_AddComponent_FailsWithSameType(t *testing.T) {
	expect := expectate.Expect(t)

	entity := engine.NewEntity()

	dummy1 := new(DummyComponent)
	dummy1.rType = "type"
	dummy2 := new(DummyComponent)
	dummy2.rType = "type"
	dummy3 := new(DummyComponent)
	dummy3.rType = "type"

	entity.AddComponent(dummy1)
	err := entity.AddComponent(dummy2)
	err2 := entity.AddComponent(dummy3)

	expect(err).ToBe(engine.ErrComponentMustBeUnique)
	expect(err2).ToBe(engine.ErrComponentMustBeUnique)
}

func TestEntity_CallsOnUpdate_WithOneComponent(t *testing.T) {
	expect := expectate.Expect(t)

	entity := engine.NewEntity()
	dummy := new(DummyComponent)
	dummy.rType = "foo"
	entity.AddComponent(dummy)

	entity.Update(1.0)

	expect(dummy.wasCalled).ToBe(true)
	expect(dummy.wasCalledWithDT).ToBe(1.0)
}

func TestEntity_CallsOnUpdate_WithManyComponents(t *testing.T) {
	expect := expectate.Expect(t)

	entity := engine.NewEntity()

	dummy1 := new(DummyComponent)
	dummy1.rType = "type1"
	dummy2 := new(DummyComponent)
	dummy2.rType = "type2"
	dummy3 := new(DummyComponent)
	dummy3.rType = "type3"
	dummy4 := new(DummyComponent)
	dummy4.rType = "type4"

	entity.AddComponent(dummy1)
	entity.AddComponent(dummy2)
	entity.AddComponent(dummy3)
	entity.AddComponent(dummy4)

	entity.Update(1.0)

	expect(dummy1.wasCalled).ToBe(true)
	expect(dummy1.wasCalledWithDT).ToBe(1.0)
	expect(dummy2.wasCalled).ToBe(true)
	expect(dummy2.wasCalledWithDT).ToBe(1.0)
	expect(dummy3.wasCalled).ToBe(true)
	expect(dummy3.wasCalledWithDT).ToBe(1.0)
	expect(dummy4.wasCalled).ToBe(true)
	expect(dummy4.wasCalledWithDT).ToBe(1.0)
}
