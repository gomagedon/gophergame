package entity_test

import (
	"testing"

	"github.com/gomagedon/expectate"
	"github.com/gomagedon/gophergame/engine/entity"
)

func TestCollection(t *testing.T) {
	var expect expectate.ExpectorFunc
	var collection *entity.Collection

	setup := func(t *testing.T) {
		expect = expectate.Expect(t)
		collection = entity.NewCollection()
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
			entitiesToAdd []*entity.Entity
		}

		tests := []Test{
			{
				name:          "one entity",
				entitiesToAdd: []*entity.Entity{entity.New("foo")},
			},
			{
				name: "two entities",
				entitiesToAdd: []*entity.Entity{
					entity.New("foo"),
					entity.New("bar"),
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

		collection.AddChild(entity.New("foo"))
		collection.AddChild(entity.New("bar"))
		collection.AddChild(entity.New("foobar"))

		children := collection.GetChildren()
		copyOfChildren := append([]*entity.Entity{}, children...)

		children[0] = nil
		children[1] = entity.New("random")

		for _, child := range collection.GetChildren() {
			expect(includes(copyOfChildren, child)).ToBe(true)
		}
	})

	// Test
	t.Run("AddChild() returns error when same entity is added twice", func(t *testing.T) {
		setup(t)

		myEntity := entity.New("my entity")

		err := collection.AddChild(myEntity)
		expect(err).ToBe(nil)
		err = collection.AddChild(myEntity)
		expect(err).ToBe(entity.ErrDuplicateEntity)
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

		fooEntity := entity.New("foo")
		barEntity := entity.New("bar")
		foobarEntity := entity.New("foobar")

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
		expect(err).ToBe(entity.ErrNoSuchEntity)
	})

	// Test
	t.Run("RemoveChild() removes child from collection", func(t *testing.T) {
		setup(t)

		myEntity := entity.New("my entity")
		collection.AddChild(myEntity)
		expect(collection.GetChild("my entity")).ToBe(myEntity)

		err := collection.RemoveChild("my entity")
		expect(err).ToBe(nil)
		if collection.GetChild("my entity") != nil {
			t.Fatalf("Child was not removed")
		}
	})
}

func includes(arr []*entity.Entity, entity *entity.Entity) bool {
	for _, element := range arr {
		if element == entity {
			return true
		}
	}
	return false
}
