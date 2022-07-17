package naiad

func (window *Window) AddShape (shape Shape) {
	window.shapes = append(window.shapes, shape)
}
