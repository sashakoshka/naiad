package naiad

import "github.com/faiface/pixel"

func NewShapePolygon (
	x, y float64,
) (
	shape *ShapePolygon,
) {
	shape = &ShapePolygon { }
	shape.matrix = pixel.IM.Moved(pixel.V(x, y))
	return
}

func (shape *ShapePolygon) Push (point Point) {
	shape.points = append(shape.points, point)
}
