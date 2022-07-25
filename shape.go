package naiad

import "github.com/faiface/pixel"

// ShapeKind determines what sort of shape a Shape should be cast to. It can be
// retrieved by calling the Kind() method.
type ShapeKind int

const (
	ShapeKindPath = iota
	ShapeKindText
	ShapeKindGroup
)

// Shape is an interface representing an on-screen shape. For a shape to be
// inserted into naiad's shape heirarchy, it must support these behaviors.
type Shape interface {
	// draw draws the shape onto the specified target.
	draw (target pixel.Target)

	// Kind returns what kind of shape it is.
	Kind () (kind ShapeKind)

	// Dirty returns wether the shape is dirty or not.
	Dirty () (isDirty bool)

	// SetDirty causes the shape to be flagged as dirty.
	SetDirty ()

	// SetClean causes the shape to be flagged as clean.
	SetClean ()

	// Bounds returns the bounds of the shape, relative to its origin. This
	// is usually in the top left, which means min will usually be (0, 0).
	// Hovever, in shapes such as paths, the min bound may be in a different
	// spot due to things such as stroke thickness and points with negative
	// coordinates.
	Bounds () (min, max Vector)

	// TODO: have a method that takes in mouse coordinates, unprojects them,
	// and returns what shapes (including itself, if need be) that are being
	// hovered over by the mouse. Have an argument that determines if the
	// method should recurse into child objects so that a breadth-first
	// search can be accomplished.

	// setParent sets the shape's parent. This does not actually parent the
	// shape - it should be called by the parent as the shape is being
	// parented.
	setParent (parent Shape)
}

// shapeBase is a struct which should be included in all shapes. It defines some
// basic behaviors and properties such as position, bounds, and dirtiness.
type shapeBase struct {
	Style

	parent Shape
	
	matrix   pixel.Matrix
	position Vector

	min Vector
	max Vector
	
	clean bool
}

// SetPosition sets the position of the shape.
func (base *shapeBase) SetPosition (position Vector) {
	if base.position != position {
		base.position = position
		base.SetDirty()
	}
	base.calculateTransform()
}

// SetX sets the x position of the shape.
func (base *shapeBase) SetX (x float64) {
	if base.position.X() != x {
		base.position.SetX(x)
		base.SetDirty()
	}
	base.calculateTransform()
}

// SetY sets the y position of the shape.
func (base *shapeBase) SetY (y float64) {
	if base.position.Y() != y {
		base.position.SetY(y)
		base.SetDirty()
	}
	base.calculateTransform()
}

// X returns the x position of the shape.
func (base *shapeBase) X () (x float64) {
	return base.position.X()
}

// Y returns the y position of the shape.
func (base *shapeBase) Y () (y float64) {
	return base.position.Y()
}


// TODO: remove the style struct. Shapes are too different and need their own
// independent set of styling.
func (base *shapeBase) SetThickness (thickness float64) {
	if base.Style.thickness == thickness { return }
	base.Style.thickness = thickness
	base.SetDirty()
}

func (base *shapeBase) SetOpen (open bool) {
	if base.Style.open == open { return }
	base.Style.open = open
	base.SetDirty()
}

// Dirty returns wether the shape is dirty or not.
func (base *shapeBase) Dirty () (isDirty bool) {
	return !base.clean
}

// SetDirty causes the shape to be flagged as dirty.
func (base *shapeBase) SetDirty () {
	if base.Dirty() { return }
	base.clean = false

	// if this shape needs to be redrawn, then so do all of its parents.
	if base.parent != nil {
		base.parent.SetDirty()
	}
}

// SetClean causes the shape to be flagged as clean.
func (base *shapeBase) SetClean () {
	base.clean = true
}

// Bounds returns the bounds of the shape, relative to its origin. This is
// usually in the top left, which means min will usually be (0, 0). Hovever, in
// shapes such as paths, the min bound may be in a different spot due to things
// such as stroke thickness and points with negative coordinates.
func (base *shapeBase) Bounds () (min, max Vector) {
	return base.min, base.max
}

// setParent sets the shape's parent. This does not actually parent the shape -
// it should be called by the parent as the shape is being parented.
func (base *shapeBase) setParent (parent Shape) {
	base.parent = parent
}

// contractMin ensures that the shape's minimum bound is no greater than min.
func (base *shapeBase) contractMin (min Vector) {
	if min.X() < base.min.X() {
		base.min.SetX(min.X())
	}
	
	if min.Y() < base.min.Y() {
		base.min.SetY(min.Y())
	}
}

// expandMax ensures that the shape's maximum bound is no less than max.
func (base *shapeBase) expandMax (max Vector) {
	if max.X() > base.max.X() {
		base.max.SetX(max.X())
	}
	
	if max.Y() > base.max.Y() {
		base.max.SetY(max.Y())
	}
}

// calculateTransform recalcualtes the transformation matrix of the shape.
func (base *shapeBase) calculateTransform () {
	// recalculate matrix
	base.matrix = pixel.IM.Moved (
		pixel.V(base.position.X(), base.position.Y()))
}
