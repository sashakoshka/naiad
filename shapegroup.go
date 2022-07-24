package naiad

import "image/color"
import "github.com/faiface/pixel"
import "github.com/faiface/pixel/pixelgl"

// TODO: put this at the root of the window instead of a list of shapes.
type ShapeGroup struct {
	shapeBase
	shapes []Shape
	canvas *pixelgl.Canvas
}

func NewShapeGroup (
	x, y float64,
	w, h float64,
) (
	group *ShapeGroup,
) {
	group = &ShapeGroup { }
	group.SetPosition(V(x, y))
	group.SetBounds(V(w, h))
	return
}

func (group *ShapeGroup) SetBounds (max Vector) {
	group.max = max
	group.SetDirty()
}

func (group *ShapeGroup) Push (shape Shape) {
	shape.setParent(group)
	group.shapes = append(group.shapes, shape)
	group.SetDirty()
}

func (group *ShapeGroup) Pop () (shape Shape) {
	shape.setParent(nil)
	shape = group.shapes[len(group.shapes) - 1]
	group.shapes = group.shapes[:len(group.shapes) - 1]
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
		group.canvas.Bounds().Max != group.max.pixellate()
	
	if needNewCanvas {
		group.SetDirty()
		group.canvas = pixelgl.NewCanvas (pixel.R (
			0, 0,
			group.max.X(),
			group.max.Y()))
	}

	// draw all shapes
	if group.Dirty () {
		group.canvas.Clear(color.RGBA { 0, 0, 0, 0 })
		for _, shape := range group.shapes {
			shape.draw(group.canvas)
		}
	}

	// draw group's canvas onto target
	group.canvas.Draw(target, group.matrix)
	group.SetClean()
}
