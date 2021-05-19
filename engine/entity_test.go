package engine_test

import (
	"testing"

	"github.com/gomagedon/expectate"
	"github.com/gomagedon/gophergame/engine"
)

// MockBehavior
type MockBehavior struct {
	Parent    *engine.Entity
	WasCalled bool
	DeltaTime float64
	name      string
}

func (mock MockBehavior) Name() string {
	return mock.name
}

func (mock *MockBehavior) OnUpdate(parent *engine.Entity, dt float64) {
	mock.Parent = parent
	mock.DeltaTime = dt
}

func TestEntity(t *testing.T) {
	// Test
	t.Run("Uses name in constructor", func(t *testing.T) {
		expect := expectate.Expect(t)

		entity1 := engine.NewEntity("foo")
		expect(entity1.Name()).ToBe("foo")

		entity2 := engine.NewEntity("bar")
		expect(entity2.Name()).ToBe("bar")
	})

	// Test
	t.Run("AddBehavior()", func(t *testing.T) {
		t.Run("Returns err when behavior has no type", func(t *testing.T) {
			expect := expectate.Expect(t)

			myEntity := engine.NewEntity("my entity")
			behavior := &MockBehavior{name: ""}

			err := myEntity.AddBehavior(behavior)
			expect(err).ToBe(engine.ErrBehaviorMustHaveType)
		})

		// Test
		t.Run("Returns err with duplicate behavior", func(t *testing.T) {
			expect := expectate.Expect(t)

			myEntity := engine.NewEntity("my entity")

			firstBehavior := &MockBehavior{name: "foo"}
			duplicateBehavior := &MockBehavior{name: "foo"}

			err := myEntity.AddBehavior(firstBehavior)
			expect(err).ToBe(nil)
			err = myEntity.AddBehavior(duplicateBehavior)
			expect(err).ToBe(engine.ErrBehaviorMustBeUnique)
		})
	})

	// Test
	t.Run("Update()", func(t *testing.T) {
		var expect expectate.ExpectorFunc
		var myEntity *engine.Entity
		var behaviors []*MockBehavior

		setup := func(t *testing.T) {
			expect = expectate.Expect(t)

			myEntity = engine.NewEntity("my entity")

			behaviors = []*MockBehavior{
				{name: "type1"},
				{name: "type2"},
				{name: "type3"},
			}

			for _, behavior := range behaviors {
				myEntity.AddBehavior(behavior)
			}
		}

		// Test
		t.Run("Passes delta time to behaviors", func(t *testing.T) {
			setup(t)

			myEntity.Update(123.0)
			for _, behavior := range behaviors {
				expect(behavior.DeltaTime).ToBe(123.0)
			}

			myEntity.Update(99.0)
			for _, behavior := range behaviors {
				expect(behavior.DeltaTime).ToBe(99.0)
			}
		})

		// Test
		t.Run("Passes parent to behaviors", func(t *testing.T) {
			setup(t)

			myEntity.Update(99.0)

			for _, behavior := range behaviors {
				expect(behavior.Parent).ToBe(myEntity)
			}
		})
	})

	// Test
	t.Run("GetBehavior()", func(t *testing.T) {
		var expect expectate.ExpectorFunc
		var myEntity *engine.Entity

		setup := func(t *testing.T) {
			expect = expectate.Expect(t)
			myEntity = engine.NewEntity("my entity")
		}

		// Test
		t.Run("Returns nil if behavior does not exist", func(t *testing.T) {
			setup(t)

			behavior := myEntity.GetBehavior("non-existent")
			expect(behavior).ToBe(nil)
		})
		// Test
		t.Run("Returns behavior if exists", func(t *testing.T) {
			setup(t)

			fooBehavior := &MockBehavior{name: "foo"}
			myEntity.AddBehavior(fooBehavior)

			behavior := myEntity.GetBehavior("foo")
			expect(behavior).ToBe(fooBehavior)
		})
	})

	t.Run("RemoveBehavior()", func(t *testing.T) {
		t.Run("Returns error if behavior doesn't exist", func(t *testing.T) {
			expect := expectate.Expect(t)

			myEntity := engine.NewEntity("my entity")

			err := myEntity.RemoveBehavior("non-existant")

			expect(err).ToBe(engine.ErrBehaviorDoesNotExist)
		})

		t.Run("Removes behavior if exists", func(t *testing.T) {
			expect := expectate.Expect(t)

			myEntity := engine.NewEntity("my entity")
			myEntity.AddBehavior(&MockBehavior{name: "foo"})

			err := myEntity.RemoveBehavior("foo")
			expect(err).ToBe(nil)
			expect(myEntity.GetBehavior("foo")).ToBe(nil)
		})
	})
}
