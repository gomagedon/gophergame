package entity_test

import (
	"testing"

	"github.com/gomagedon/expectate"
	"github.com/gomagedon/gophergame/engine/entity"
)

// MockComponent
type MockComponent struct {
	Parent    *entity.Entity
	WasCalled bool
	DeltaTime float64
	typeName  string
}

func (mock MockComponent) Type() string {
	return mock.typeName
}

func (mock *MockComponent) OnUpdate(parent *entity.Entity, dt float64) {
	mock.Parent = parent
	mock.DeltaTime = dt
}

func TestEntity(t *testing.T) {
	t.Run("Uses name in constructor", func(t *testing.T) {
		expect := expectate.Expect(t)

		entity1 := entity.New("foo")
		expect(entity1.Name()).ToBe("foo")

		entity2 := entity.New("bar")
		expect(entity2.Name()).ToBe("bar")
	})

	t.Run("AddComponent()", func(t *testing.T) {
		t.Run("Returns err when component has no type", func(t *testing.T) {
			expect := expectate.Expect(t)

			myEntity := entity.New("my entity")
			component := &MockComponent{typeName: ""}

			err := myEntity.AddComponent(component)
			expect(err).ToBe(entity.ErrComponentMustHaveType)
		})

		t.Run("Returns err with duplicate component", func(t *testing.T) {
			expect := expectate.Expect(t)

			myEntity := entity.New("my entity")

			firstComponent := &MockComponent{typeName: "foo"}
			duplicateComponent := &MockComponent{typeName: "foo"}

			err := myEntity.AddComponent(firstComponent)
			expect(err).ToBe(nil)
			err = myEntity.AddComponent(duplicateComponent)
			expect(err).ToBe(entity.ErrComponentMustBeUnique)
		})
	})

	t.Run("Update()", func(t *testing.T) {
		var expect expectate.ExpectorFunc
		var myEntity *entity.Entity
		var components []*MockComponent

		setup := func(t *testing.T) {
			expect = expectate.Expect(t)

			myEntity = entity.New("my entity")

			components = []*MockComponent{
				{typeName: "type1"},
				{typeName: "type2"},
				{typeName: "type3"},
			}

			for _, component := range components {
				myEntity.AddComponent(component)
			}
		}

		t.Run("Passes delta time to components", func(t *testing.T) {
			setup(t)

			myEntity.Update(123.0)
			for _, component := range components {
				expect(component.DeltaTime).ToBe(123.0)
			}

			myEntity.Update(99.0)
			for _, component := range components {
				expect(component.DeltaTime).ToBe(99.0)
			}
		})

		t.Run("Passes parent to components", func(t *testing.T) {
			setup(t)

			myEntity.Update(99.0)

			for _, component := range components {
				expect(component.Parent).ToBe(myEntity)
			}
		})
	})
}
