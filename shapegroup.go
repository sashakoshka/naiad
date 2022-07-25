package naiad

import "image/color"
import "github.com/faiface/pixel"
import "github.com/faiface/pixel/pixelgl"

// ShapeGroup is a group containing other shapes. It has its own internal buffer
// that the shapes inside of it will be drawn to, reducing the need for redraws.
// It should be used to group objects composed of several different shapes
// together, especially if they move together.
type ShapeGroup struct {
	shapeBase
	shapes []Shape
	canvas *pixelgl.Canvas
}

// NewShapeGroup creates a new shape group with the specified position and
// bounds.
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

// SetBounds sets the size of the shape group.
func (group *ShapeGroup) SetBounds (max Vector) {
	if group.max == max { return }
	group.max = max
	group.SetDirty()
}

// Push adds a new shape to the shape group.
func (group *ShapeGroup) Push (shape Shape) {
	shape.setParent(group)
	group.shapes = append(group.shapes, shape)
	group.SetDirty()
	shape.SetDirty()
}

// Pop removes the last shape from the shape group, and returns it.
func (group *ShapeGroup) Pop () (shape Shape) {
	shape.setParent(nil)
	shape = group.shapes[len(group.shapes) - 1]
	group.shapes = group.shapes[:len(group.shapes) - 1]
	group.SetDirty()
	shape.SetDirty()
	return
}

// Kind returns ShapeKindGroup.
func (group *ShapeGroup) Kind () (kind ShapeKind) {
	return ShapeKindGroup
}

// Contains takes in mouse coordinates, and determines if they are inside of the
// shape. If it is, the returned slice will have one item pointing to this
// shape. If not, the returned slice will be nill.
func (group *ShapeGroup) Contains (position Vector) (shapes []Shape) {
	relativePosition := group.matrix.Unproject(position.pixellate())
	bounds := pixel.R(group.min.x, group.min.y, group.max.x, group.max.y)
	
	if bounds.Contains(relativePosition) {
		shapes = append(shapes, Shape(group))
	} else {
		return
	}

	// range backward over shapes, becasue we want to get the top shape
	// that's in contact with the position.
	for index := len(group.shapes) - 1; index >= 0; index -- {
		shape := group.shapes[index]
		childContains := shape.Contains(vFromPixel(relativePosition))
		
		if len(childContains) > 0 {
			// there can only be one hovered shape - so we append
			// its results to the list and return.
			shapes = append(shapes, childContains...)
			return
		}
	}
	return
}

// draw draws the shape group onto the specified target, and marks the group as
// clean. If the group is marked as dirty, it will redraw all shapes inside of
// it. Otherwise, it will just draw what's in its buffer. This method
// automatically resizes the shape group's internal buffer as needed.
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
