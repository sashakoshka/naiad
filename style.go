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
	precision int
}

func (style *Style) Thickness () (thickness float64) {
	return style.thickness
}

func (style *Style) Precision () (precision int) {
	return style.precision
}

func (style *Style) SetThickness (thickness float64) {
	style.thickness = thickness
}

func (style *Style) SetPrecision (precision int) {
	style.precision = precision
}
