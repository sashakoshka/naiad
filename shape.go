package naiad

import "image/color"
import "github.com/faiface/pixel"
import "github.com/faiface/pixel/imdraw"

type ShapeKind int

const (
	ShapeKindPath = iota
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

func (base *shapeBase) Dirty () (isDirty bool) {
	return !base.clean
}

func (base *shapeBase) SetDirty () {
	base.clean = false
}

func (base *shapeBase) SetClean () {
	base.clean = true
}
