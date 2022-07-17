package naiad

func (window *Window) AddShape (shape Shape) {
	window.shapes.PushBack(shape)
}
