package naiad

import "github.com/faiface/pixel/imdraw"

// TODO: put this at the root of the window instead of a list of shapes.
type ShapeGroup struct {
	shapeBase
	shapes []Shape
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
	group.shapes = append(group.shapes, shape)
	group.calculateBounds()
	group.SetDirty()
}

func (group *ShapeGroup) Pop () (shape Shape) {
	shape = group.shapes[len(group.shapes) - 1]
	group.shapes = group.shapes[:len(group.shapes) - 1]
	group.calculateBounds()
	group.SetDirty()
	return
}

func (polygon *ShapeGroup) Kind () (kind ShapeKind) {
	return ShapeKindPath
}

func (shape *ShapeGroup) draw (artist *imdraw.IMDraw) {
	// TODO: if internal buffer is nil, or bounds does not match, resize
	// buffer and force redraw.
	// TODO: range over shapes, and if any of them are dirty, redraw all of
	// them to internal buffer.
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

	group.recalculateTransform()
}
