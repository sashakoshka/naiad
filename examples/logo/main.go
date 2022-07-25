package main

import "os"
import "image"
import "image/color"
import _ "image/png"
import "github.com/sashakoshka/naiad"

func main () {
	naiad.Run(run)
}

func run () {
	iconFile, err := os.Open("icon.png")
	if err != nil { panic(err) }
	icon, _, err := image.Decode(iconFile)
	if err != nil { panic(err) }
	iconFile.Close()

	window := naiad.Window { }
	window.SetTitle("Naiad Logo")
	window.SetSize(naiad.V(336, 80))
	window.SetTransparent(true)
	window.SetIcon(icon)
	window.Open()

	colors := []color.Color {
		color.RGBA {	R: 32,	G: 0,	B: 255,	A: 255	}, // blue
		color.RGBA {	R: 252,	G: 0,	B: 212,	A: 255	}, // magenta
		color.RGBA {	R: 252,	G: 72,	B: 0,	A: 255	}, // red
		color.RGBA {	R: 252,	G: 104,	B: 0,	A: 255	}, // orange
		color.RGBA {	R: 252,	G: 240,	B: 0,	A: 255	}, // yellow
		color.RGBA {	R: 28,	G: 252,	B: 0,	A: 255	}, // green
		
		color.RGBA {	R: 0,	G: 0,	B: 0,	A: 0	},
		color.RGBA {	R: 252,	G: 104,	B: 0,	A: 0	}, // orange
	}

	logo := naiad.NewShapeGroup(0, 0, 336, 80)
	window.Push(logo)

	// N
	nGroup := naiad.NewShapeGroup(0, 0, 80, 80)
	shape := naiad.NewShapePath(8, 8)
	shape.Push(naiad.P(0,  0,  colors[1], naiad.LineCapRound))
	shape.Push(naiad.P(64, 0,  colors[6], naiad.LineCapRound))
	shape.SetThickness(8)
	nGroup.Push(shape)
	shape = naiad.NewShapePath(8, 8)
	shape.Push(naiad.P(0,  0,  colors[1], naiad.LineCapRound))
	shape.Push(naiad.P(0,  64, colors[0], naiad.LineCapRound))
	shape.Push(naiad.P(64,  0, colors[3], naiad.LineCapRound))
	shape.SetThickness(8)
	shape.SetOpen(true)
	nGroup.Push(shape)
	shape = naiad.NewShapePath(72, 8)
	shape.Push(naiad.P(0, 0,  colors[3], naiad.LineCapRound))
	shape.Push(naiad.P(0, 64, colors[4], naiad.LineCapRound))
	shape.SetThickness(8)
	nGroup.Push(shape)
	logo.Push(nGroup)

	// A
	aGroup := naiad.NewShapeGroup(80, 0, 80, 80)
	shape = naiad.NewShapePath(8, 24)
	shape.Push(naiad.P(0,  0, colors[6], naiad.LineCapRound))
	shape.Push(naiad.P(48, 0, colors[5], naiad.LineCapRound))
	shape.SetThickness(8)
	aGroup.Push(shape)
	shape = naiad.NewShapePath(8, 8)
	shape.Push(naiad.P(0,  0,  colors[3], naiad.LineCapRound))
	shape.Push(naiad.P(0,  64, colors[4], naiad.LineCapRound))
	shape.Push(naiad.P(64, 0,  colors[5], naiad.LineCapRound))
	shape.SetThickness(8)
	shape.SetOpen(true)
	aGroup.Push(shape)
	logo.Push(aGroup)

	// I
	iGroup := naiad.NewShapeGroup(160, 0, 80, 80)
	shape = naiad.NewShapePath(8, 8)
	shape.Push(naiad.P(0, 0,  colors[5], naiad.LineCapRound))
	shape.Push(naiad.P(0, 64, colors[0], naiad.LineCapRound))
	shape.SetThickness(8)
	iGroup.Push(shape)
	logo.Push(iGroup)

	// A
	a2Group := naiad.NewShapeGroup(176, 0, 80, 80)
	shape = naiad.NewShapePath(8, 24)
	shape.Push(naiad.P(0,  0, colors[6], naiad.LineCapRound))
	shape.Push(naiad.P(48, 0, colors[1], naiad.LineCapRound))
	shape.SetThickness(8)
	a2Group.Push(shape)
	shape = naiad.NewShapePath(8, 8)
	shape.Push(naiad.P(0,  0,  colors[5], naiad.LineCapRound))
	shape.Push(naiad.P(0,  64, colors[0], naiad.LineCapRound))
	shape.Push(naiad.P(64, 0,  colors[1], naiad.LineCapRound))
	shape.SetThickness(8)
	shape.SetOpen(true)
	a2Group.Push(shape)
	logo.Push(a2Group)

	// D
	dGroup := naiad.NewShapeGroup(256, 0, 80 ,80)
	shape = naiad.NewShapePath(8, 8)
	shape.Push(naiad.P(0,  0,  colors[0], naiad.LineCapRound))
	shape.Push(naiad.P(64, 64, colors[1], naiad.LineCapRound))
	shape.Push(naiad.P(0,  64, colors[2], naiad.LineCapRound))
	shape.SetThickness(8)
	dGroup.Push(shape)
	logo.Push(dGroup)

	window.Poll()

	for !window.Closed() {
		if window.IsHovered(dGroup) {
			shape.SetThickness(16)
		} else {
			shape.SetThickness(8)
		}
		window.Await(0)
	}
}
