package naiad

import "github.com/faiface/pixel/imdraw"

type ShapePath struct {
	shapeBase
	points []Point
}

func NewShapePath (
	x, y float64,
) (
	shape *ShapePath,
) {
	shape = &ShapePath { }
	shape.SetPosition(V(x, y))
	return
}

func (shape *ShapePath) Push (point Point) {
	shape.points = append(shape.points, point)
	shape.calculateBounds()
}

func (shape *ShapePath) Pop () (point Point) {
	point = shape.points[len(shape.points) - 1]
	shape.points = shape.points[:len(shape.points) - 1]
	shape.calculateBounds()
	return
}

func (polygon *ShapePath) Kind () (kind ShapeKind) {
	return ShapeKindPath
}

func (shape *ShapePath) draw (artist *imdraw.IMDraw) {
	artist.SetMatrix(shape.matrix)

	for _, point := range shape.points {
		artist.Color = point.color
		artist.EndShape = imdraw.EndShape(point.cap)
		artist.Push(point.pixellate())
	}

	if shape.Open() {
		artist.Line(shape.thickness)
	} else {
		artist.Polygon(shape.thickness)
	}
}

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

	shape.recalculateTransform()
}
