package naiad

import "github.com/faiface/pixel/pixelgl"

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
// bounds the previous time window.Await or window.Poll etc. was called.
func (window *Window) MousePreviousPosition () (previousPosition Vector) {
	return window.mousePreviousPosition
}

// MouseHover returns a list of all shapes the mouse is hoverint over. The first
// shape is the least specific, and the last shape is the most specific.
func (window *Window) MouseHover () (hoveredShapes []Shape) {
	return window.mouseHover
}

func shapeIsIn (shape Shape, shapes []Shape) (within bool) {
	for _, checkShape := range shapes {
		if checkShape == shape { return true }
	}
	return false
}

// IsHovered returns whether or not the specified shape is being hovered over
// by the mouse.
func (window *Window) IsHovered (shape Shape) (hovered bool) {
	return shapeIsIn(shape, window.mouseHover)
}

// MouseLeftHold returns a list of all shapes that were under the mouse when its
// left button started being held down. If the mouse's left button is not
// currently being held down, this function returns nil.
func (window *Window) MouseLeftHold () (heldShapes []Shape) {
	return window.mouseLeftHold
}

// IsLeftHeld returns whether or not the specified shape is being held by the
// left mouse button.
func (window *Window) IsLeftHeld (shape Shape) (held bool) {
	return shapeIsIn(shape, window.mouseLeftHold)
}

// MouseLeftClicked returns a list of alll shapes that were clicked with the
// left mouse button. This emulates behavior found in most other programs, where
// for an object to be clicked the mouse must be held down in that object, and
// then released in that object.
func (window *Window) MouseLeftClick () (clickedShapes []Shape) {
	return window.mouseLeftClick
}

// IsLeftClicked returns whether or not the specified shape was just clicked
// with the left mouse button.
func (window *Window) IsLeftClicked (shape Shape) (clicked bool) {
	return shapeIsIn(shape, window.mouseLeftClick)
}

// Typed returns the string of characters typed since the window input was last
// updated.
func (window *Window) Typed () (text string) {
	return window.window.Typed()
}

// Pressed returns whether a key is currently being pressed.
func (window *Window) Pressed (key Key) (pressed bool) {
	return window.window.Pressed(pixelgl.Button(key))
}

// Repeated returns whether a repeat event is being triggered for a key (when it
// has been held down).
func (window *Window) Repeated (key Key) (repeated bool) {
	return window.window.Repeated(pixelgl.Button(key))
}

// JustPressed returns whether a key has just been pressed down.
func (window *Window) JustPressed (key Key) (pressed bool) {
	return window.window.JustPressed(pixelgl.Button(key))
}

// JustReleased returns whether a key has just been released.
func (window *Window) JustReleased (key Key) (released bool) {
	return window.window.JustReleased(pixelgl.Button(key))
}
