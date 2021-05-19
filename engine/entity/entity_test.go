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
	name      string
}

func (mock MockComponent) Name() string {
	return mock.name
}

func (mock *MockComponent) OnUpdate(parent *entity.Entity, dt float64) {
	mock.Parent = parent
	mock.DeltaTime = dt
}

func TestEntity(t *testing.T) {
	// Test
	t.Run("Uses name in constructor", func(t *testing.T) {
		expect := expectate.Expect(t)

		entity1 := entity.New("foo")
		expect(entity1.Name()).ToBe("foo")

		entity2 := entity.New("bar")
		expect(entity2.Name()).ToBe("bar")
	})

	// Test
	t.Run("AddComponent()", func(t *testing.T) {
		t.Run("Returns err when component has no type", func(t *testing.T) {
			expect := expectate.Expect(t)

			myEntity := entity.New("my entity")
			component := &MockComponent{name: ""}

			err := myEntity.AddComponent(component)
			expect(err).ToBe(entity.ErrComponentMustHaveType)
		})

		// Test
		t.Run("Returns err with duplicate component", func(t *testing.T) {
			expect := expectate.Expect(t)

			myEntity := entity.New("my entity")

			firstComponent := &MockComponent{name: "foo"}
			duplicateComponent := &MockComponent{name: "foo"}

			err := myEntity.AddComponent(firstComponent)
			expect(err).ToBe(nil)
			err = myEntity.AddComponent(duplicateComponent)
			expect(err).ToBe(entity.ErrComponentMustBeUnique)
		})
	})

	// Test
	t.Run("Update()", func(t *testing.T) {
		var expect expectate.ExpectorFunc
		var myEntity *entity.Entity
		var components []*MockComponent

		setup := func(t *testing.T) {
			expect = expectate.Expect(t)

			myEntity = entity.New("my entity")

			components = []*MockComponent{
				{name: "type1"},
				{name: "type2"},
				{name: "type3"},
			}

			for _, component := range components {
				myEntity.AddComponent(component)
			}
		}

		// Test
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

		// Test
		t.Run("Passes parent to components", func(t *testing.T) {
			setup(t)

			myEntity.Update(99.0)

			for _, component := range components {
				expect(component.Parent).ToBe(myEntity)
			}
		})
	})

	// Test
	t.Run("GetComponent()", func(t *testing.T) {
		var expect expectate.ExpectorFunc
		var myEntity *entity.Entity

		setup := func(t *testing.T) {
			expect = expectate.Expect(t)
			myEntity = entity.New("my entity")
		}

		// Test
		t.Run("Returns nil if component does not exist", func(t *testing.T) {
			setup(t)

			component := myEntity.GetComponent("non-existent")
			expect(component).ToBe(nil)
		})
		// Test
		t.Run("Returns component if exists", func(t *testing.T) {
			setup(t)

			fooComponent := &MockComponent{name: "foo"}
			myEntity.AddComponent(fooComponent)

			component := myEntity.GetComponent("foo")
			expect(component).ToBe(fooComponent)
		})
	})

	t.Run("RemoveComponent()", func(t *testing.T) {
		t.Run("Returns error if component doesn't exist", func(t *testing.T) {
			expect := expectate.Expect(t)

			myEntity := entity.New("my entity")

			err := myEntity.RemoveComponent("non-existant")

			expect(err).ToBe(entity.ErrComponentDoesNotExist)
		})

		t.Run("Removes component if exists", func(t *testing.T) {
			expect := expectate.Expect(t)

			myEntity := entity.New("my entity")
			myEntity.AddComponent(&MockComponent{name: "foo"})

			err := myEntity.RemoveComponent("foo")
			expect(err).ToBe(nil)
			expect(myEntity.GetComponent("foo")).ToBe(nil)
		})
	})
}
