package naiad

import "image"
import "image/color"
import "github.com/faiface/pixel"
import "github.com/faiface/pixel/pixelgl"

type Window struct {
	title       string
	icon        []image.Image
	size        Vector
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
	
	window *pixelgl.Window
	root   *ShapeGroup
}

/* Open brings the window on screen. If the window has already been opened, this
 * method does nothing.
 */
func (window *Window) Open () (err error) {
	if window.window != nil { return }

	var icon []pixel.Picture
	for _, size := range window.icon {
		icon = append(icon, pixel.PictureDataFromImage(size))
	}
	
	window.root = NewShapeGroup(0, 0, window.size.X(), window.size.Y())

	window.window, err = pixelgl.NewWindow (pixelgl.WindowConfig {
		Title:  window.title,
		Icon:   icon,
		Bounds: pixel.R (
			0, 0,
			window.size.X(),
			window.size.Y()),
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

/* draw redraws all shapes that need to be redrawn. If force is set to true, it
 * will clear the window and redraw all shapes regardless.
 */
func (window *Window) draw () {
	if window.window == nil { return }

	if window.root.Dirty() {
		window.window.Clear(color.RGBA { 0, 0, 0, 0 })	
		window.root.draw(window.window)
	}
	
	window.window.SwapBuffers()
}

/* Push adds a shape to the window's root group.
 */
func (window *Window) Push (shape Shape) {
	window.root.Push(shape)
}

/* Pop removes the last added shape from the window's root group.
 */
func (window *Window) Pop () (shape Shape) {
	return window.root.Pop()
}
