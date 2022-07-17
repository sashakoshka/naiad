package naiad

import "github.com/faiface/pixel"

func NewShapeRectangle (
	x, y float64,
	start, end Point,
) (
	shape *ShapeRectangle,
) {
	shape = &ShapeRectangle { }
	shape.matrix = pixel.IM.Moved(pixel.V(x, y))
	shape.start  = start
	shape.end    = end
	return
}
