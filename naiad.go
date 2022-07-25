package naiad

import "github.com/faiface/pixel/pixelgl"

// Run starts naiad. All graphics and window manipulation needs to be performed
// within the callback function.
func Run (callback func ()) {
	pixelgl.Run(callback)
}
