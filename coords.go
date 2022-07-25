package naiad

import "github.com/faiface/pixel"

// Vector represents a coordinate pair. It can be used to specify things such as
// bounds, or 2D coordinates.
type Vector struct {
	x float64
	y float64
}

// V is a quick way to create a new vector out of x and y coordinates.
func V (x, y float64) (vector Vector) {
	vector.x = x
	vector.y = y
	return
}

// pixellate converts a naiad vector into a pixel vector.
func (vector Vector) pixellate () (vec pixel.Vec) {
	return pixel.Vec { X: vector.x, Y: vector.y }
}

// vFromPixel converts a pixel vector into a naiad vector.
func vFromPixel (vec pixel.Vec) (vector Vector) {
	return V(vec.X, vec.Y)
}

// X returns the x value of the vector
func (vector Vector) X () (x float64) {
	return vector.x
}

// Y returns thw x value of the vector
func (vector Vector) Y () (y float64) {
	return vector.y
}

// SetX sets the x value of the vector
func (vector* Vector) SetX (x float64) {
	vector.x = x
}

// SetY sets the y value of the vector
func (vector *Vector) SetY (y float64) {
	vector.y = y
}
