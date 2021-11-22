package engine_test

import (
	"fmt"
	"testing"

	"github.com/gomagedon/expectate"
	"github.com/gomagedon/gophergame/engine"
)

// MockBehavior
type MockBehavior struct {
	Parent    *engine.Entity
	DeltaTime float64
}

func (mock *MockBehavior) Init(parent *engine.Entity) {
	mock.Parent = parent
}

func (mock *MockBehavior) Update(dt float64) {
	mock.DeltaTime = dt
}

func TestEntity(t *testing.T) {
	var expect expectate.ExpectorFunc // testing utility

	setup := func(t *testing.T) {
		expect = expectate.Expect(t)
	}

	// Test
	t.Run("Uses name in constructor", func(t *testing.T) {
		setup(t)

		entity1 := engine.NewEntity("foo")
		expect(entity1.Name()).ToBe("foo")

		entity2 := engine.NewEntity("bar")
		expect(entity2.Name()).ToBe("bar")
	})

	// Test
	t.Run("AddBehavior()", func(t *testing.T) {
		// Test
		t.Run("Returns err with duplicate behavior", func(t *testing.T) {
			setup(t)

			myEntity := engine.NewEntity("my entity")

			myEntity.AddBehavior("foo", new(MockBehavior))
			err := myEntity.AddBehavior("foo", new(MockBehavior))
			expect(err).ToBe(engine.ErrBehaviorIsNotUnique)
		})

		// Test
		t.Run("Inits behavior", func(t *testing.T) {
			setup(t)

			myEntity := engine.NewEntity("my entity")
			behavior := new(MockBehavior)
			myEntity.AddBehavior("mock behavior", behavior)

			expect(behavior.Parent).ToBe(myEntity)
		})
	})

	// Test
	t.Run("Update() passes delta time to behaviors", func(t *testing.T) {
		setup(t)

		myEntity := engine.NewEntity("my entity")
		behaviors := addMockBehaviors(myEntity, 10)

		myEntity.Update(3.14159)
		for _, behavior := range behaviors {
			expect(behavior.DeltaTime).ToBe(3.14159)
		}

		myEntity.Update(99.0)
		for _, behavior := range behaviors {
			expect(behavior.DeltaTime).ToBe(99.0)
		}
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

			behavior := new(MockBehavior)
			myEntity.AddBehavior("foo", behavior)

			expect(myEntity.GetBehavior("foo")).ToBe(behavior)
		})
	})

	// Test
	t.Run("RemoveBehavior()", func(t *testing.T) {
		// Test
		t.Run("Returns error if behavior doesn't exist", func(t *testing.T) {
			setup(t)

			myEntity := engine.NewEntity("my entity")

			err := myEntity.RemoveBehavior("non-existant")
			expect(err).ToBe(engine.ErrBehaviorDoesNotExist)
		})

		// Test
		t.Run("Removes behavior if exists", func(t *testing.T) {
			setup(t)

			myEntity := engine.NewEntity("my entity")
			myEntity.AddBehavior("foo", new(MockBehavior))

			myEntity.RemoveBehavior("foo")
			expect(myEntity.GetBehavior("foo")).ToBe(nil)
		})

		// Test
		t.Run("Removed behavior is not updated", func(t *testing.T) {
			setup(t)

			myEntity := engine.NewEntity("my entity")
			behavior := new(MockBehavior)
			myEntity.AddBehavior("foo", behavior)

			err := myEntity.RemoveBehavior("foo")
			expect(err).ToBe(nil)

			myEntity.Update(123.0)
			expect(behavior.DeltaTime).ToBe(0.0) // behavior was not updated
		})
	})

	// Test
	t.Run("Has a transform", func(t *testing.T) {
		// Test
		t.Run("That is not nil", func(t *testing.T) {
			setup(t)

			myEntity := engine.NewEntity("my entity")

			expect(myEntity.Transform()).NotToBe(nil)
		})

		// Test
		t.Run("With expected default values", func(t *testing.T) {
			setup(t)

			myEntity := engine.NewEntity("my entity")

			expect(myEntity.Transform().GetBox()).ToEqual(engine.Box{
				X: 1,
				Y: 1,
				W: 1,
				H: 1,
			})
		})
	})
}

func addMockBehaviors(entity *engine.Entity, count int) []*MockBehavior {
	behaviors := []*MockBehavior{}
	for i := 0; i < count; i++ {
		behaviors = append(behaviors, new(MockBehavior))
		entity.AddBehavior(fmt.Sprint("behavior", i), behaviors[i])
	}
	return behaviors
}
