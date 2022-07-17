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

func (window *Window) Show () (err error) {
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

func (window *Window) SetTitle (title string) {
	if window.window != nil {
		window.window.SetTitle(title)
	}

	window.title = title
}

func (window *Window) SetBounds (bounds Rectangle) {
	if window.window != nil {
		window.window.SetBounds(bounds.pixellate())
	}

	window.bounds = bounds
}

func (window *Window) Bounds () (bounds Rectangle) {
	return window.bounds
}

func (window *Window) SetTransparent (transparent bool) {
	if window.window != nil { return }
	window.transparent = transparent
}

func (window *Window) Closed () (closed bool) {
	if window.window == nil { return true }
	return window.window.Closed()
}

func (window *Window) Await (timeout time.Duration) {
	if window.window == nil { return }
	window.window.UpdateInputWait(timeout)

	newBounds := rFromPixel(window.window.Bounds())
	window.draw(newBounds != window.bounds)
	window.bounds = newBounds
}

func (window *Window) Poll () {
	if window.window == nil { return }
	window.window.Update()
	
	newBounds := rFromPixel(window.window.Bounds())
	window.draw(newBounds != window.bounds)
	window.bounds = newBounds
}

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
