package core_test

import (
	"testing"

	"github.com/gomagedon/expectate"
	"github.com/gomagedon/gophergame/core"
)

func TestTransform(t *testing.T) {
	var expect expectate.ExpectorFunc // testing utility

	setup := func(t *testing.T) {
		expect = expectate.Expect(t)
	}

	// Test
	t.Run("GetPosition() returns its position", func(t *testing.T) {
		setup(t)

		transform := core.NewTransform(10.0, 20.0, 0, 0)
		position := transform.GetPosition()

		expect(position.X).ToBe(10.0)
		expect(position.Y).ToBe(20.0)
	})

	// Test
	t.Run("GetSize() returns its size", func(t *testing.T) {
		setup(t)

		transform := core.NewTransform(0, 0, 11.0, 21.0)
		scale := transform.GetSize()

		expect(scale.X).ToBe(11.0)
		expect(scale.Y).ToBe(21.0)
	})

	// Test
	t.Run("GetBox() returns x, y, w, and h", func(t *testing.T) {
		setup(t)

		transform := core.NewTransform(5.0, 10.0, 15.0, 20.0)
		box := transform.GetBox()
		expect(box).ToEqual(core.Box{
			X: 5.0,
			Y: 10.0,
			W: 15.0,
			H: 20.0,
		})
	})

	// Test
	t.Run("Translate() changes its position", func(t *testing.T) {
		setup(t)

		transform := core.NewTransform(10, 20, 0, 0)

		transform.Translate(core.Vector{X: 3, Y: -5})
		position := transform.GetPosition()

		expect(position.X).ToBe(13.0)
		expect(position.Y).ToBe(15.0)
	})

	// Test
	t.Run("Scale() multiplies its size", func(t *testing.T) {
		setup(t)

		transform := core.NewTransform(0, 0, 20.0, 40.0)

		transform.Scale(core.Vector{X: 2.5, Y: 0.25})
		scale := transform.GetSize()

		expect(scale.X).ToBe(50.0)
		expect(scale.Y).ToBe(10.0)
	})

	// Test
	t.Run("SetPosition() sets its position", func(t *testing.T) {
		setup(t)

		transform := core.NewTransform(0, 0, 0, 0)

		transform.SetPosition(core.Vector{X: 11.0, Y: 30.0})
		position := transform.GetPosition()

		expect(position.X).ToBe(11.0)
		expect(position.Y).ToBe(30.0)
	})

	// Test
	t.Run("SetSize() sets its size", func(t *testing.T) {
		setup(t)

		transform := core.NewTransform(0, 0, 0, 0)

		transform.SetSize(core.Vector{X: 3.0, Y: 5.0})
		size := transform.GetSize()

		expect(size.X).ToBe(3.0)
		expect(size.Y).ToBe(5.0)
	})
}
