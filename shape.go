package naiad

import "image/color"
import "github.com/faiface/pixel"
import "github.com/faiface/pixel/imdraw"

type ShapeKind int

const (
	ShapeKindPolygon = iota
	ShapeKindPolyline
	ShapeKindRectangle
	ShapeKindCircle
	ShapeKindCircleArc
	ShapeKindEllipse
	ShapeKindEllipseArc
	ShapeKindText
)

type LineCap int

const (
	LineCapNone = iota
	LineCapSharp
	LineCapRound
)

type Point struct {
	Vector
	color color.Color
	cap   LineCap
}

func P (x, y float64, color color.Color, cap LineCap) (point Point) {
	point.x = x
	point.y = y
	point.color = color
	point.cap   = cap
	return
}

type Dirty struct {
	clean bool
}

type Transform struct {
	matrix pixel.Matrix
}

type Shape interface {
	/* draw draws the shape onto the specified target.
	 */
	draw (artist *imdraw.IMDraw)

	/* Kind returns what kind of shape it is.
	 */
	Kind () (kind ShapeKind)

	/* Dirty returns wether the shape is dirty or not.
	 */
	Dirty () (isDirty bool)

	/* SetDirty causes the shape to be flagged as dirty.
	 */
	SetDirty ()

	/* SetClean causes the shape to be flagged as clean.
	 */
	SetClean ()
}

type shapeBase struct {
	Transform
	Style

	clean bool
}

type shapePoly struct {
	points []Point
}

func (shape *shapePoly) Push (point Point) {
	shape.points = append(shape.points, point)
}

func (base *shapeBase) Dirty () (isDirty bool) {
	return !base.clean
}

func (base *shapeBase) SetDirty () {
	base.clean = false
}

func (base *shapeBase) SetClean () {
	base.clean = true
}

type ShapePolygon struct {
	shapeBase
	shapePoly
}

type ShapePolyline struct {
	shapeBase
	shapePoly
}

type ShapeRectangle struct {
	shapeBase
	start Point
	end   Point
}

type ShapeCircle struct {
	shapeBase
	center Point
	size   float64
}

type ShapeCircleArc struct {
	shapeBase
	center Point
	size   float64
	low    float64
	high   float64
}

func (polygon *ShapePolygon) Kind () (kind ShapeKind) {
	return ShapeKindPolygon
}

func (polygon *ShapePolyline) Kind () (kind ShapeKind) {
	return ShapeKindPolyline
}

func (rectangle *ShapeRectangle) Kind () (kind ShapeKind) {
	return ShapeKindRectangle
}

func (circle *ShapeCircle) Kind () (kind ShapeKind) {
	return ShapeKindCircle
}

func (circle *ShapeCircleArc) Kind () (kind ShapeKind) {
	return ShapeKindCircleArc
}

func (polygon *ShapePolygon) draw (artist *imdraw.IMDraw) {
	artist.SetMatrix(polygon.matrix)

	for _, point := range polygon.points {
		artist.Color = point.color
		artist.EndShape = imdraw.EndShape(point.cap)
		artist.Push(point.pixellate())
	}

	artist.Polygon(polygon.thickness)
}

func (polyline *ShapePolyline) draw (artist *imdraw.IMDraw) {
	artist.SetMatrix(polyline.matrix)
	
	for _, point := range polyline.points {
		artist.Color    = point.color
		artist.EndShape = imdraw.EndShape(point.cap)
		artist.Push(point.pixellate())
	}

	artist.Line(polyline.thickness)
}

func (rectangle *ShapeRectangle) draw (artist *imdraw.IMDraw) {
	artist.SetMatrix(rectangle.matrix)
	
	artist.Color    = rectangle.start.color
	artist.EndShape = imdraw.EndShape(rectangle.start.cap)
	artist.Push(rectangle.start.pixellate())
	
	artist.Color    = rectangle.end.color
	artist.EndShape = imdraw.EndShape(rectangle.end.cap)
	artist.Push(rectangle.end.pixellate())

	artist.Rectangle(rectangle.thickness)
}

func (circle *ShapeCircle) draw (artist *imdraw.IMDraw) {
	artist.SetMatrix(circle.matrix)
	artist.Push(circle.center.pixellate())
	artist.Circle(circle.size, circle.thickness)
}
