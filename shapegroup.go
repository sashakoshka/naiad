package naiad

import "github.com/faiface/pixel"
import "github.com/faiface/pixel/pixelgl"

// TODO: put this at the root of the window instead of a list of shapes.
type ShapeGroup struct {
	shapeBase
	shapes []Shape
	canvas *pixelgl.Canvas
	calculatedW float64
	calculatedH float64
}

func NewShapeGroup (
	x, y float64,
) (
	group *ShapeGroup,
) {
	group = &ShapeGroup { }
	group.SetPosition(V(x, y))
	return
}

func (group *ShapeGroup) Push (shape Shape) {
	shape.setParent(group)
	group.shapes = append(group.shapes, shape)
	group.calculateBounds()
	group.SetDirty()
}

func (group *ShapeGroup) Pop () (shape Shape) {
	shape.setParent(nil)
	shape = group.shapes[len(group.shapes) - 1]
	group.shapes = group.shapes[:len(group.shapes) - 1]
	group.calculateBounds()
	group.SetDirty()
	return
}

func (group *ShapeGroup) Kind () (kind ShapeKind) {
	return ShapeKindGroup
}

func (group *ShapeGroup) draw (target pixel.Target) {
	// if we don't have a canvas, or the canvas is not the size we need,
	// create a new one.
	needNewCanvas :=
		group.canvas == nil ||
		group.canvas.Bounds().Max.X != group.calculatedW ||
		group.canvas.Bounds().Max.Y != group.calculatedH
	
	if needNewCanvas {
		group.SetDirty()
		group.canvas = pixelgl.NewCanvas (pixel.R (
			0, 0,
			group.calculatedW,
			group.calculatedH))
	}

	// draw all shapes
	if group.Dirty () {
		for _, shape := range group.shapes {
			shape.draw(target)
		}
	}

	// TODO: draw group's target onto target
	// group.artist.Draw(target)
	group.SetClean()
}

func (group *ShapeGroup) calculateBounds () {
	if len(group.shapes) > 0 {
		group.min,
		group.max = group.shapes[0].GetBounds()
	} else {
		group.min = Vector { }
		group.max = Vector { }
	}

	for _, shape := range group.shapes {
		shapeMin,
		shapeMax := shape.GetBounds()
		group.contractMin(shapeMin)
		group.expandMax(shapeMax)
	}

	group.calculatedW = group.max.Y() - group.min.Y()
	group.calculatedH = group.max.Y() - group.min.Y()

	group.calculateTransform()
}
