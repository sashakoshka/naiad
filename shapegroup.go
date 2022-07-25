package naiad

import "image/color"
import "github.com/faiface/pixel"
import "github.com/faiface/pixel/pixelgl"

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
	defer group.SetClean()
	
	// don't even attempt to do anything if bounds are less than one
	if group.max.X() < 1 || group.max.Y() < 1 {
		return
	}

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
	// jesus christ this is ugly. i wish pixel didn't have a coordinate
	// system dragged out from the deepest layer of hell.
	group.canvas.Draw (
		target,
		group.matrix.Moved(group.max.pixellate().Scaled(0.5)))
}
