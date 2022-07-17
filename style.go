package naiad

type RGBA struct {
	R uint32
	G uint32
	B uint32
	A uint32
}

func (rgba RGBA) RGBA () (r, g, b, a uint32) {
	return rgba.R, rgba.G, rgba.B, rgba.A
}

type Style struct {
	thickness float64
	open      bool
}

func (style *Style) Thickness () (thickness float64) {
	return style.thickness
}

func (style *Style) Open () (closed bool) {
	return style.open
}

func (style *Style) SetThickness (thickness float64) {
	style.thickness = thickness
}

func (style *Style) SetOpen (open bool) {
	style.open = open
}
