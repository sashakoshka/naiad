package naiad

import "image"
import "github.com/faiface/pixel"

//  SetTitle sets the window title.
func (window *Window) SetTitle (title string) {
	if window.window != nil {
		window.window.SetTitle(title)
	}

	window.title = title
}

// SetIcon takes in different resolutions of the same icon (all images) and sets
// the window icon. This will not do anything after open has been called.
func (window *Window) SetIcon (icon ...image.Image) {
	if window.window != nil {
		return
	}	

	window.icon = icon
}

// SetBounds sets the size of the window to the dimensions specified by a
// vector.
func (window *Window) SetBounds (size Vector) {
	if window.window != nil {
		window.window.SetBounds (pixel.R (
			0, 0,
			window.size.X(),
			window.size.Y()))
	}

	window.size = size
}

// Bounds returns the dimensions of the window as a vector.
func (window *Window) Bounds () (min, max Vector) {
	return V(0, 0), window.size
}

// TODO: create X, Y, SetX, and SetY methods

// SetTransparent sets whether or not the window has a transparent framebuffer,
// if supported.
func (window *Window) SetTransparent (transparent bool) {
	if window.window != nil { return }
	window.transparent = transparent
}
