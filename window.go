package naiad

import "image"
import "image/color"
import "github.com/faiface/pixel"
import "github.com/faiface/pixel/pixelgl"

// Window represents an on-screen window.
type Window struct {
	title    string
	icon     []image.Image
	size     Vector
	position Vector
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

	mousePosition         Vector
	mousePreviousPosition Vector

	mouseHover     []Shape
	mouseLeftHold  []Shape
	mouseLeftClick []Shape

	boundsDirty bool

	window *pixelgl.Window

	ShapeGroup
}

// Open brings the window on screen. If the window has already been opened, this
// method does nothing.
func (window *Window) Open() (err error) {
	if window.window != nil {
		return
	}

	var icon []pixel.Picture
	for _, size := range window.icon {
		icon = append(icon, pixel.PictureDataFromImage(size))
	}

	window.ShapeGroup.SetBounds(V(window.size.X(), window.size.Y()))
	window.ShapeGroup.SetPosition(V(0, 0))

	window.window, err = pixelgl.NewWindow(pixelgl.WindowConfig{
		Title: window.title,
		Icon:  icon,
		Bounds: pixel.R(
			0, 0,
			window.size.X(),
			window.size.Y()),
		Position:               window.position.pixellate(),
		Resizable:              !window.fixedSize,
		Undecorated:            window.undecorated,
		NoIconify:              !window.autoIconify,
		AlwaysOnTop:            window.alwaysOnTop,
		TransparentFramebuffer: window.transparent,
		VSync:                  !window.noVsync,
		Maximized:              window.maximized,
	})

	return
}

// draw redraws all shapes that need to be redrawn. If force is set to true, it
// will clear the window and redraw all shapes regardless.
func (window *Window) draw() {
	if window.window == nil {
		return
	}

	if window.Dirty() {
		window.window.Clear(color.RGBA{0, 0, 0, 0})
		window.ShapeGroup.draw(window.window)
	}

	window.window.SwapBuffers()
}
