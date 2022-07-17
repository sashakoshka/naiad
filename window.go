package naiad

import "time"
import "github.com/faiface/pixel/imdraw"
import "github.com/faiface/pixel/pixelgl"

type Window struct {
	Vector
	
	title       string
	// TODO: icon
	bounds      Rectangle
	position    Vector
	// TODO: monitor
	// Equivalent of Resizable, but the other way around. Being able to
	// resize the window is the default behavior.
	fixedSize   bool
	undecorated bool
	// Equivalent of NoIconify, but the other way around. Not iconifying the
	// fullscreen window on focus loss is the default behavior.
	autoIconify bool
	alwaysOnTop bool
	transparent bool
	noVsync     bool
	maximized   bool

	boundsDirty bool
	artist *imdraw.IMDraw
	shapes []Shape
	
	window *pixelgl.Window
}

/* Show brings the window on screen. This should only be called once.
 */
func (window *Window) Show () (err error) {
	if window.window != nil { return }

	window.artist = imdraw.New(nil)
	window.window, err = pixelgl.NewWindow (pixelgl.WindowConfig {
		Title:                  window.title,
		Bounds:                 window.bounds.pixellate(),
		Position:               window.position.pixellate(),
		Resizable:             !window.fixedSize,
		Undecorated:            window.undecorated,
		NoIconify:             !window.autoIconify,
		AlwaysOnTop:            window.alwaysOnTop,
		TransparentFramebuffer: window.transparent,
		VSync:                 !window.noVsync,
		Maximized:		window.maximized,
	})

	return
}

/* SetTitle sets the window title.
 */
func (window *Window) SetTitle (title string) {
	if window.window != nil {
		window.window.SetTitle(title)
	}

	window.title = title
}

/* SetBounds sets the bounds of the window to a rectangle.
 */
func (window *Window) SetBounds (bounds Rectangle) {
	if window.window != nil {
		window.window.SetBounds(bounds.pixellate())
	}

	window.bounds = bounds
}

/* Bounds returns the bounds of the window as a rectangle
 */
func (window *Window) Bounds () (bounds Rectangle) {
	return window.bounds
}

/* SetTransparent sets whether or not the window has a transparent framebuffer,
 * if supported.
 */
func (window *Window) SetTransparent (transparent bool) {
	if window.window != nil { return }
	window.transparent = transparent
}

/* Closed returns. whether the window is closed
 */
func (window *Window) Closed () (closed bool) {
	if window.window == nil { return true }
	return window.window.Closed()
}

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
	newBounds := rFromPixel(window.window.Bounds())
	window.draw(newBounds != window.bounds)
	window.bounds = newBounds
}

/* draw redraws all shapes that need to be redrawn. If force is set to true, it
 * will clear the window and redraw all shapes regardless.
 */
func (window *Window) draw (force bool) {
	if window.window == nil { return }

	updated := force
	if force {
		window.window.Clear(RGBA { 0, 0, 0, 0 })
	}

	window.artist.Clear()
	for _, shape := range window.shapes {
		if !shape.Dirty() && !force { continue }

		shape.draw(window.artist)
		shape.SetClean()
		updated = true
	}

	if updated {
		window.artist.Draw(window.window)
	}
	window.window.SwapBuffers()
}
