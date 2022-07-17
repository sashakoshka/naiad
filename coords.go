package naiad

import "github.com/faiface/pixel"

type Vector struct {
	x float64
	y float64
}

type Rectangle struct {
	Vector
	w float64
	h float64
}

func V (x, y float64) (vector Vector) {
	vector.x = x
	vector.y = y
	return
}

func R (x, y, w, h float64) (rectangle Rectangle) {
	rectangle.x = x
	rectangle.y = y
	rectangle.w = w
	rectangle.h = h
	return
}

func (vector Vector) pixellate () (vec pixel.Vec) {
	return pixel.V(vector.x, vector.y)
}

func (rectangle Rectangle) pixellate () (rect pixel.Rect) {
	return pixel.R (
		rectangle.x,
		rectangle.y,
		rectangle.x + rectangle.w,
		rectangle.y + rectangle.h)
}

func vFromPixel (vec pixel.Vec) (vector Vector) {
	return V(vec.X, vec.Y)
}

func rFromPixel (rect pixel.Rect) (rectangle Rectangle) {
	return R (
		rect.Min.X,
		rect.Min.Y,
		rect.Max.X - rect.Min.X,
		rect.Max.Y - rect.Min.Y)
}

func (vector Vector) X () (x float64) {
	return vector.x
}

func (vector Vector) Y () (y float64) {
	return vector.y
}

func (rectangle Rectangle) W () (w float64) {
	return rectangle.w
}

func (rectangle Rectangle) H () (h float64) {
	return rectangle.h
}
