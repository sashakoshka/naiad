package naiad

import "image/color"
import "github.com/faiface/pixel"

type ShapeKind int

const (
	ShapeKindPath = iota
	ShapeKindText
	ShapeKindGroup
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

type Shape interface {
	/* draw draws the shape onto the specified target.
	 */
	// TODO: only include target, have each shape create its own artist.
	draw (target pixel.Target)

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

	/* GetBounds returns the shape's bounds, mapped to real coordinates on
	 * the screen.
	 */
	GetBounds () (min, max Vector)

	/* setParent shets the shape's parent.
	 */
	setParent (parent Shape)
}

type shapeBase struct {
	Style

	parent Shape
	
	matrix   pixel.Matrix
	position Vector

	min     Vector
	max     Vector
	realMin Vector
	realMax Vector
	
	clean bool
}

func (base *shapeBase) setParent (parent Shape) {
	base.parent = parent
}

func (base *shapeBase) SetPosition (position Vector) {
	base.position = position
	base.calculateTransform()
	base.SetDirty()
}

func (base *shapeBase) SetX (x float64) {
	base.position.SetX(x)
	base.calculateTransform()
	base.SetDirty()
}

func (base *shapeBase) SetY (y float64) {
	base.position.SetY(y)
	base.calculateTransform()
	base.SetDirty()
}

func (base *shapeBase) X () (x float64) {
	return base.position.X()
}

func (base *shapeBase) Y () (y float64) {
	return base.position.Y()
}

func (base *shapeBase) SetThickness (thickness float64) {
	base.Style.thickness = thickness
	base.SetDirty()
}

func (base *shapeBase) SetOpen (open bool) {
	base.Style.open = open
	base.SetDirty()
}

func (base *shapeBase) contractMin (min Vector) {
	if min.X() < base.min.X() {
		base.min.SetX(min.X())
	}
	
	if min.Y() < base.min.Y() {
		base.min.SetY(min.Y())
	}
}

func (base *shapeBase) expandMax (max Vector) {
	if max.X() > base.max.X() {
		base.max.SetX(max.X())
	}
	
	if max.Y() > base.max.Y() {
		base.max.SetY(max.Y())
	}
}

func (base *shapeBase) calculateTransform () {
	// recalculate matrix
	base.matrix = pixel.IM.Moved (
		pixel.V(base.position.X(), base.position.Y()))

	// TODO: this will not work for rotation. need to go over all points and
	// project them, then find bounds again.
	minVector := base.matrix.Project(base.min.pixellate())
	maxVector := base.matrix.Project(base.max.pixellate())

	// the shape bounds need to encompass everything that gets drawn - so we
	// must account for border thickness.
	thicknessOffset := base.Thickness() / 2
	minVector = minVector.Add(pixel.V(-thicknessOffset, -thicknessOffset))
	maxVector = minVector.Add(pixel.V( thicknessOffset,  thicknessOffset))
	
	base.realMin = vFromPixel(minVector)
	base.realMax = vFromPixel(maxVector)
}

func (base *shapeBase) GetBounds () (min, max Vector) {
	return base.realMin, base.realMax
}

func (base *shapeBase) Dirty () (isDirty bool) {
	return !base.clean
}

func (base *shapeBase) SetDirty () {
	base.clean = false

	// if this shape needs to be redrawn, then so do all of its parents.
	if base.parent != nil {
		base.parent.SetDirty()
	}
}

func (base *shapeBase) SetClean () {
	base.clean = true
}
