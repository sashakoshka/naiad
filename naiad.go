package naiad

// import "github.com/faiface/pixel"
import "github.com/faiface/pixel/pixelgl"

func Run (callback func ()) {
	pixelgl.Run(callback)
}
