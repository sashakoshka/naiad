// Naiad lets you create windows with shapes in them. It is designed for creating
// graphical user interfaces. It wraps pixel
// and provides a simple way to create and manipulate persistent shapes on screen.
// It is as cross platform as pixel is, which is to say, very. This library is not
// designed to be used in conjunction with pixel, it abstracts it away. Eventually
// it will just use raw OpenGL as a backend.
//
// The reason this exists is because the browser DOM is very useful for building
// user interfaces because you can just make things and the browser abstracts away
// all the rendering logic. However, it is inherently resource intensive, has too
// much vestigial nonsense built into it, and it wasn't even designed for
// constructing user interfaces (being more oriented towards text documents).
//
// Naiad is designed to let you have your cake and eat it too. You have full
// freedom to create whatever ridiculous shapes you want, with the ease of working
// with a DOM, and the assurance that your program will not end up being a 127
// megabyte, cpu eating, ram eating piece of garbage.
//
// On-screen shapes in Naiad are structured using groups. These groups contain
// paths (and soon text as well). At the top there is a root group which you are
// not allowed to mess with, but you can add things to it and remove things from
// it through methods defined on the window object.
//
// Paths, in turn, are just lists of points. Each point contains information like
// position, color, and line cap type. The entire path contains stuff like line
// thickness, and if it should be filled or stroked.
package naiad

import "github.com/faiface/pixel/pixelgl"

// Run starts naiad. All graphics and window manipulation needs to be performed
// within the callback function.
func Run(callback func()) {
	pixelgl.Run(callback)
}
