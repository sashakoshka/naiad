package naiad

import "image/color"
import "github.com/faiface/pixel"
import "github.com/faiface/pixel/imdraw"

// LineCap determines how corners and ends of lines should look. Each point has
// one of these values.
type LineCap int

const (
	LineCapNone LineCap = iota
	LineCapSharp
	LineCapRound
)

// Point represents a single point on a path. It has information about its
// position, color, and line cap type.
type Point struct {
	Vector
	color color.Color
	cap   LineCap
}

// P provides a quick way to create a point given a position, color, and line
// cap type.
func P (x, y float64, color color.Color, cap LineCap) (point Point) {
	point.x = x
	point.y = y
	point.color = color
	point.cap   = cap
	return
}

// ShapePath is a polygon (or line) composed of a set of points.
type ShapePath struct {
	shapeBase
	points []Point
	artist *imdraw.IMDraw
	thickness float64
	open      bool
}

// NewShapePath creates a new path at the given x and y coordinates.
func NewShapePath (
	x, y float64,
) (
	shape *ShapePath,
) {
	shape = &ShapePath { }
	shape.SetPosition(V(x, y))
	return
}

// Push adds a point to the shape.
func (shape *ShapePath) Push (point Point) {
	shape.points = append(shape.points, point)
	shape.calculateBounds()
	shape.SetDirty()
}

// Pop removes the last point from the shape and returns it.
func (shape *ShapePath) Pop () (point Point) {
	point = shape.points[len(shape.points) - 1]
	shape.points = shape.points[:len(shape.points) - 1]
	shape.calculateBounds()
	shape.SetDirty()
	return
}

// Kind returns ShapeKindPath
func (shape *ShapePath) Kind () (kind ShapeKind) {
	return ShapeKindPath
}

// SetThickness sets the stroke thickness of the shape. If the thickness is 0,
// the shape will be filled instead.
func (shape *ShapePath) SetThickness (thickness float64) {
	if shape.thickness == thickness { return }
	shape.thickness = thickness
	shape.SetDirty()
}

// SetOpen sets whether or not the shape is open on one end. Set this to false
// to create a polyline as opposed to a polygon.
func (shape *ShapePath) SetOpen (open bool) {
	if shape.open == open { return }
	shape.open = open
	shape.SetDirty()
}

// Thickness returns the stroke thickness of the shape.
func (shape *ShapePath) Thickness () (thickness float64) {
	return shape.thickness
}

// Open returns whether or not the shape is open on one end.
func (shape *ShapePath) Open () (closed bool) {
	return shape.open
}

// draw draws the shape onto the specified target. Paths don't make use of the
// clean/dirty flag yet, but this resets it anyway.
func (shape *ShapePath) draw (target pixel.Target) {
	defer shape.SetClean()

	if shape.artist == nil {
		shape.artist = imdraw.New(nil)
	}

	shape.artist.Clear()
	shape.artist.SetMatrix(shape.matrix)

	for _, point := range shape.points {
		shape.artist.Color = point.color
		shape.artist.EndShape = imdraw.EndShape(point.cap)
		shape.artist.Push(point.pixellate())
	}

	if shape.Open() {
		shape.artist.Line(shape.thickness)
	} else {
		shape.artist.Polygon(shape.thickness)
	}
	
	shape.artist.Draw(target)
}

// calculateBounds iterates through all points and determines the bounds of the
// shape relative to its origin.
func (shape *ShapePath) calculateBounds () {
	if len(shape.points) > 0 {
		shape.min = shape.points[0].Vector
		shape.max = shape.points[0].Vector
	} else {
		shape.min = Vector { }
		shape.max = Vector { }
	}

	for _, point := range shape.points {
		shape.contractMin(point.Vector)
		shape.expandMax(point.Vector)
	}

	shape.calculateTransform()
}
