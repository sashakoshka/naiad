package naiad

import "time"

// Await waits for an event to occur, or the timeout to elapse. It then redraws
// the screen if needed. If the timeout is zero, there won't actually be a
// timeout and it will just wait forever for an event.
func (window *Window) Await (timeout time.Duration) {
	if window.window == nil { return }
	window.draw()
	window.window.UpdateInputWait(timeout)
	window.processEvents()
}

// Poll polls events, and redraws the screen if needed. This is non-blocking.
func (window *Window) Poll () {
	if window.window == nil { return }
	window.draw()
	window.window.Update()
	window.processEvents()
}

// processEvents reacts to any events that have been recieved, and redraws the
// screen if needed
func (window *Window) processEvents () {
	newSize := vFromPixel(window.window.Bounds().Max)
	if newSize != window.size {
		window.root.SetBounds(newSize)
	}
	window.size = newSize

	window.mousePreviousPosition = window.mousePosition
	window.mousePosition         = vFromPixel(window.window.MousePosition())
	if window.mousePosition != window.mousePreviousPosition {
		window.detectMouseOver()
	}
}

// detectMouseOver detects which shape is being hovered over by the mouse.
func (window *Window) detectMouseOver () {
	window.mouseHover = window.root.Contains(window.mousePosition)
}
