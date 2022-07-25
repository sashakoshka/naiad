package naiad

// Closed returns. whether the window is closed
func (window *Window) Closed () (closed bool) {
	if window.window == nil { return true }
	return window.window.Closed()
}
