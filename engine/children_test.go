package engine_test

import (
	"testing"

	"github.com/gomagedon/expectate"
	"github.com/gomagedon/gophergame/engine"
)

func TestChildren(t *testing.T) {
	var expect expectate.ExpectorFunc // testing utility
	var children *engine.Children

	setup := func(t *testing.T) {
		expect = expectate.Expect(t)
		children = engine.NewChildren()
	}

	// Test
	t.Run("Is empty when created", func(t *testing.T) {
		setup(t)

		expect(len(children.All())).ToBe(0)
	})

	// Test
	t.Run("All() returns entities added to it", func(t *testing.T) {
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
					children.Add(entity)
				}

				children := children.All()
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
	t.Run("All() is immutable", func(t *testing.T) {
		setup(t)

		children.Add(engine.NewEntity("foo"))
		children.Add(engine.NewEntity("bar"))
		children.Add(engine.NewEntity("foobar"))

		entities := children.All()
		copyOfEntities := append([]*engine.Entity{}, entities...)

		entities[0] = nil
		entities[1] = engine.NewEntity("random")

		for _, entity := range children.All() {
			expect(includes(copyOfEntities, entity)).ToBe(true)
		}
	})

	// Test
	t.Run("Add() returns error when same entity is added twice", func(t *testing.T) {
		setup(t)

		myEntity := engine.NewEntity("my entity")

		err := children.Add(myEntity)
		expect(err).ToBe(nil)
		err = children.Add(myEntity)
		expect(err).ToBe(engine.ErrDuplicateEntity)
	})

	// Test
	t.Run("Get() returns nil if child does not exist", func(t *testing.T) {
		setup(t)

		child := children.Get("non-existent")
		if child != nil {
			t.Fatalf("Expected child to be nil; Got: %v", child)
		}
	})

	// Test
	t.Run("Get() returns added child with name", func(t *testing.T) {
		setup(t)

		fooEntity := engine.NewEntity("foo")
		barEntity := engine.NewEntity("bar")
		foobarEntity := engine.NewEntity("foobar")

		children.Add(fooEntity)
		children.Add(barEntity)
		children.Add(foobarEntity)

		expect(children.Get("foo")).ToBe(fooEntity)
		expect(children.Get("bar")).ToBe(barEntity)
		expect(children.Get("foobar")).ToBe(foobarEntity)
	})

	// Test
	t.Run("Remove() returns error if child does not exist", func(t *testing.T) {
		setup(t)

		err := children.Remove("non-existent")
		expect(err).ToBe(engine.ErrNoSuchEntity)
	})

	// Test
	t.Run("Remove() removes child from children", func(t *testing.T) {
		setup(t)

		myEntity := engine.NewEntity("my entity")
		children.Add(myEntity)
		expect(children.Get("my entity")).ToBe(myEntity)

		err := children.Remove("my entity")
		expect(err).ToBe(nil)
		if children.Get("my entity") != nil {
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
