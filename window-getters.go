package naiad

// Closed returns. whether the window is closed
func (window *Window) Closed () (closed bool) {
	if window.window == nil { return true }
	return window.window.Closed()
}

// MousePosition returns the current position of the mouse inside the window
// bounds.
func (window *Window) MousePosition () (position Vector) {
	return window.mousePosition
}

// MousePreviousPosition returns the position of the mouse inside the window
// bounds the previous time window.Await or window.Poll etc. was called
func (window *Window) MousePreviousPosition () (previousPosition Vector) {
	return window.mousePreviousPosition
}

// MouseHover returns a list of all shapes the mouse is hoverint over. The first
// shape is the least specific, and the last shape is the most specific.
func (window *Window) MouseHover () (hoveredShapes []Shape) {
	return window.mouseHover
}

// IsHovered returns whether or not the specified shape is being hovered over
// by the mouse.
func (window *Window) IsHovered (shape Shape) (hovered bool) {
	for _, checkShape := range window.mouseHover {
		if checkShape == shape { return true }
	}
	return false
}
