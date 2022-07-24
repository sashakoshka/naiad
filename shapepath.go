package naiad

import "github.com/faiface/pixel"
import "github.com/faiface/pixel/imdraw"

type ShapePath struct {
	shapeBase
	points []Point
	artist *imdraw.IMDraw
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
	shape.SetDirty()
}

func (shape *ShapePath) Pop () (point Point) {
	point = shape.points[len(shape.points) - 1]
	shape.points = shape.points[:len(shape.points) - 1]
	shape.calculateBounds()
	shape.SetDirty()
	return
}

func (shape *ShapePath) Kind () (kind ShapeKind) {
	return ShapeKindPath
}

func (shape *ShapePath) draw (target pixel.Target) {
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

	shape.calculateTransform()
	shape.SetClean()
}
