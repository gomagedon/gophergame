package behaviors_test

import (
	"fmt"
	"testing"

	"github.com/gomagedon/expectate"
	"github.com/gomagedon/gophergame/behaviors"
	"github.com/gomagedon/gophergame/engine"
)

func TestTransform(t *testing.T) {
	var expect expectate.ExpectorFunc

	setup := func(t *testing.T) {
		expect = expectate.Expect(t)
	}

	// Test
	t.Run("GetPosition() returns its position", func(t *testing.T) {
		setup(t)

		transform := behaviors.NewTransform(10.0, 20.0, 0, 0)
		position := transform.GetPosition()

		expect(position.X).ToBe(10.0)
		expect(position.Y).ToBe(20.0)
	})

	// Test
	t.Run("GetScale() returns its scale", func(t *testing.T) {
		setup(t)

		transform := behaviors.NewTransform(0, 0, 11.0, 21.0)
		scale := transform.GetScale()

		expect(scale.X).ToBe(11.0)
		expect(scale.Y).ToBe(21.0)
	})

	// Test
	t.Run("GetBox() returns x, y, w, and h", func(t *testing.T) {
		setup(t)

		transform := behaviors.NewTransform(5.0, 10.0, 15.0, 20.0)
		box := transform.GetBox()
		expect(box).ToEqual(engine.Box{
			X: 5.0,
			Y: 10.0,
			W: 15.0,
			H: 20.0,
		})
	})

	// Test
	t.Run("Translate()", func(t *testing.T) {
		t.Run("Changes its position", func(t *testing.T) {
			setup(t)

			transform := behaviors.NewTransform(10.0, 20.0, 0, 0)
			engine.NewEntity("foo").AddBehavior(behaviors.TRANSFORM, transform)

			transform.Translate(engine.Vector{X: 3, Y: -5})
			position := transform.GetPosition()

			expect(position.X).ToBe(13.0)
			expect(position.Y).ToBe(15.0)
		})
		// Test
		t.Run("Moves all child entities", func(t *testing.T) {
			setup(t)

			parent := engine.NewEntity("parent")
			parentTransform := behaviors.NewTransform(10, 20, 30, 40)
			parent.AddBehavior(behaviors.TRANSFORM, parentTransform)

			child1 := engine.NewEntity("child1")
			child1Transform := behaviors.NewTransform(0, 0, 32, 32)
			child1.AddBehavior(behaviors.TRANSFORM, child1Transform)

			child2 := engine.NewEntity("child2")
			child2Transform := behaviors.NewTransform(-10, -10, 32, 32)
			child2.AddBehavior(behaviors.TRANSFORM, child2Transform)

			parent.Children().Add(child1)
			parent.Children().Add(child2)

			parentTransform.Translate(engine.Vector{X: 5, Y: 5})

			parent.Update(0.0)

			expect(child1Transform.GetPosition()).ToEqual(engine.Vector{X: 5, Y: 5})
			expect(child2Transform.GetPosition()).ToEqual(engine.Vector{X: -5, Y: -5})
		})
	})

	// Test
	t.Run("Scale()", func(t *testing.T) {
		// Test
		t.Run("Multiplies its scale", func(t *testing.T) {
			setup(t)

			transform := behaviors.NewTransform(0, 0, 20.0, 40.0)
			engine.NewEntity("foo").AddBehavior(behaviors.TRANSFORM, transform)

			transform.Scale(engine.Vector{X: 2.5, Y: 0.25})
			scale := transform.GetScale()

			expect(scale.X).ToBe(50.0)
			expect(scale.Y).ToBe(10.0)
		})

		// Test
		t.Run("Scales all child entities and proportional distance", func(t *testing.T) {
			type TestCase struct {
				parentBox engine.Box
				child1Box engine.Box
				child2Box engine.Box

				scaleBy engine.Vector

				expectedChild1Box engine.Box
				expectedChild2Box engine.Box
			}

			tt := []TestCase{
				{
					parentBox: engine.Box{X: 100, Y: 100, W: 32, H: 32},
					child1Box: engine.Box{X: 90, Y: 90, W: 16, H: 16},
					child2Box: engine.Box{X: 120, Y: 110, W: 16, H: 16},

					scaleBy: engine.Vector{X: 2, Y: 1},

					expectedChild1Box: engine.Box{X: 80, Y: 90, W: 32, H: 16},
					expectedChild2Box: engine.Box{X: 140, Y: 110, W: 32, H: 16},
				},
				{
					parentBox: engine.Box{X: 100, Y: 100, W: 32, H: 32},
					child1Box: engine.Box{X: 90, Y: 90, W: 16, H: 16},
					child2Box: engine.Box{X: 120, Y: 110, W: 16, H: 16},

					scaleBy: engine.Vector{X: 1, Y: 2},

					expectedChild1Box: engine.Box{X: 90, Y: 80, W: 16, H: 32},
					expectedChild2Box: engine.Box{X: 120, Y: 120, W: 16, H: 32},
				},
				{
					parentBox: engine.Box{X: 100, Y: 100, W: 32, H: 32},
					child1Box: engine.Box{X: 90, Y: 90, W: 16, H: 16},
					child2Box: engine.Box{X: 120, Y: 110, W: 16, H: 16},

					scaleBy: engine.Vector{X: 3, Y: 0.5},

					expectedChild1Box: engine.Box{X: 70, Y: 95, W: 48, H: 8},
					expectedChild2Box: engine.Box{X: 160, Y: 105, W: 48, H: 8},
				},
			}

			var parentTransform *behaviors.Transform
			var child1Transform *behaviors.Transform
			var child2Transform *behaviors.Transform

			setupEntitiesWithTransforms := func(tc TestCase) {
				parent := engine.NewEntity("parent")
				child1 := engine.NewEntity("child1")
				child2 := engine.NewEntity("child2")

				parent.Children().Add(child1)
				parent.Children().Add(child2)

				parentTransform = behaviors.NewTransformFromBox(tc.parentBox)
				parent.AddBehavior(behaviors.TRANSFORM, parentTransform)

				child1Transform = behaviors.NewTransformFromBox(tc.child1Box)
				child1.AddBehavior(behaviors.TRANSFORM, child1Transform)

				child2Transform = behaviors.NewTransformFromBox(tc.child2Box)
				child2.AddBehavior(behaviors.TRANSFORM, child2Transform)
			}

			for i, tc := range tt {
				// Test
				t.Run(fmt.Sprint(i), func(t *testing.T) {
					setup(t)

					setupEntitiesWithTransforms(tc)

					parentTransform.Scale(tc.scaleBy)

					expect(child1Transform.GetBox()).ToEqual(tc.expectedChild1Box)
					expect(child2Transform.GetBox()).ToEqual(tc.expectedChild2Box)
				})
			}
		})
	})

	// Test
	t.Run("SetPosition()", func(t *testing.T) {
		t.Run("Sets its position", func(t *testing.T) {
			setup(t)

			transform := behaviors.NewTransform(0, 0, 0, 0)

			transform.SetPosition(engine.Vector{X: 11.0, Y: 30.0})
			position := transform.GetPosition()

			expect(position.X).ToBe(11.0)
			expect(position.Y).ToBe(30.0)
		})
	})

	// Test
	t.Run("SetScale() sets its scale", func(t *testing.T) {
		setup(t)

		transform := behaviors.NewTransform(0, 0, 0, 0)

		transform.SetScale(engine.Vector{X: 3.0, Y: 5.0})
		scale := transform.GetScale()

		expect(scale.X).ToBe(3.0)
		expect(scale.Y).ToBe(5.0)
	})
}
