package engine_test

import (
	"testing"

	"github.com/gomagedon/expectate"
	"github.com/gomagedon/gophergame/engine"
)

func TestEntityCollection(t *testing.T) {
	var expect expectate.ExpectorFunc
	var collection *engine.EntityCollection

	setup := func(t *testing.T) {
		expect = expectate.Expect(t)
		collection = engine.NewEntityCollection()
	}

	// Test
	t.Run("Is empty when created", func(t *testing.T) {
		setup(t)

		expect(len(collection.GetChildren())).ToBe(0)
	})

	// Test
	t.Run("GetChildren() returns entities added to it", func(t *testing.T) {
		type Test struct {
			name          string
			entitiesToAdd []*engine.Entity
		}

		tests := []Test{
			{
				name:          "one entity",
				entitiesToAdd: []*engine.Entity{engine.NewEntity("foo")},
			},
			{
				name: "two entities",
				entitiesToAdd: []*engine.Entity{
					engine.NewEntity("foo"),
					engine.NewEntity("bar"),
				},
			},
		}

		for _, tc := range tests {
			// Test
			t.Run(tc.name, func(t *testing.T) {
				setup(t)

				for _, entity := range tc.entitiesToAdd {
					collection.AddChild(entity)
				}

				children := collection.GetChildren()
				expect(len(children)).ToBe(len(tc.entitiesToAdd))

				for _, entity := range tc.entitiesToAdd {
					if !includes(children, entity) {
						t.Fatalf("Added entity was not present in children")
					}
				}
			})
		}
	})

	// Test
	t.Run("GetChildren() is immutable", func(t *testing.T) {
		setup(t)

		collection.AddChild(engine.NewEntity("foo"))
		collection.AddChild(engine.NewEntity("bar"))
		collection.AddChild(engine.NewEntity("foobar"))

		children := collection.GetChildren()
		copyOfChildren := append([]*engine.Entity{}, children...)

		children[0] = nil
		children[1] = engine.NewEntity("random")

		for _, child := range collection.GetChildren() {
			expect(includes(copyOfChildren, child)).ToBe(true)
		}
	})

	// Test
	t.Run("AddChild() returns error when same entity is added twice", func(t *testing.T) {
		setup(t)

		entity := engine.NewEntity("foo")

		err := collection.AddChild(entity)
		expect(err).ToBe(nil)
		err = collection.AddChild(entity)
		expect(err).ToBe(engine.ErrDuplicateEntity)
	})

	// Test
	t.Run("GetChild() returns nil if child does not exist", func(t *testing.T) {
		setup(t)

		child := collection.GetChild("non-existent")
		if child != nil {
			t.Fatalf("Expected child to be nil; Got: %v", child)
		}
	})

	// Test
	t.Run("GetChild() returns added child with name", func(t *testing.T) {
		setup(t)

		fooEntity := engine.NewEntity("foo")
		barEntity := engine.NewEntity("bar")
		foobarEntity := engine.NewEntity("foobar")

		collection.AddChild(fooEntity)
		collection.AddChild(barEntity)
		collection.AddChild(foobarEntity)

		expect(collection.GetChild("foo")).ToBe(fooEntity)
		expect(collection.GetChild("bar")).ToBe(barEntity)
		expect(collection.GetChild("foobar")).ToBe(foobarEntity)
	})

	// Test
	t.Run("RemoveChild() returns error if child does not exist", func(t *testing.T) {
		setup(t)

		err := collection.RemoveChild("non-existent")
		expect(err).ToBe(engine.ErrNoSuchEntity)
	})

	// Test
	t.Run("RemoveChild() removes child from collection", func(t *testing.T) {
		setup(t)

		myEntity := engine.NewEntity("my entity")
		collection.AddChild(myEntity)
		expect(collection.GetChild("my entity")).ToBe(myEntity)

		err := collection.RemoveChild("my entity")
		expect(err).ToBe(nil)
		if collection.GetChild("my entity") != nil {
			t.Fatalf("Child was not removed")
		}
	})
}

func includes(arr []*engine.Entity, entity *engine.Entity) bool {
	for _, element := range arr {
		if element == entity {
			return true
		}
	}
	return false
}
