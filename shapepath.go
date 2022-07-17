package naiad

import "github.com/faiface/pixel"
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
	shape.matrix = pixel.IM.Moved(pixel.V(x, y))
	return
}

func (shape *ShapePath) Push (point Point) {
	shape.points = append(shape.points, point)
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
