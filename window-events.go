package naiad

import "time"

/* Await waits for an event to occur, or the timeout to elapse. It then redraws
 * the screen if needed. If the timeout is zero, there won't actually be a
 * timeout and it will just wait forever for an event.
 */
func (window *Window) Await (timeout time.Duration) {
	if window.window == nil { return }
	window.window.UpdateInputWait(timeout)
	window.processEvents()
}

/* Poll polls events, and redraws the screen if needed. This is non-blocking.
 */
func (window *Window) Poll () {
	if window.window == nil { return }
	window.window.Update()
	window.processEvents()
}

/* processEvents reacts to any events that have been recieved, and redraws the
 * screen if needed
 */
func (window *Window) processEvents () {
	newSize := vFromPixel(window.window.Bounds().Max)
	window.draw(newSize != window.size)
	window.size = newSize
}

/* detectMouseOver detects which shape is being hovered over by the mouse.
 */
// func (window *Window) detectMouseOver () {
	// 
// }
