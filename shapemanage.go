package naiad

func (window *Window) AddShape (shape Shape) {
	window.shapes.PushBack(shape)
}

func (window *Window) RemoveShape (shape Shape) {
	element := window.shapes.Front()
	for element != nil {
		if shape == element.Value { break }
		element = element.Next()
	}
}
