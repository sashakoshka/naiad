package naiad

import "github.com/faiface/pixel"

type Vector struct {
	x float64
	y float64
}

func V (x, y float64) (vector Vector) {
	vector.x = x
	vector.y = y
	return
}

func (vector Vector) pixellate () (vec pixel.Vec) {
	return pixel.V(vector.x, vector.y)
}

func vFromPixel (vec pixel.Vec) (vector Vector) {
	return V(vec.X, vec.Y)
}

func (vector Vector) X () (x float64) {
	return vector.x
}

func (vector Vector) Y () (y float64) {
	return vector.y
}
