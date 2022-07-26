package naiad

import "time"
import "github.com/faiface/pixel/pixelgl"

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

// processEvents reacts to any events that have been received, and redraws the
// screen if needed.
func (window *Window) processEvents () {
	// update window size
	newSize := vFromPixel(window.window.Bounds().Max)
	if newSize != window.size {
		window.ShapeGroup.SetBounds(newSize)
	}
	window.size = newSize

	// update mouse position and hover
	window.mousePreviousPosition = window.mousePosition
	window.mousePosition         = vFromPixel(window.window.MousePosition())
	if window.mousePosition != window.mousePreviousPosition {
		window.mouseHover = window.Contains(window.mousePosition)
	}

	// update mouse buttons
	window.mouseLeftClick = nil
	if window.window.JustPressed(pixelgl.MouseButtonLeft) {
		window.mouseLeftHold = window.mouseHover
	} else if window.window.JustReleased(pixelgl.MouseButtonLeft) {
		// set leftclick to the intersection of hover and lefthold
		for index, holdShape := range window.mouseLeftHold {
			// as soon as these two slices differ, we have reached
			// the end of any possible intersection
			if index     >= len(window.mouseHover)   { break }
			if holdShape != window.mouseHover[index] { break }

			window.mouseLeftClick = append (
				window.mouseLeftClick,
				holdShape)
		}
		window.mouseLeftHold = nil
	}
}
