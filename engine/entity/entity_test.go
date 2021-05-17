package entity_test

import (
	"testing"

	"github.com/gomagedon/expectate"
	"github.com/gomagedon/gophergame/engine/entity"
)

/* Mocks */

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
