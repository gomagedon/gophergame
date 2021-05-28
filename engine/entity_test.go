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
	var expect expectate.ExpectorFunc

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
			expect(err).ToBe(engine.ErrBehaviorMustHaveUniqueName)
		})

		// Test
		t.Run("Inits behavior if InitBehavior", func(t *testing.T) {
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

		behaviors := []*MockBehavior{}
		for i := 0; i < 10; i++ {
			myEntity.AddBehavior(fmt.Sprint("behavior", i), new(MockBehavior))
		}

		myEntity.Update(123.0) // arbitrary number
		for _, behavior := range behaviors {
			expect(behavior.DeltaTime).ToBe(123.0)
		}

		myEntity.Update(99.0) // another arbitrary number
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
}
